package copier

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// These flags define options for tag handling
const (
	// Denotes that a destination field must be copied to. If copying fails then a panic will ensue.
	tagMust uint8 = 1 << iota

	// Denotes that the program should not panic when the must flag is on and
	// value is not copied. The program will return an error instead.
	tagNoPanic

	// Ignore a destation field from being copied to.
	tagIgnore

	// Denotes that the value as been copied
	hasCopied
)

// Copy copy things
func Copy(toValue interface{}, fromValue interface{}) (err error) {
	return copy(toValue, fromValue, false)
}

type Option struct {
	IgnoreEmpty bool
}

// CopyWithOption copy with option
func CopyWithOption(toValue interface{}, fromValue interface{}, option Option) (err error) {
	return copy(toValue, fromValue, option.IgnoreEmpty)
}

func copy(toValue interface{}, fromValue interface{}, ignoreEmpty bool) (err error) {
	var (
		isSlice bool
		amount  = 1
		from    = indirect(reflect.ValueOf(fromValue)) //read note 判断是指针还是结构体类型来返回对应的参数的值
		to      = indirect(reflect.ValueOf(toValue))   //read note 判断是指针还是结构体类型来返回对应的参数的值
	)

	//read note 不可寻址的数据类型可以参考：https://github.com/hyper0x/Golang_Puzzlers/blob/master/src/puzzlers/article15/q1/demo35.go
	if !to.CanAddr() {
		return errors.New("copy to value is unaddressable")
	}

	// Return is from value is invalid
	if !from.IsValid() {
		return
	}

	//read note 返回参数的具体类型
	fromType := indirectType(from.Type())
	toType := indirectType(to.Type())

	// Just set it if possible to assign
	// And need to do copy anyway if the type is struct
	//read note 判断是否非结构体并且可以直接赋值
	if fromType.Kind() != reflect.Struct && from.Type().AssignableTo(to.Type()) {
		to.Set(from)
		return
	}

	//read note from和to都是map
	if fromType.Kind() == reflect.Map && toType.Kind() == reflect.Map {
		//read note 判断map的key的结构类型是否可以转换，因为这边已经判断是map，所以直接通过 .key来获取，不担心是否报错
		if !fromType.Key().ConvertibleTo(toType.Key()) {
			return
		}
		//read note 判断要转换的Map是否为空，为空则进行初始化
		if to.IsNil() {
			to.Set(reflect.MakeMapWithSize(toType, from.Len()))
		}
		//read note 遍历Map的所有key
		for _, k := range from.MapKeys() {
			//read note 根据to的key类型创建一个新的key
			toKey := indirect(reflect.New(toType.Key()))
			//read note 设置key的值
			if !set(toKey, k) {
				continue
			}

			//read note 设置value值
			toValue := indirect(reflect.New(toType.Elem()))
			if !set(toValue, from.MapIndex(k)) {
				//read note 对嵌套的结构体进行copy
				err = Copy(toValue.Addr().Interface(), from.MapIndex(k).Interface())
				if err != nil {
					continue
				}
			}
			to.SetMapIndex(toKey, toValue)
		}
	}

	//read note 只要有一个数据结构不是结构体直接返回?
	if fromType.Kind() != reflect.Struct || toType.Kind() != reflect.Struct {
		return
	}

	//read note 切片处理：设置切片的长度

	//read note 如果Result是Slice类型
	if to.Kind() == reflect.Slice {
		//read note 设置isSlice
		isSlice = true
		//read note amount初始是1，如果被复制的对象也是数组，则这边修改成数组长度，否则1表示只处理结构体
		if from.Kind() == reflect.Slice {
			amount = from.Len()
		}
	}

	//read note 循环被复制的切片的长度
	for i := 0; i < amount; i++ {
		var dest, source reflect.Value

		//read note Result的结果是数组
		if isSlice {
			// source
			// read note 如果Origin的类型是数组，需要根据index进行获取结构体
			if from.Kind() == reflect.Slice {
				source = indirect(from.Index(i))
			} else {
				// read note 如果Origin的类型是结构体，直接获取该结构体
				source = indirect(from)
			}
			// dest
			dest = indirect(reflect.New(toType).Elem())
		} else {
			//read note Result的结果不是数组，直接获取结构体
			source = indirect(from)
			dest = indirect(to)
		}

		// Get tag options
		//read note 获取tag的所有标签
		tagBitFlags := map[string]uint8{}
		if dest.IsValid() {
			//read note 根据结构体type获取所有的Field对应的tag标签
			tagBitFlags = getBitFlags(toType)
		}

		// check source

		//read note 如果是非零值
		if source.IsValid() {
			//read note 获取结构体的所有Field的信息（数组）
			fromTypeFields := deepFields(fromType)
			// fmt.Printf("%#v", fromTypeFields)
			// Copy from field to field or method

			//todo 循环所有的Field处理
			for _, field := range fromTypeFields {
				name := field.Name

				// Get bit flags for field
				//read note 根据name获取tag数据
				fieldFlags, _ := tagBitFlags[name]

				// Check if we should ignore copying
				//read note ignore标签对结构体的影响处理
				if (fieldFlags & tagIgnore) != 0 {
					continue
				}

				//read note Origin的方法和Result字段同名的处理

				if fromField := source.FieldByName(name); fromField.IsValid() && !shouldIgnore(fromField, ignoreEmpty) {
					// has field

					//read note 根据名称获取Result结构体的字段.
					if toField := dest.FieldByName(name); toField.IsValid() {
						if toField.CanSet() {
							if !set(toField, fromField) {
								if err := Copy(toField.Addr().Interface(), fromField.Interface()); err != nil {
									return err
								}
							} else {
								//read note 赋值完成，设置对应的标识
								if fieldFlags != 0 {
									// Note that a copy was made

									//read note 设置复制标识
									tagBitFlags[name] = fieldFlags | hasCopied
								}
							}
						}
					} else {
						// try to set to method
						var toMethod reflect.Value
						//read note 通过名称找到对应的Method
						if dest.CanAddr() {
							toMethod = dest.Addr().MethodByName(name)
						} else {
							toMethod = dest.MethodByName(name)
						}
						//read note 【被转换对象的方法调用】的校验还比较严格,这边可以看出来只能有一个字段，并且字段类型要对应上
						if toMethod.IsValid() && toMethod.Type().NumIn() == 1 && fromField.Type().AssignableTo(toMethod.Type().In(0)) {
							//read note 调用声明的方法
							toMethod.Call([]reflect.Value{fromField})
						}
					}
				}
			}

			//read note Result方法 与Origin字段同名的处理

			// Copy from method to field
			//read note 处理目标结构体的方法，目标结构体的方法要和被复制结构体的字段名一致，就是这边控制的
			for _, field := range deepFields(toType) {
				name := field.Name

				//read note 根据Result的字段，获取Origin同名的方法
				var fromMethod reflect.Value
				if source.CanAddr() {
					fromMethod = source.Addr().MethodByName(name)
				} else {
					fromMethod = source.MethodByName(name)
				}

				//read note 如果方法符合规则，没有入参，有一个出参，则进行对应方法的调用处理
				if fromMethod.IsValid() && fromMethod.Type().NumIn() == 0 && fromMethod.Type().NumOut() == 1 && !shouldIgnore(fromMethod, ignoreEmpty) {
					if toField := dest.FieldByName(name); toField.IsValid() && toField.CanSet() {
						values := fromMethod.Call([]reflect.Value{})
						if len(values) >= 1 {
							//read note 进行字段的设值
							set(toField, values[0])
						}
					}
				}
			}
		}

		//read note 转换结果Result是切片的处理：分成两种情况，被复制的是 结构体指针 和 结构体
		if isSlice {
			if dest.Addr().Type().AssignableTo(to.Type().Elem()) {
				to.Set(reflect.Append(to, dest.Addr()))
			} else if dest.Type().AssignableTo(to.Type().Elem()) {
				to.Set(reflect.Append(to, dest))
			}
		}
		//read note 这边是不是会有一个问题，就是err是不是会被覆盖，前面的字段有错误，最后一个没有错误则会覆盖之前的error
		err = checkBitFlags(tagBitFlags)
	}
	return
}

func shouldIgnore(v reflect.Value, ignoreEmpty bool) bool {
	if !ignoreEmpty {
		return false
	}

	return v.IsZero()
}

//read note 根据结构体类型，获取结构体对应的Field切片，注意这边的【Anonymous】表示的匿名变量，匿名变量的Field这边需要特殊处理
func deepFields(reflectType reflect.Type) []reflect.StructField {
	var fields []reflect.StructField

	//read note 判断是不是结构体,只能对结构体进行处理
	if reflectType = indirectType(reflectType); reflectType.Kind() == reflect.Struct {
		//read note 循环处理对应的所有Field
		for i := 0; i < reflectType.NumField(); i++ {
			v := reflectType.Field(i)
			//read note 对【嵌入（匿名）字段】结构体 的所有结构体进行添加
			if v.Anonymous {
				fields = append(fields, deepFields(v.Type)...)
			} else {
				fields = append(fields, v)
			}
		}
	}

	return fields
}

func indirect(reflectValue reflect.Value) reflect.Value {
	for reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}
	return reflectValue
}

func indirectType(reflectType reflect.Type) reflect.Type {
	for reflectType.Kind() == reflect.Ptr || reflectType.Kind() == reflect.Slice {
		reflectType = reflectType.Elem()
	}
	return reflectType
}

func set(to, from reflect.Value) bool {

	//read note IsValid返回是否非零值的结果.所以这边的处理是针对非零值
	if from.IsValid() {

		//read note 前置条件：处理to的类型，如果from为空，则直接设置空值返回

		//	to是指针类型特殊处理
		if to.Kind() == reflect.Ptr {
			// set `to` to nil if from is nil
			//read note 如果from是空，则直接设置to为零值返回
			if from.Kind() == reflect.Ptr && from.IsNil() {
				to.Set(reflect.Zero(to.Type()))
				return true
			} else if to.IsNil() {
				//read note 如果to是nil且不满足上面的from为nil的条件，这个时候要给to设置默认值
				to.Set(reflect.New(to.Type().Elem()))
			}
			//read note 指针的转换处理
			to = to.Elem()
		}

		//read note from和to类型的转换处理，这边当from是ptr类型的时候，会调用set进行递归处理

		//read note 如果类型可以进行转换，则要设置对应的值（具体什么类型可以转换需要看一下源码，这里不多赘述）
		if from.Type().ConvertibleTo(to.Type()) {
			to.Set(from.Convert(to.Type()))
		} else if scanner, ok := to.Addr().Interface().(sql.Scanner); ok {
			//read note sql.Scanner 这个不知道具体是干嘛的.
			err := scanner.Scan(from.Interface())
			if err != nil {
				return false
			}
		} else if from.Kind() == reflect.Ptr {
			//read note from是指针类型，处理成结构体进行赋值（相当于递归再往下走）
			return set(to, from.Elem())
		} else {
			//read note 其他不能转换的直接返回false
			return false
		}
	}

	//read note 零值直接返回true，零值不处理
	return true
}

// parseTags Parses struct tags and returns uint8 bit flags.
func parseTags(tag string) (flags uint8) {
	for _, t := range strings.Split(tag, ",") {
		switch t {
		case "-":
			flags = tagIgnore
			return
		case "must":
			flags = flags | tagMust
		case "nopanic":
			flags = flags | tagNoPanic
		}
	}
	return
}

// getBitFlags Parses struct tags for bit flags.
func getBitFlags(toType reflect.Type) map[string]uint8 {
	//read note 存储的结构是  FieldName->tag对应的二进制数据(tag标签转换成程序标识)
	flags := map[string]uint8{}
	//read note 根据结构体的类型获取对应的Field切片
	toTypeFields := deepFields(toType)

	// Get a list dest of tags
	//read note 循环Field切片，获取切片对应的tag数据
	for _, field := range toTypeFields {
		tags := field.Tag.Get("copier") //tag标签是【copier】
		if tags != "" {
			//read note tag标签转换成程序处理标识（这边也是使用二进制的处理方式）
			flags[field.Name] = parseTags(tags)
		}
	}
	return flags
}

// checkBitFlags Checks flags for error or panic conditions.
func checkBitFlags(flagsList map[string]uint8) (err error) {
	// Check flag conditions were met
	//read note 循环map（FieldName->tag对应的二进制数据）
	for name, flags := range flagsList {
		//read note 如果字段没有被复制
		if flags&hasCopied == 0 {
			switch {
			case flags&tagMust != 0 && flags&tagNoPanic != 0:
				//read note 处理1：返回错误信息
				err = fmt.Errorf("Field %s has must tag but was not copied", name)
				return
			case flags&(tagMust) != 0:
				//read note 处理2：直接报错
				panic(fmt.Sprintf("Field %s has must tag but was not copied", name))
			}
		}
	}
	return
}

// The mapstructure package exposes functionality to convert an
// abitrary map[string]interface{} into a native Go structure.
//
// The Go structure can be arbitrarily complex, containing slices,
// other structs, etc. and the decoder will properly decode nested
// maps and so on into the proper structures in the native Go struct.
// See the examples to see what the decoder is capable of.
package mapstructure

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

//read note 钩子，起作用的地方是在Decode之前，这边主要的作用可以说比如 在转换之前把字符串转成Time类型等等，具体可以查看：https://godoc.org/github.com/mitchellh/mapstructure#DecodeHookFunc
//read note  https://github.com/mitchellh/mapstructure/blob/master/decode_hooks.go 这边提供了一些默认的hook方法，可以看一下
type DecodeHookFunc func(reflect.Kind, reflect.Kind, interface{}) (interface{}, error)

// DecoderConfig is the configuration that is used to create a new decoder
// and allows customization of various aspects of decoding.
type DecoderConfig struct {
	// DecodeHook, if set, will be called before any decoding and any
	// type conversion (if WeaklyTypedInput is on). This lets you modify
	// the values before they're set down onto the resulting struct.
	//
	// If an error is returned, the entire decode will fail with that
	// error.
	//read note 在Decode之前会调用该Hook,这边的注释说明,hook可以让你在Decode之前修改你的结构体现有值
	DecodeHook DecodeHookFunc

	// If ErrorUnused is true, then it is an error for there to exist
	// keys in the original map that were unused in the decoding process
	// (extra keys).
	//read note 如果这个字段设置为true的话，在map当中只要有key不能转换成结构体的字段就会进行报错，感觉正常情况下是不会设置成true...
	ErrorUnused bool

	// If WeaklyTypedInput is true, the decoder will make the following
	// "weak" conversions:
	//
	//   - bools to string (true = "1", false = "0")
	//   - numbers to string (base 10)
	//   - bools to int/uint (true = 1, false = 0)
	//   - strings to int/uint (base implied by prefix)
	//   - int to bool (true if value != 0)
	//   - string to bool (accepts: 1, t, T, TRUE, true, True, 0, f, F,
	//     FALSE, false, False. Anything else is an error)
	//   - empty array = empty map and vice versa
	//
	//read note 弱类型转换，这个字段如果设置为true，则从某种类型 to 某种类型 会按照上面的转换规则来进行转换
	WeaklyTypedInput bool

	// Metadata is the struct that will contain extra metadata about
	// the decoding. If this is nil, then no metadata will be tracked.
	//read note 没搞明白这个是干啥的
	Metadata *Metadata

	// Result is a pointer to the struct that will contain the decoded
	// value.
	//read note 转换的结果
	Result interface{}

	// The tag name that mapstructure reads for field names. This
	// defaults to "mapstructure"
	//read note tag名称，默认是 mapstructure
	TagName string
}

// A Decoder takes a raw interface value and turns it into structured
// data, keeping track of rich error information along the way in case
// anything goes wrong. Unlike the basic top-level Decode method, you can
// more finely control how the Decoder behaves using the DecoderConfig
// structure. The top-level Decode method is just a convenience that sets
// up the most basic Decoder.
type Decoder struct {
	//read note 转换工具的配置
	config *DecoderConfig
}

// Metadata contains information about decoding a structure that
// is tedious or difficult to get otherwise.
type Metadata struct {
	// Keys are the keys of the structure which were successfully decoded
	Keys []string

	// Unused is a slice of keys that were found in the raw value but
	// weren't decoded since there was no matching field in the result interface
	Unused []string
}

// Decode takes a map and uses reflection to convert it into the
// given Go native structure. val must be a pointer to a struct.
func Decode(m interface{}, rawVal interface{}) error {
	config := &DecoderConfig{
		Metadata: nil,
		Result:   rawVal,
	}

	//read note 创建包含配置的Decoder对象
	decoder, err := NewDecoder(config)
	if err != nil {
		return err
	}

	//read note 进行Decode操作
	return decoder.Decode(m)
}

// DecodePath takes a map and uses reflection to convert it into the
// given Go native structure. Tags are used to specify the mapping
// between fields in the map and structure
func DecodePath(m map[string]interface{}, rawVal interface{}) error {
	config := &DecoderConfig{
		Metadata: nil,
		Result:   nil,
	}

	decoder, err := NewPathDecoder(config)
	if err != nil {
		return err
	}

	_, err = decoder.DecodePath(m, rawVal)
	return err
}

// DecodeSlicePath decodes a slice of maps against a slice of structures that
// contain specified tags
func DecodeSlicePath(ms []map[string]interface{}, rawSlice interface{}) error {
	reflectRawSlice := reflect.TypeOf(rawSlice)
	rawKind := reflectRawSlice.Kind()
	rawElement := reflectRawSlice.Elem()

	//read note 校验是否为切片
	if (rawKind == reflect.Ptr && rawElement.Kind() != reflect.Slice) ||
		(rawKind != reflect.Ptr && rawKind != reflect.Slice) {
		return fmt.Errorf("Incompatible Value, Looking For Slice : %v : %v", rawKind, rawElement.Kind())
	}

	config := &DecoderConfig{
		Metadata: nil,
		Result:   nil,
	}

	decoder, err := NewPathDecoder(config)
	if err != nil {
		return err
	}

	// Create a slice large enough to decode all the values
	//read note 构造一个新的数组
	valSlice := reflect.MakeSlice(rawElement, len(ms), len(ms))

	// Iterate over the maps and decode each one
	//read note 循环被转换的map数组
	for index, m := range ms {
		sliceElementType := rawElement.Elem()
		if sliceElementType.Kind() != reflect.Ptr {
			// A slice of objects
			//read note 如果转换的结果是结构体类型，处理转换
			obj := reflect.New(rawElement.Elem())
			decoder.DecodePath(m, reflect.Indirect(obj))
			indexVal := valSlice.Index(index)
			indexVal.Set(reflect.Indirect(obj))
		} else {
			// A slice of pointers
			//read note 如果转换的结果是指针类型，处理转换
			obj := reflect.New(rawElement.Elem().Elem())
			decoder.DecodePath(m, reflect.Indirect(obj))
			indexVal := valSlice.Index(index)
			indexVal.Set(obj)
		}
	}

	// Set the new slice
	//read note 设置转换后的数组数据
	reflect.ValueOf(rawSlice).Elem().Set(valSlice)
	return nil
}

// NewDecoder returns a new decoder for the given configuration. Once
// a decoder has been returned, the same configuration must not be used
// again.
func NewDecoder(config *DecoderConfig) (*Decoder, error) {
	val := reflect.ValueOf(config.Result)
	if val.Kind() != reflect.Ptr {
		return nil, errors.New("result must be a pointer")
	}

	val = val.Elem()
	if !val.CanAddr() {
		return nil, errors.New("result must be addressable (a pointer)")
	}

	if config.Metadata != nil {
		if config.Metadata.Keys == nil {
			config.Metadata.Keys = make([]string, 0)
		}

		if config.Metadata.Unused == nil {
			config.Metadata.Unused = make([]string, 0)
		}
	}

	if config.TagName == "" {
		config.TagName = "mapstructure"
	}

	result := &Decoder{
		config: config,
	}

	return result, nil
}

// NewPathDecoder returns a new decoder for the given configuration.
// This is used to decode path specific structures
func NewPathDecoder(config *DecoderConfig) (*Decoder, error) {
	if config.Metadata != nil {
		if config.Metadata.Keys == nil {
			config.Metadata.Keys = make([]string, 0)
		}

		if config.Metadata.Unused == nil {
			config.Metadata.Unused = make([]string, 0)
		}
	}

	if config.TagName == "" {
		config.TagName = "mapstructure"
	}

	result := &Decoder{
		config: config,
	}

	return result, nil
}

// Decode decodes the given raw interface to the target pointer specified
// by the configuration.
func (d *Decoder) Decode(raw interface{}) error {
	//read note config中的Result配置的是指针，如果不是指针是没办法完成配置的.
	//	所有这边的调用是通过 reflect.ValueOf().Elem()
	return d.decode("", raw, reflect.ValueOf(d.config.Result).Elem())
}

// DecodePath decodes the raw interface against the map based on the
// specified tags
func (d *Decoder) DecodePath(m map[string]interface{}, rawVal interface{}) (bool, error) {
	decoded := false

	var val reflect.Value
	reflectRawValue := reflect.ValueOf(rawVal)
	kind := reflectRawValue.Kind()

	// Looking for structs and pointers to structs
	//read note 对传入的 转换结果的数据类型进行判断，转换成对应的结构体类型
	switch kind {
	case reflect.Ptr:
		val = reflectRawValue.Elem()
		if val.Kind() != reflect.Struct {
			return decoded, fmt.Errorf("Incompatible Type : %v : Looking For Struct", kind)
		}
	case reflect.Struct:
		var ok bool
		val, ok = rawVal.(reflect.Value)
		if ok == false {
			return decoded, fmt.Errorf("Incompatible Type : %v : Looking For reflect.Value", kind)
		}
	default:
		return decoded, fmt.Errorf("Incompatible Type : %v", kind)
	}

	// Iterate over the fields in the struct
	//read note 循环结构体的所有Field
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag
		//read note 这边的tag是通过 jpath来识别的
		tagValue := tag.Get("jpath")

		// Is this a field without a tag
		//read note 没有tag的单独处理
		if tagValue == "" {
			//read note 如果是结构体，调用DecodePath处理
			if valueField.Kind() == reflect.Struct {
				// We have a struct that may have indivdual tags. Process separately
				d.DecodePath(m, valueField)
				continue
			} else if valueField.Kind() == reflect.Ptr && reflect.TypeOf(valueField).Kind() == reflect.Struct {
				//read note 指针的处理，转换成结构体也是类型的DecodePath的处理

				// We have a pointer to a struct
				if valueField.IsNil() {
					// Create the object since it doesn't exist
					valueField.Set(reflect.New(valueField.Type().Elem()))
					decoded, _ = d.DecodePath(m, valueField.Elem())
					if decoded == false {
						// If nothing was decoded for this object return the pointer to nil
						valueField.Set(reflect.NewAt(valueField.Type().Elem(), nil))
					}
					continue
				}

				d.DecodePath(m, valueField.Elem())
				continue
			}
		}

		// Use mapstructure to populate the fields
		//read note jpath后面支持的别名可以是递进到更进去的层次，通过.标识，比如说 Age.Birth.  表示 Age:{Birth:100,。。。}
		keys := strings.Split(tagValue, ".")
		//read note 通过keys去查找数据
		data := d.findData(m, keys)
		//read note 如果在map中找到的数据不为nil。需要开始设值操作
		if data != nil {
			//read note 如果当前Filed的类型是Slice
			if valueField.Kind() == reflect.Slice {
				// Ignore a slice of maps - This sucks but not sure how to check
				//read note 如果Field是map数组，这边没办法处理，只能通过下面的decode方法进行具体的处理
				if strings.Contains(valueField.Type().String(), "map[") {
					goto normal_decode
				}

				// We have a slice
				mapSlice := data.([]interface{})
				if len(mapSlice) > 0 {
					// Test if this is a slice of more maps
					//read note 转换成切片的数据如果不是 map的数组，则跳转到下面通过 decode处理.因为这边是切片对切片，map对结构体
					_, ok := mapSlice[0].(map[string]interface{})
					if ok == false {
						goto normal_decode
					}

					// Extract the maps out and run it through DecodeSlicePath
					//read note 组转map数组
					ms := make([]map[string]interface{}, len(mapSlice))
					for index, m2 := range mapSlice {
						ms[index] = m2.(map[string]interface{})
					}

					//调用DecodeSlicePath，转换 map数组 -》 结构体数组
					DecodeSlicePath(ms, valueField.Addr().Interface())
					continue
				}
			}
		normal_decode:
			//read note 通过decode处理，这边同样应该支持 mapstructure的tag标签
			decoded = true
			err := d.decode("", data, valueField)
			if err != nil {
				return false, err
			}
		}
	}

	return decoded, nil
}

// Decodes an unknown data type into a specific reflection value.
//read note 入参说明：
//	name： 字段名称
//	data： 被转换的数据
// 	val ： 转换最终的结构体
func (d *Decoder) decode(name string, data interface{}, val reflect.Value) error {
	//read note 结构体为空直接返回
	if data == nil {
		// If the data is nil, then we don't set anything.
		return nil
	}

	dataVal := reflect.ValueOf(data)
	//read note 非IsValid的对象，设置零值
	if !dataVal.IsValid() {
		// If the data value is invalid, then we just set the value
		// to be the zero value.
		val.Set(reflect.Zero(val.Type()))
		return nil
	}

	//read note hook的调用，调用的结果data会用在下面的判断中
	if d.config.DecodeHook != nil {
		// We have a DecodeHook, so let's pre-process the data.
		var err error
		data, err = d.config.DecodeHook(d.getKind(dataVal), d.getKind(val), data)
		if err != nil {
			return err
		}
	}

	var err error
	//read note 获得 转换最终的结构体 的类型
	dataKind := d.getKind(val)
	//read note 依据类型进行不同的转换处理
	switch dataKind {
	case reflect.Bool:
		err = d.decodeBool(name, data, val)
	case reflect.Interface:
		err = d.decodeBasic(name, data, val)
	case reflect.String:
		err = d.decodeString(name, data, val)
	case reflect.Int:
		err = d.decodeInt(name, data, val)
	case reflect.Uint:
		err = d.decodeUint(name, data, val) //read note 处理与decodeInt类似
	case reflect.Float32:
		err = d.decodeFloat(name, data, val) //read note 处理与decodeInt类似
	case reflect.Struct:
		err = d.decodeStruct(name, data, val)
	case reflect.Map:
		err = d.decodeMap(name, data, val)
	case reflect.Slice:
		err = d.decodeSlice(name, data, val)
	default:
		// If we reached this point then we weren't able to decode it
		//read note 没办法处理指针类型,比如本来传入的就是指针的指针，这边是拒绝处理的
		return fmt.Errorf("%s: unsupported type: %s", name, dataKind)
	}

	// If we reached here, then we successfully decoded SOMETHING, so
	// mark the key as used if we're tracking metadata.
	//read note 对处理玩的Metadata进行组装，但是这边看好像一定是不会处理的，因为传进来的name一直是空字符串
	if d.config.Metadata != nil && name != "" {
		d.config.Metadata.Keys = append(d.config.Metadata.Keys, name)
	}

	return err
}

// findData locates the data by walking the keys down the map
func (d *Decoder) findData(m map[string]interface{}, keys []string) interface{} {
	//read note 这是一个递归方法，所以递归的结束就是keys最终变成一个值，查找该值，能找到则返回，否则返回nil
	if len(keys) == 1 {
		if value, ok := m[keys[0]]; ok == true {
			return value
		}
		return nil
	}

	//read note 继续进行递归，下一次递归就是往下一个key开始，可以理解就是map一层一层的下去查找对应的key.
	if value, ok := m[keys[0]]; ok == true {
		if m, ok := value.(map[string]interface{}); ok == true {
			return d.findData(m, keys[1:])
		}
	}

	return nil
}

func (d *Decoder) getKind(val reflect.Value) reflect.Kind {
	kind := val.Kind()

	switch {
	case kind >= reflect.Int && kind <= reflect.Int64:
		return reflect.Int
	case kind >= reflect.Uint && kind <= reflect.Uint64:
		return reflect.Uint
	case kind >= reflect.Float32 && kind <= reflect.Float64:
		return reflect.Float32
	default:
		return kind
	}
}

// This decodes a basic type (bool, int, string, etc.) and sets the
// value to "data" of that type.
func (d *Decoder) decodeBasic(name string, data interface{}, val reflect.Value) error {
	//read note 如果被转换的结果是interface，则只判断是否 AssignableTo
	dataVal := reflect.ValueOf(data)
	dataValType := dataVal.Type()
	if !dataValType.AssignableTo(val.Type()) {
		return fmt.Errorf(
			"'%s' expected type '%s', got '%s'",
			name, val.Type(), dataValType)
	}

	val.Set(dataVal)
	return nil
}

func (d *Decoder) decodeString(name string, data interface{}, val reflect.Value) error {
	dataVal := reflect.ValueOf(data)
	dataKind := d.getKind(dataVal)

	//read note string这边的转换规则如下：
	//	1、如果被转换的就是string，则直接转换成string
	//	开启了弱类型转换标识的
	//	2、bool转换从1或0
	//	3、数值类型按照十进制转换成字符串
	//	4、float类型，按照64位转换
	switch {
	case dataKind == reflect.String:
		val.SetString(dataVal.String())
	case dataKind == reflect.Bool && d.config.WeaklyTypedInput:
		if dataVal.Bool() {
			val.SetString("1")
		} else {
			val.SetString("0")
		}
	case dataKind == reflect.Int && d.config.WeaklyTypedInput:
		val.SetString(strconv.FormatInt(dataVal.Int(), 10))
	case dataKind == reflect.Uint && d.config.WeaklyTypedInput:
		val.SetString(strconv.FormatUint(dataVal.Uint(), 10))
	case dataKind == reflect.Float32 && d.config.WeaklyTypedInput:
		val.SetString(strconv.FormatFloat(dataVal.Float(), 'f', -1, 64))
	default:
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s'",
			name, val.Type(), dataVal.Type())
	}

	return nil
}

func (d *Decoder) decodeInt(name string, data interface{}, val reflect.Value) error {
	dataVal := reflect.ValueOf(data)
	dataKind := d.getKind(dataVal)

	//read note int这边的转换大致如下
	//	1、数值类型，直接转换成int，不考虑精度
	//	如果开启了弱类型转换标识
	//	2、bool类型按照0,1转换
	//	3、字符串通过 ParseInt转换
	//  4、否则错误
	switch {
	case dataKind == reflect.Int:
		val.SetInt(dataVal.Int())
	case dataKind == reflect.Uint:
		val.SetInt(int64(dataVal.Uint()))
	case dataKind == reflect.Float32:
		val.SetInt(int64(dataVal.Float()))
	case dataKind == reflect.Bool && d.config.WeaklyTypedInput:
		if dataVal.Bool() {
			val.SetInt(1)
		} else {
			val.SetInt(0)
		}
	case dataKind == reflect.String && d.config.WeaklyTypedInput:
		i, err := strconv.ParseInt(dataVal.String(), 0, val.Type().Bits())
		if err == nil {
			val.SetInt(i)
		} else {
			return fmt.Errorf("cannot parse '%s' as int: %s", name, err)
		}
	default:
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s'",
			name, val.Type(), dataVal.Type())
	}

	return nil
}

func (d *Decoder) decodeUint(name string, data interface{}, val reflect.Value) error {
	dataVal := reflect.ValueOf(data)
	dataKind := d.getKind(dataVal)

	switch {
	case dataKind == reflect.Int:
		val.SetUint(uint64(dataVal.Int()))
	case dataKind == reflect.Uint:
		val.SetUint(dataVal.Uint())
	case dataKind == reflect.Float32:
		val.SetUint(uint64(dataVal.Float()))
	case dataKind == reflect.Bool && d.config.WeaklyTypedInput:
		if dataVal.Bool() {
			val.SetUint(1)
		} else {
			val.SetUint(0)
		}
	case dataKind == reflect.String && d.config.WeaklyTypedInput:
		i, err := strconv.ParseUint(dataVal.String(), 0, val.Type().Bits())
		if err == nil {
			val.SetUint(i)
		} else {
			return fmt.Errorf("cannot parse '%s' as uint: %s", name, err)
		}
	default:
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s'",
			name, val.Type(), dataVal.Type())
	}

	return nil
}

//read note 入参说明：
//	name： 字段名称
//	data： 被转换的数据
// 	val ： 转换最终的结构体
func (d *Decoder) decodeBool(name string, data interface{}, val reflect.Value) error {
	dataVal := reflect.ValueOf(data)
	dataKind := d.getKind(dataVal)

	//read note 这边的转换规则如下：
	//	1、bool类型直接转换
	//	开启了弱类型转换标识的
	//	2、数值类型判断是否为0，0是false
	//	3、string类型，先进行bool转换，如果不能转换，空字符表示false，否则错误
	//	4、其他类型错误
	switch {
	case dataKind == reflect.Bool:
		val.SetBool(dataVal.Bool())
	case dataKind == reflect.Int && d.config.WeaklyTypedInput:
		val.SetBool(dataVal.Int() != 0)
	case dataKind == reflect.Uint && d.config.WeaklyTypedInput:
		val.SetBool(dataVal.Uint() != 0)
	case dataKind == reflect.Float32 && d.config.WeaklyTypedInput:
		val.SetBool(dataVal.Float() != 0)
	case dataKind == reflect.String && d.config.WeaklyTypedInput:
		b, err := strconv.ParseBool(dataVal.String())
		if err == nil {
			val.SetBool(b)
		} else if dataVal.String() == "" {
			val.SetBool(false)
		} else {
			return fmt.Errorf("cannot parse '%s' as bool: %s", name, err)
		}
	default:
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s'",
			name, val.Type(), dataVal.Type())
	}

	return nil
}

func (d *Decoder) decodeFloat(name string, data interface{}, val reflect.Value) error {
	dataVal := reflect.ValueOf(data)
	dataKind := d.getKind(dataVal)

	switch {
	case dataKind == reflect.Int:
		val.SetFloat(float64(dataVal.Int()))
	case dataKind == reflect.Uint:
		val.SetFloat(float64(dataVal.Uint()))
	case dataKind == reflect.Float32:
		val.SetFloat(float64(dataVal.Float()))
	case dataKind == reflect.Bool && d.config.WeaklyTypedInput:
		if dataVal.Bool() {
			val.SetFloat(1)
		} else {
			val.SetFloat(0)
		}
	case dataKind == reflect.String && d.config.WeaklyTypedInput:
		f, err := strconv.ParseFloat(dataVal.String(), val.Type().Bits())
		if err == nil {
			val.SetFloat(f)
		} else {
			return fmt.Errorf("cannot parse '%s' as float: %s", name, err)
		}
	default:
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s'",
			name, val.Type(), dataVal.Type())
	}

	return nil
}

//read note 入参说明：
//	name： 字段名称
//	data： 被转换的数据
// 	val ： 转换最终的结构体
func (d *Decoder) decodeMap(name string, data interface{}, val reflect.Value) error {
	valType := val.Type()
	valKeyType := valType.Key()
	valElemType := valType.Elem()

	// Make a new map to hold our result
	//read note 通过反射创建新的map
	mapType := reflect.MapOf(valKeyType, valElemType)
	valMap := reflect.MakeMap(mapType)

	// Check input type
	dataVal := reflect.Indirect(reflect.ValueOf(data))
	//read note 如果被转换的结构体不是map，这边需要满足
	//	1、弱类型标识打开
	//	2、空切片或空数组
	//	才会返回空map，否则组装错误信息
	if dataVal.Kind() != reflect.Map {
		// Accept empty array/slice instead of an empty map in weakly typed mode
		if d.config.WeaklyTypedInput &&
			(dataVal.Kind() == reflect.Slice || dataVal.Kind() == reflect.Array) &&
			dataVal.Len() == 0 {
			val.Set(valMap)
			return nil
		} else {
			return fmt.Errorf("'%s' expected a map, got '%s'", name, dataVal.Kind())
		}
	}

	// Accumulate errors
	errors := make([]string, 0)

	//read note 遍历被转换结构体的所有key
	for _, k := range dataVal.MapKeys() {
		fieldName := fmt.Sprintf("%s[%s]", name, k)

		// First decode the key into the proper type
		//read note 对key进行转换 decode
		currentKey := reflect.Indirect(reflect.New(valKeyType))
		if err := d.decode(fieldName, k.Interface(), currentKey); err != nil {
			errors = appendErrors(errors, err)
			continue
		}

		// Next decode the data into the proper type
		//read note 对value进行转换 decode
		v := dataVal.MapIndex(k).Interface()
		currentVal := reflect.Indirect(reflect.New(valElemType))
		if err := d.decode(fieldName, v, currentVal); err != nil {
			errors = appendErrors(errors, err)
			continue
		}

		//read note 对结果map进行组装
		valMap.SetMapIndex(currentKey, currentVal)
	}

	// Set the built up map to the value
	//read note map设值
	val.Set(valMap)

	// If we had errors, return those
	if len(errors) > 0 {
		return &Error{errors}
	}

	return nil
}

func (d *Decoder) decodeSlice(name string, data interface{}, val reflect.Value) error {
	dataVal := reflect.Indirect(reflect.ValueOf(data))
	dataValKind := dataVal.Kind()
	valType := val.Type()
	valElemType := valType.Elem()

	// Make a new slice to hold our result, same size as the original data.
	//read note 构造新的切片
	sliceType := reflect.SliceOf(valElemType)
	valSlice := reflect.MakeSlice(sliceType, dataVal.Len(), dataVal.Len())

	// Check input type
	//read note 同样的弱类型转换，空的map直接返回空数组（弱类型标识为true的情况下）
	if dataValKind != reflect.Array && dataValKind != reflect.Slice {
		// Accept empty map instead of array/slice in weakly typed mode
		if d.config.WeaklyTypedInput && dataVal.Kind() == reflect.Map && dataVal.Len() == 0 {
			val.Set(valSlice)
			return nil
		} else {
			return fmt.Errorf(
				"'%s': source data must be an array or slice, got %s", name, dataValKind)
		}
	}

	// Accumulate any errors
	errors := make([]string, 0)

	//read note 遍历原来的数组，进行每个元素的转换 decode
	for i := 0; i < dataVal.Len(); i++ {
		currentData := dataVal.Index(i).Interface()
		currentField := valSlice.Index(i)

		fieldName := fmt.Sprintf("%s[%d]", name, i)
		if err := d.decode(fieldName, currentData, currentField); err != nil {
			errors = appendErrors(errors, err)
		}
	}

	// Finally, set the value to the slice we built up
	//read note 设置切片的值
	val.Set(valSlice)

	// If there were errors, we return those
	if len(errors) > 0 {
		return &Error{errors}
	}

	return nil
}

//read note 入参说明：
//	name： 字段名称
//	data： 被转换的数据
// 	val ： 转换最终的结构体
func (d *Decoder) decodeStruct(name string, data interface{}, val reflect.Value) error {
	dataVal := reflect.Indirect(reflect.ValueOf(data))
	dataValKind := dataVal.Kind()
	//read note 被转换的类型不是map，错误
	if dataValKind != reflect.Map {
		return fmt.Errorf("'%s' expected a map, got '%s'", name, dataValKind)
	}

	dataValType := dataVal.Type()
	//read note map的key不是string或者Interface类型，错误
	if kind := dataValType.Key().Kind(); kind != reflect.String && kind != reflect.Interface {
		return fmt.Errorf("'%s' needs a map with string keys, has '%s' keys", name, dataValType.Key().Kind())
	}

	//read note 这边组装了两个数组，一个是映射成功的key数组，一个是未映射的key数组
	dataValKeys, dataValKeysUnused := d.getUseAndUnUseKeyList(dataVal)

	// This slice will keep track of all the structs we'll be decoding. There can be more than one struct if there are embedded structs  that are squashed.
	//read note 正常情况下只会有一个结构体才对，但是这边的处理，如果结构体中的某一个字段是匿名结构体，则会加入到structs中进行处理
	// structs[0]是最外层的结构体，structs[1:]是最外层结构体中的匿名字段代表的结构体
	structs := make([]reflect.Value, 1, 5)
	structs[0] = val
	errors := make([]string, 0)

	// Compile the list of all the fields that we're going to be decoding from all the structs.
	fields := make(map[*reflect.StructField]reflect.Value) //read note 创建【field->value】的map

	//read note 循环structs数组，在循环的过程中遇到 【匿名结构体并且标注了squash的结构体字段会加入到structs数组中】
	d.structLoopForFieldList(structs, errors, val, fields)

	//read note 循环Filed数组，对每一个map能映射到的字段进行赋值（字段名不区分大小写）
	d.fieldsLoopForDecode(fields, errors, dataVal, dataValKeys, dataValKeysUnused, name)

	//read note errorUnused标识打开，如果map有未转换的可以，则组装错误信息
	d.produceErrorWithUnused(dataValKeysUnused, name, errors)

	if len(errors) > 0 {
		return &Error{errors}
	}

	// Add the unused keys to the list of unused keys if we're tracking metadata
	//read note 组装metadata.这边还要判断name不为空
	d.produceMetadata(dataValKeysUnused, name)

	return nil
}

func (d *Decoder) structLoopForFieldList(structs []reflect.Value, errors []string, val reflect.Value, fields map[*reflect.StructField]reflect.Value) {
	for len(structs) > 0 {
		structVal := structs[0]
		structs = structs[1:]

		structType := structVal.Type()
		//read note 循环结构体的所有Field字段
		for i := 0; i < structType.NumField(); i++ {
			fieldType := structType.Field(i)

			//read note 匿名字段处理
			if fieldType.Anonymous {
				fieldKind := fieldType.Type.Kind()
				//read note 非结构体，错误
				if fieldKind != reflect.Struct {
					errors = appendErrors(errors,
						fmt.Errorf("%s: unsupported type: %s", fieldType.Name, fieldKind))
					continue
				}

				// We have an embedded field. We "squash" the fields down
				// if specified in the tag.
				//read note squash标签只作用在匿名字段上，会对匿名字段结构体的字段进行深入赋值.
				squash := false
				tagParts := strings.Split(fieldType.Tag.Get(d.config.TagName), ",")
				for _, tag := range tagParts[1:] {
					if tag == "squash" {
						squash = true
						break
					}
				}

				//read note 如果是squash标签标识，会加入到structs数组中，在下一次循环开始之前进行循环
				if squash {
					structs = append(structs, val.FieldByName(fieldType.Name))
					continue
				}
			}

			// Normal struct field, store it away
			//read note 组装map指向
			fields[&fieldType] = structVal.Field(i)
		}
	}
}

func (d *Decoder) fieldsLoopForDecode(fields map[*reflect.StructField]reflect.Value, errors []string, dataVal reflect.Value, dataValKeys map[reflect.Value]struct{}, dataValKeysUnused map[interface{}]struct{}, name string) {
	//read note 循环所有的Field，这边的field数组可能是当前结构体的，也可能是结构体中的匿名字段的field
	for fieldType, field := range fields {
		fieldName := fieldType.Name

		//read note 获取tag标签，进行【,】切割，这边的标识默认是 mapstructure，比如说 mapstructure：name.获取的就是name
		tagValue := fieldType.Tag.Get(d.config.TagName)
		tagValue = strings.SplitN(tagValue, ",", 2)[0]
		if tagValue != "" {
			fieldName = tagValue
		}

		//read note 根据key获取map中的值
		rawMapKey := reflect.ValueOf(fieldName)
		rawMapVal := dataVal.MapIndex(rawMapKey)
		//read note 循环Key数组，进行大小写不敏感匹配，获得map中可以，对应的value
		if !rawMapVal.IsValid() {
			// Do a slower search by iterating over each key and
			// doing case-insensitive search.
			for dataValKey, _ := range dataValKeys {
				mK, ok := dataValKey.Interface().(string)
				if !ok {
					// Not a string key
					continue
				}

				//read note 不区分大小写
				if strings.EqualFold(mK, fieldName) {
					rawMapKey = dataValKey
					rawMapVal = dataVal.MapIndex(dataValKey)
					break
				}
			}

			if !rawMapVal.IsValid() {
				// There was no matching key in the map for the value in
				// the struct. Just ignore.
				continue
			}
		}

		// Delete the key we're using from the unused map so we stop tracking
		//read note 未使用的key数组把被转换的key删掉
		delete(dataValKeysUnused, rawMapKey.Interface())

		//read note 非IsValid直接报错
		if !field.IsValid() {
			// This should never happen
			panic("field is not valid")
		}

		// If we can't set the field, then it is unexported or something,
		// and we just continue onwards.
		if !field.CanSet() {
			continue
		}

		// If the name is empty string, then we're at the root, and we
		// don't dot-join the fields.
		if name != "" {
			fieldName = fmt.Sprintf("%s.%s", name, fieldName)
		}

		//read note 字段再递归进去处理，可能是对应的不同的类型
		if err := d.decode(fieldName, rawMapVal.Interface(), field); err != nil {
			errors = appendErrors(errors, err)
		}
	}
}

func (d *Decoder) produceErrorWithUnused(dataValKeysUnused map[interface{}]struct{}, name string, errors []string) {
	if d.config.ErrorUnused && len(dataValKeysUnused) > 0 {
		keys := make([]string, 0, len(dataValKeysUnused))
		for rawKey, _ := range dataValKeysUnused {
			keys = append(keys, rawKey.(string))
		}
		sort.Strings(keys)

		err := fmt.Errorf("'%s' has invalid keys: %s", name, strings.Join(keys, ", "))
		errors = appendErrors(errors, err)
	}
}

func (d *Decoder) produceMetadata(dataValKeysUnused map[interface{}]struct{}, name string) {
	if d.config.Metadata != nil {
		for rawKey, _ := range dataValKeysUnused {
			key := rawKey.(string)
			if name != "" {
				key = fmt.Sprintf("%s.%s", name, key)
			}

			d.config.Metadata.Unused = append(d.config.Metadata.Unused, key)
		}
	}
}

func (d *Decoder) getUseAndUnUseKeyList(dataVal reflect.Value) (map[reflect.Value]struct{}, map[interface{}]struct{}) {
	dataValKeys := make(map[reflect.Value]struct{})
	dataValKeysUnused := make(map[interface{}]struct{})
	//read note 循环map的key，组装key数组和未使用的key数组
	for _, dataValKey := range dataVal.MapKeys() {
		dataValKeys[dataValKey] = struct{}{}
		dataValKeysUnused[dataValKey.Interface()] = struct{}{}
	}
	return dataValKeys, dataValKeysUnused
}

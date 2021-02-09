// Copyright 2013 Dario Castañé. All rights reserved.
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Based on src/pkg/reflect/deepequal.go from official
// golang's stdlib.

package mergo

import (
	"fmt"
	"reflect"
)

//read note 判断结构体是否有可导出字段
func hasMergeableFields(dst reflect.Value) (exported bool) {
	for i, n := 0, dst.NumField(); i < n; i++ {
		field := dst.Type().Field(i)
		//read note 匿名字段如果是结构体需要判断该结构体中是否有可导出字段
		if field.Anonymous && dst.Field(i).Kind() == reflect.Struct {
			exported = exported || hasMergeableFields(dst.Field(i))
		} else if isExportedComponent(&field) {
			exported = exported || len(field.PkgPath) == 0
		}
	}
	return
}

//read note 判断字段是否可导出，这边通过pkgPath和字段名首字母大小写确定
func isExportedComponent(field *reflect.StructField) bool {
	pkgPath := field.PkgPath
	//read note 查看一下pkgPath的注释，如果为空字符串表示可导出，否则不可导出
	if len(pkgPath) > 0 {
		return false
	}
	//read note 这边也是判断字段是否可导出
	c := field.Name[0]
	if 'a' <= c && c <= 'z' || c == '_' {
		return false
	}
	return true
}

//read note 配置，这边采用Option的方式进行传参设置，通过Option修改配置的字段
type Config struct {
	Overwrite                    bool         //read note 是否覆盖
	AppendSlice                  bool         //read note 是否添加到对应数组中
	TypeCheck                    bool         //read note 是否进行类型检查
	Transformers                 Transformers //read note 自定义类型转换
	overwriteWithEmptyValue      bool         //read note 是否覆盖空值
	overwriteSliceWithEmptyValue bool         //read note 是否覆盖空数组
	sliceDeepCopy                bool         //read note 数组中的元素是否进行深度克隆
	debug                        bool
}

type Transformers interface {
	Transformer(reflect.Type) func(dst, src reflect.Value) error
}

// Traverses recursively both values, assigning src's fields values to dst.
// The map argument tracks comparisons that have already been seen, which allows
// short circuiting on recursive types.
func deepMerge(dst, src reflect.Value, visited map[uintptr]*visit, depth int, config *Config) (err error) {
	overwrite := config.Overwrite
	typeCheck := config.TypeCheck
	overwriteWithEmptySrc := config.overwriteWithEmptyValue
	overwriteSliceWithEmptySrc := config.overwriteSliceWithEmptyValue
	sliceDeepCopy := config.sliceDeepCopy

	//read note 零值返回
	if !src.IsValid() {
		return
	}
	//read note 这个地方应该是根据visited判断字段是否被deepMerge
	if dst.CanAddr() {
		addr := dst.UnsafeAddr()
		h := 17 * addr
		seen := visited[h]
		typ := dst.Type()
		//read note 这边是为了避免结构体中和结构体同类型的字段赋值同样的值。如果遇到这种情况会跳过赋值，因为已经赋值过了
		for p := seen; p != nil; p = p.next {
			if p.ptr == addr && p.typ == typ {
				return nil
			}
		}
		// Remember, remember...
		//read note 记录已经被merge的字段
		visited[h] = &visit{addr, typ, seen}
	}
	//read note 这个转换方法只有在 结果结构体或结构体字段不为空的情况下才会处理
	if config.Transformers != nil && !isEmptyValue(dst) {
		if fn := config.Transformers.Transformer(dst.Type()); fn != nil {
			err = fn(dst, src)
			return
		}
	}

	switch dst.Kind() {
	case reflect.Struct:
		//read note 这边会遍历结构体的字段，判断是否有可导出的字段，如果有，才会进行Field遍历，实现deepMerge
		if hasMergeableFields(dst) {
			for i, n := 0, dst.NumField(); i < n; i++ {
				//read note 对每一个字段进行deepMerge操作
				if err = deepMerge(dst.Field(i), src.Field(i), visited, depth+1, config); err != nil {
					return
				}
			}
		} else {
			//read note 空值覆盖或者是有值覆盖，这个条件需要仔细阅读一下
			if (isReflectNil(dst) || overwrite) && (!isEmptyValue(src) || overwriteWithEmptySrc) {
				dst.Set(src)
			}
		}
	case reflect.Map:
		//read note 初始化map
		if dst.IsNil() && !src.IsNil() {
			dst.Set(reflect.MakeMap(dst.Type()))
		}

		//read note 这个地方按道理如果是两个相同的结构体，字段或者结构体的类型不可能不一样，所以这边有可能进去吗？
		if src.Kind() != reflect.Map {
			if overwrite {
				dst.Set(src)
			}
			return
		}

		//read note 遍历map的key
		for _, key := range src.MapKeys() {
			//read note 获取被merge的map的value
			srcElement := src.MapIndex(key)
			if !srcElement.IsValid() {
				continue
			}
			//read note 获取merge结果的map的value
			dstElement := dst.MapIndex(key)
			//read note fallthrough表示当前的case处理完之后还会继续往下执行一个case.
			switch srcElement.Kind() {
			case reflect.Chan, reflect.Func, reflect.Map, reflect.Interface, reflect.Slice:
				if srcElement.IsNil() {
					//read note 零值覆盖
					if overwrite {
						dst.SetMapIndex(key, srcElement)
					}
					continue
				}
				fallthrough
			default:
				if !srcElement.CanInterface() {
					continue
				}

				switch reflect.TypeOf(srcElement.Interface()).Kind() {
				case reflect.Struct:
					fallthrough
				case reflect.Ptr:
					fallthrough
				case reflect.Map:
					//read note map、struct、ptr的处理类似，都需要获取value对应值,然后往下进行deepMerge操作
					srcMapElm := srcElement
					dstMapElm := dstElement
					if srcMapElm.CanInterface() {
						srcMapElm = reflect.ValueOf(srcMapElm.Interface())
						if dstMapElm.IsValid() {
							dstMapElm = reflect.ValueOf(dstMapElm.Interface())
						}
					}
					if err = deepMerge(dstMapElm, srcMapElm, visited, depth+1, config); err != nil {
						return
					}
				case reflect.Slice:
					srcSlice := reflect.ValueOf(srcElement.Interface())

					var dstSlice reflect.Value
					//read note 如果结果字段是数组，判断是否空值进行初始化或者获取值
					if !dstElement.IsValid() || dstElement.IsNil() {
						dstSlice = reflect.MakeSlice(srcSlice.Type(), 0, srcSlice.Len())
					} else {
						dstSlice = reflect.ValueOf(dstElement.Interface())
					}

					//read note 判断直接覆盖的条件，如果满足，会把被merge的数组覆盖到原数组上
					if (!isEmptyValue(src) || overwriteWithEmptySrc || overwriteSliceWithEmptySrc) && (overwrite || isEmptyValue(dst)) && !config.AppendSlice && !sliceDeepCopy {
						if typeCheck && srcSlice.Type() != dstSlice.Type() {
							return fmt.Errorf("cannot override two slices with different type (%s, %s)", srcSlice.Type(), dstSlice.Type())
						}
						dstSlice = srcSlice
					} else if config.AppendSlice {
						//read note 如果标识数组append，则会对数组值进行append操作
						if srcSlice.Type() != dstSlice.Type() {
							return fmt.Errorf("cannot append two slices with different type (%s, %s)", srcSlice.Type(), dstSlice.Type())
						}
						dstSlice = reflect.AppendSlice(dstSlice, srcSlice)
					} else if sliceDeepCopy {
						//read note 如果数组需要深度克隆，则会遍历数组的所有元素，对每个元素进行deepMerge
						i := 0
						for ; i < srcSlice.Len() && i < dstSlice.Len(); i++ {
							srcElement := srcSlice.Index(i)
							dstElement := dstSlice.Index(i)

							if srcElement.CanInterface() {
								srcElement = reflect.ValueOf(srcElement.Interface())
							}
							if dstElement.CanInterface() {
								dstElement = reflect.ValueOf(dstElement.Interface())
							}

							if err = deepMerge(dstElement, srcElement, visited, depth+1, config); err != nil {
								return
							}
						}

					}
					//read note 设置map的值
					dst.SetMapIndex(key, dstSlice)
				}
			}
			//read note 这边的判断是如果上面的处理进行了赋值，则跳过这个key的处理,否则再往下就是默认设值的处理了
			if dstElement.IsValid() && !isEmptyValue(dstElement) && (reflect.TypeOf(srcElement.Interface()).Kind() == reflect.Map || reflect.TypeOf(srcElement.Interface()).Kind() == reflect.Slice) {
				continue
			}
			//read note 设置转换结果为被转换的字段值，如果声明了覆盖标识（override）
			if srcElement.IsValid() && ((srcElement.Kind() != reflect.Ptr && overwrite) || !dstElement.IsValid() || isEmptyValue(dstElement)) {
				if dst.IsNil() {
					dst.Set(reflect.MakeMap(dst.Type()))
				}
				dst.SetMapIndex(key, srcElement)
			}
		}
	case reflect.Slice:
		//read note 这边的slice的处理和map中对slice的处理类似，就不多赘述
		if !dst.CanSet() {
			break
		}
		if (!isEmptyValue(src) || overwriteWithEmptySrc || overwriteSliceWithEmptySrc) && (overwrite || isEmptyValue(dst)) && !config.AppendSlice && !sliceDeepCopy {
			dst.Set(src)
		} else if config.AppendSlice {
			if src.Type() != dst.Type() {
				return fmt.Errorf("cannot append two slice with different type (%s, %s)", src.Type(), dst.Type())
			}
			dst.Set(reflect.AppendSlice(dst, src))
		} else if sliceDeepCopy {
			for i := 0; i < src.Len() && i < dst.Len(); i++ {
				srcElement := src.Index(i)
				dstElement := dst.Index(i)
				if srcElement.CanInterface() {
					srcElement = reflect.ValueOf(srcElement.Interface())
				}
				if dstElement.CanInterface() {
					dstElement = reflect.ValueOf(dstElement.Interface())
				}

				if err = deepMerge(dstElement, srcElement, visited, depth+1, config); err != nil {
					return
				}
			}
		}
	case reflect.Ptr:
		fallthrough
	case reflect.Interface:
		//read note ptr和Interface的处理是一样的

		//read note 空值的处理，判断是否覆盖.
		if isReflectNil(src) {
			if overwriteWithEmptySrc && dst.CanSet() && src.Type().AssignableTo(dst.Type()) {
				dst.Set(src)
			}
			break
		}

		//read note 不是Interface类型，这边dst控制、ptr和结构体类型进行处理
		if src.Kind() != reflect.Interface {
			if dst.IsNil() || (src.Kind() != reflect.Ptr && overwrite) {
				if dst.CanSet() && (overwrite || isEmptyValue(dst)) {
					dst.Set(src)
				}
			} else if src.Kind() == reflect.Ptr {
				if err = deepMerge(dst.Elem(), src.Elem(), visited, depth+1, config); err != nil {
					return
				}
			} else if dst.Elem().Type() == src.Type() {
				if err = deepMerge(dst.Elem(), src, visited, depth+1, config); err != nil {
					return
				}
			} else {
				return ErrDifferentArgumentsTypes
			}
			break
		}

		//read note 空值处理
		if dst.IsNil() || overwrite {
			if dst.CanSet() && (overwrite || isEmptyValue(dst)) {
				dst.Set(src)
			}
			break
		}

		//read note 类型相同，进行deepMerge处理
		if dst.Elem().Kind() == src.Elem().Kind() {
			if err = deepMerge(dst.Elem(), src.Elem(), visited, depth+1, config); err != nil {
				return
			}
			break
		}
	default:
		//read note 默认情况下的处理，set值或者是直接赋值
		mustSet := (isEmptyValue(dst) || overwrite) && (!isEmptyValue(src) || overwriteWithEmptySrc)
		if mustSet {
			if dst.CanSet() {
				dst.Set(src)
			} else {
				dst = src
			}
		}
	}

	return
}

// Merge will fill any empty for value type attributes on the dst struct using corresponding
// src attributes if they themselves are not empty. dst and src must be valid same-type structs
// and dst must be a pointer to struct.
// It won't merge unexported (private) fields and will do recursively any exported field.
func Merge(dst, src interface{}, opts ...func(*Config)) error {
	//read note 这边的入参是被复制的结果结构体在前
	return merge(dst, src, opts...)
}

// MergeWithOverwrite will do the same as Merge except that non-empty dst attributes will be overridden by
// non-empty src attribute values.
// Deprecated: use Merge(…) with WithOverride
func MergeWithOverwrite(dst, src interface{}, opts ...func(*Config)) error {
	return merge(dst, src, append(opts, WithOverride)...)
}

//read note ------------------------------------
// 下面这几个就是Option进行设置Config的函数处理
// WithTransformers adds transformers to merge, allowing to customize the merging of some types.
func WithTransformers(transformers Transformers) func(*Config) {
	return func(config *Config) {
		config.Transformers = transformers
	}
}

// WithOverride will make merge override non-empty dst attributes with non-empty src attributes values.
func WithOverride(config *Config) {
	config.Overwrite = true
}

// WithOverwriteWithEmptyValue will make merge override non empty dst attributes with empty src attributes values.
func WithOverwriteWithEmptyValue(config *Config) {
	config.Overwrite = true
	config.overwriteWithEmptyValue = true
}

// WithOverrideEmptySlice will make merge override empty dst slice with empty src slice.
func WithOverrideEmptySlice(config *Config) {
	config.overwriteSliceWithEmptyValue = true
}

// WithAppendSlice will make merge append slices instead of overwriting it.
func WithAppendSlice(config *Config) {
	config.AppendSlice = true
}

// WithTypeCheck will make merge check types while overwriting it (must be used with WithOverride).
func WithTypeCheck(config *Config) {
	config.TypeCheck = true
}

// WithSliceDeepCopy will merge slice element one by one with Overwrite flag.
func WithSliceDeepCopy(config *Config) {
	config.sliceDeepCopy = true
	config.Overwrite = true
}

//read note ------------------------------------

func merge(dst, src interface{}, opts ...func(*Config)) error {
	//read note 判断指针类型和结构体类型是否正确
	if dst != nil && reflect.ValueOf(dst).Kind() != reflect.Ptr {
		return ErrNonPointerAgument
	}
	var (
		vDst, vSrc reflect.Value
		err        error
	)

	config := &Config{}

	//read note optional设置config
	for _, opt := range opts {
		opt(config)
	}

	if vDst, vSrc, err = resolveValues(dst, src); err != nil {
		return err
	}
	//read note 被merge的两个结构体类型必须一样，否则不能进行merge
	if vDst.Type() != vSrc.Type() {
		return ErrDifferentArgumentsTypes
	}
	//read note 进行deepMerge
	return deepMerge(vDst, vSrc, make(map[uintptr]*visit), 0, config)
}

// IsReflectNil is the reflect value provided nil
func isReflectNil(v reflect.Value) bool {
	k := v.Kind()
	switch k {
	case reflect.Interface, reflect.Slice, reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr:
		// Both interface and slice are nil if first word is 0.
		// Both are always bigger than a word; assume flagIndir.
		return v.IsNil()
	default:
		return false
	}
}

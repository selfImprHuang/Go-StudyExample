/*
 *  @Author : huangzj
 *  @Time : 2020/5/14 17:53
 *  @Description：
 */

package util

import (
	"Go-StudyExample/entity"
	"bytes"
	"encoding/gob"
	"encoding/json"
)

/*
 * @param dst 目标对象
 * @param src 源对象
 * 通过序列化的方式进行克隆
 */
func DeepCopyByGob(dst, src interface{}) error {
	var buffer bytes.Buffer
	if err := gob.NewEncoder(&buffer).Encode(src); err != nil {
		return err
	}

	return gob.NewDecoder(&buffer).Decode(dst)
}

func DeepCopyByJson(src interface{}) (*entity.SynthesisRule, error) {
	var dst = new(entity.SynthesisRule)
	b, err := json.Marshal(src)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, dst)
	return dst, err
}

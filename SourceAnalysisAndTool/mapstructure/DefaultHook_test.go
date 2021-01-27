/*
 *  @Author : huangzj
 *  @Time : 2021/1/25 16:09
 *  @Description：
 */

package example

import (
	"github.com/mitchellh/mapstructure"
	"testing"
)

func TestDefaultHook(t *testing.T) {
	//在 mitchellh/mapstructure 这个工程中提供了部分的默认Hook方法，主要作用是进行转换，具体可以查看源码
	mapstructure.StringToIPNetHookFunc()
}

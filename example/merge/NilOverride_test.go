/*
 *  @Author : huangzj
 *  @Time : 2021/2/3 9:50
 *  @Description：
 */

package merge

import (
	"fmt"
	"github.com/imdario/mergo"
	"log"
	"testing"
)

type redisConfig2 struct {
	Address string
	Port    int
	DBs     []int
}

var defaultConfig2 = redisConfig2{
	Address: "127.0.0.1",
	Port:    6381,
}

// 前提：需要配置 WithOverride 一起使用
// WithOverrideEmptySlice 		空数组覆盖
// WithOverwriteWithEmptyValue 	空字段、空数组覆盖
func TestNilOverride(t *testing.T) {
	var config redisConfig2
	config.DBs = []int{2, 3}

	if err := mergo.Merge(&config, defaultConfig2, mergo.WithOverride, mergo.WithOverrideEmptySlice); err != nil {
		log.Fatal(err)
	}

	fmt.Println("redis address: ", config.Address)
	fmt.Println("redis port: ", config.Port)
	fmt.Println("redis dbs: ", config.DBs)
}

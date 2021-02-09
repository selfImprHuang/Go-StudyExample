/*
 *  @Author : huangzj
 *  @Time : 2021/2/2 11:37
 *  @Description：
 */

package merge

import (
	"fmt"
	"log"
	"testing"

	"github.com/imdario/mergo"
)

type redisConfig struct {
	Address string
	Port    int
	DB      int
}

var defaultConfig = redisConfig{
	Address: "127.0.0.1",
	Port:    6381,
	DB:      1,
}

//基础使用，如果被merge的对象字段有值则不覆盖
func TestBaseUse(t *testing.T) {
	var config redisConfig

	if err := mergo.Merge(&config, defaultConfig); err != nil {
		log.Fatal(err)
	}

	fmt.Println("redis address: ", config.Address)
	fmt.Println("redis port: ", config.Port)
	fmt.Println("redis db: ", config.DB)

	var m = make(map[string]interface{})
	if err := mergo.Map(&m, defaultConfig); err != nil {
		log.Fatal(err)
	}

	fmt.Println(m)
}

//测试字段覆盖
func TestOverride(t *testing.T) {
	var config redisConfig

	if err := mergo.Merge(&config, defaultConfig); err != nil {
		log.Fatal(err)
	}

	fmt.Println("redis address: ", config.Address)
	fmt.Println("redis port: ", config.Port)
	fmt.Println("redis db: ", config.DB)

	var m = make(map[string]interface{})
	if err := mergo.Merge(&config, defaultConfig, mergo.WithOverride); err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)
}

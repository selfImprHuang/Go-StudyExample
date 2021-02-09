/*
 *  @Author : huangzj
 *  @Time : 2021/2/2 11:49
 *  @Description：
 */

package merge

import "testing"

import (
	"fmt"
	"log"

	"github.com/imdario/mergo"
)

type redisConfig1 struct {
	Address string
	Port    int
	DBs     []int
}

var defaultConfig1 = redisConfig1{
	Address: "127.0.0.1",
	Port:    6381,
	DBs:     []int{1},
}

func TestAppend(t *testing.T) {
	var config redisConfig1
	config.DBs = []int{2, 3}

	if err := mergo.Merge(&config, defaultConfig1, mergo.WithAppendSlice); err != nil {
		log.Fatal(err)
	}

	fmt.Println("redis address: ", config.Address)
	fmt.Println("redis port: ", config.Port)
	fmt.Println("redis dbs: ", config.DBs)
}

func TestNotAppend(t *testing.T) {
	var config redisConfig1
	config.DBs = []int{2, 3}

	if err := mergo.Merge(&config, defaultConfig1); err != nil {
		log.Fatal(err)
	}

	fmt.Println("redis address: ", config.Address)
	fmt.Println("redis port: ", config.Port)
	fmt.Println("redis dbs: ", config.DBs)
}

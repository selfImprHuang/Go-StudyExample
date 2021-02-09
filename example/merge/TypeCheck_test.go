/*
 *  @Author : huangzj
 *  @Time : 2021/2/3 9:59
 *  @Description：
 */

package merge

import (
	"fmt"
	"github.com/imdario/mergo"
	"log"
	"testing"
)

func TestWithoutTypeCheck(t *testing.T) {
	m1 := make(map[string]interface{})
	m1["dbs"] = []uint32{2, 3}

	m2 := make(map[string]interface{})
	m2["dbs"] = []int{1}

	if err := mergo.Map(&m1, &m2, mergo.WithOverride); err != nil {
		log.Fatal(err)
	}

	fmt.Println(m1)
}

func TestWithTypeCheck(t *testing.T) {
	m1 := make(map[string]interface{})
	m1["dbs"] = []uint32{2, 3}

	m2 := make(map[string]interface{})
	m2["dbs"] = []int{1}

	if err := mergo.Map(&m1, &m2, mergo.WithOverride, mergo.WithTypeCheck); err != nil {
		log.Fatal(err)
	}

	fmt.Println(m1)
}

/*
 *  @Author : huangzj
 *  @Time : 2020/7/6 11:06
 *  @Description：
 */

package example

func Assert(err error) {
	if err != nil {
		panic(err)
	}
}

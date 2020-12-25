/*
 *  @Author : huangzj
 *  @Time : 2020/12/25 15:23
 *  @Description：
 */

package carbon

import (
	"fmt"
	"github.com/uniplaces/carbon"
	"log"
	"testing"
	"time"
)

func TestTimeCreate(t *testing.T) {
	c, err := carbon.Create(2020, time.July, 24, 20, 0, 0, 0, "Japan")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The opening ceremony of next olympics will start at %s in Japan\n", c)

}

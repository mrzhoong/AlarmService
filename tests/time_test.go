package test

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeFormat(t *testing.T) {
	a := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(a)

	//	时间戳转换时间
	b := time.Unix(time.Now().Unix(), 0)
	fmt.Println(b)

	//	字符串转时间
	c, err := time.Parse("2006年01月02日 15:04:05", "2017-03-13 12:00")
	if err == nil {
		fmt.Println("Format:", c)
	}

}

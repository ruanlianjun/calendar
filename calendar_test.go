package calendar

import (
	"fmt"
	"github.com/ruanlianjun/calendar/calendar"
	"testing"
)

func TestCalendar(t *testing.T) {
	baZhi := calendar.NewBaZhi(false)
	_, _, _, _, i2, _, _ := baZhi.GetYun(0, 318416400)
	fmt.Println("yun",i2)
	getLiuNian := baZhi.GetLiuNian(2019)
	for i, m := range getLiuNian {
		fmt.Println("yun",i,m)
	}
	getLiuYue := baZhi.GetLiuYue(2020)

	for i, m := range getLiuYue {
		fmt.Println("getLiuYue",i,m)
	}

	liuShi, m := baZhi.GetLiuRISLiuShi(1)
	for i, m2 := range liuShi {
		fmt.Println(i, m2, m[i])
	}

}

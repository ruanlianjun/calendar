package calendar

import (
	"fmt"
	"github.com/ruanlianjun/calendar/calendar"
	"testing"
)

func TestCalendar(t *testing.T) {
	baZhi := calendar.NewBaZhi(false)
	_, _, _, _, i2, _, _ := baZhi.GetYun(0, 318416400000)
	fmt.Println("yun", i2)
	getLiuNian := baZhi.GetLiuNian(2019)
	for i, m := range getLiuNian {
		fmt.Println("yun", i, m)
	}
	getLiuYue := baZhi.GetLiuYue(2020)

	for i, m := range getLiuYue {
		fmt.Println("getLiuYue", i, m)
	}

	//2020-2-4 00:00:00 2020-3-4 00:00:00
	day, daymsg := baZhi.GetLiuRi(1580745600000, 1583251200000)
	for i, m := range day {
		fmt.Println("GetLiuRi ",i, m, daymsg[i])
	}

	liuShi, m := baZhi.GetLiuRISLiuShi(3)
	for i, m2 := range liuShi {
		fmt.Println("GetLiuRISLiuShi ",i, m2, m[i])
	}

}

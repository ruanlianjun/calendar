package main

import (
	"calendar/calendar"
	"fmt"
	"time"
)

func Yun(zwz bool, yy, mm, dd, hh, ii int) (map[string]int, string, []string, []string, []map[string]float64, string, map[string]map[string]map[string]string) {
	b := calendar.NewBaZhi(zwz)

	timeLocation, _ := time.LoadLocation("Asia/Shanghai")
	d := time.Date(yy, time.Month(mm), dd, hh, ii, 0, 0, timeLocation)
	tt := int(d.Unix())

	begtime, bazi, tg, dz, bigInfo, startDesc, jieqi := b.GetYun(0, tt)

	return begtime, bazi, tg, dz, bigInfo, startDesc, jieqi
}

func SoloarToLunar(yy, mm, dd, hh, ii, ss float64, zwz bool) (string, string) {
	b := calendar.NewBaZhi(zwz)
	f := b.Solar2Lunar(yy, mm, dd, hh, ii, ss)

	return fmt.Sprintf("%s-%s-%s", f["y"], f["m"], f["d"]), f["msg"]
}

func LunarToSolar(yy, mm, dd, r int, zwz bool) map[string]int {
	b := calendar.NewBaZhi(zwz)
	f := b.Lunar2Solar(yy, mm, dd, r)
	return f
}

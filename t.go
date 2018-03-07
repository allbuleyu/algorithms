package main

import (
	"github.com/allbuleyu/algorithms/leetcode"
	"fmt"
	"time"
)

func main() {
	//intSlice := make([]int, 0, 15)
	//for i := 1; i <= 5; i++ {
	//	for j := 0; j< i; j++{
	//		intSlice = append(intSlice, i)
	//	}
	//}
	//tt, _ := time.Parse("2006-01-02 15:04:05", "2018-02-01 00:00:00.000")
	//fmt.Println(findWeek(tt), findStartAndEnd(tt))
	//return
	s := "aaaaaa"
	t := []string{"aaa", "aaa"}
	//s := "barfoofoobarthefoobarman"
	//t := []string{"bar","foo","the"}
	//s := "wordgoodgoodgoodbestword"
	//t := []string{"word","good","best","good"}
	fmt.Println(leetcode.FindSubstring(s,t))
	return

	//intSlice := []int{4,5,5,6}
	intSlice := []int{2,3,1,2,4,3}
	fmt.Println(leetcode.MinSubArrayLen(7, intSlice))
	return
	//
	fmt.Println(leetcode.IsPossible(intSlice))

	//words := []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}
	//fmt.Println(leetcode.TopKFrequentStr(words, 4), "i" < "love")
}

func findWeek(d time.Time) string {
	year, week := d.ISOWeek()
	res := fmt.Sprintf("%d-%d", year,week)
	return res
}

func findStartAndEnd(d time.Time) string {
	var start, end time.Time
	switch d.Weekday() {
	case time.Monday:
		start = d.Add(0 * time.Hour * 24)
		end = d.Add(6 * time.Hour * 24)
	case time.Tuesday:
		start = d.Add(-1 * time.Hour * 24)
		end = d.Add(5 * time.Hour * 24)
	case time.Wednesday:
		start = d.Add(-2 * time.Hour * 24)
		end = d.Add(4 * time.Hour * 24)
	case time.Thursday:
		start = d.Add(-3 * time.Hour * 24)
		end = d.Add(3 * time.Hour * 24)
	case time.Friday:
		start = d.Add(-4 * time.Hour * 24)
		end = d.Add(2 * time.Hour * 24)
	case time.Saturday:
		start = d.Add(-5 * time.Hour * 24)
		end = d.Add(1 * time.Hour * 24)
	case time.Sunday:
		start = d.Add(-6 * time.Hour * 24)
		end = d.Add(0 * time.Hour * 24)
	}
	res := fmt.Sprintf("%s~%s", start.Format("01-02"), end.Format("01-02"))

	return res
}

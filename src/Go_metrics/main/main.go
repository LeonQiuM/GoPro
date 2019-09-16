package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/rcrowley/go-metrics"
)

func readLineFloat() map[float64]int {
	counts := make(map[string]int)
	//lineList := []float64{}
	lineList := make(map[float64]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		str := input.Text()
		fmt.Println(reflect.TypeOf(str))
		value, err := strconv.ParseFloat(str, 64)
		fmt.Println(value, err)
		if err == nil {
			strSlice := strings.Split(str, ".")
			dotCount := len(strSlice[len(strSlice)-1]) //  判断小数点后有多少位
			lineList[value] = dotCount
		} else {
			break
		}
		counts[str]++
	}
	return lineList
}

func readLineInt() []int64 {
	counts := make(map[string]int)
	lineList := []int64{}
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		str := input.Text()
		fmt.Println(reflect.TypeOf(str))
		value, err := strconv.ParseInt(str, 10, 64)
		fmt.Println(value, err)
		if err == nil {
			lineList = append(lineList, value)
		} else {
			break
		}
		counts[str]++
	}
	return lineList
}

func main() {
	typeArg := os.Args
	fmt.Println(typeArg)
	if typeArg[1] == "int" {
		lineList := readLineInt()
		s := metrics.NewExpDecaySample(1024, 0.015)
		h := metrics.NewHistogram(s)
		metrics.Register("baz", h)
		for _, v := range lineList {
			h.Update(int64(v))
		}

		ps := h.Percentiles([]float64{0.5, 0.75, 0.8, 0.85, 0.9, 0.95, 0.99, 0.999, 0.9999})
		fmt.Printf("* 数据中位值:\t%f\n* 数据75分位值:\t%f\n* 数据80分位值:\t%f\n* 数据85分位值:\t%f\n* 数据90分位值:\t%f\n* 数据95分位值:\t%f\n* 数据99分位值:\t%f\n* 数据99.9分位值:\t%f\n* 数据99.99分位值:\t%f\n", ps[0], ps[1], ps[2], ps[3], ps[4], ps[5], ps[6], ps[7], ps[8])
	} else if typeArg[1] == "float" {
		lineList := readLineFloat()
		s := metrics.NewExpDecaySample(1024, 0.015)
		h := metrics.NewHistogram(s)
		metrics.Register("baz", h)
		for _, v := range lineList {
			h.Update(int64(v))
		}

		ps := h.Percentiles([]float64{0.5, 0.75, 0.8, 0.85, 0.9, 0.95, 0.99, 0.999, 0.9999})
		fmt.Printf("* 数据中位值:\t%f\n* 数据75分位值:\t%f\n* 数据80分位值:\t%f\n* 数据85分位值:\t%f\n* 数据90分位值:\t%f\n* 数据95分位值:\t%f\n* 数据99分位值:\t%f\n* 数据99.9分位值:\t%f\n* 数据99.99分位值:\t%f\n", ps[0], ps[1], ps[2], ps[3], ps[4], ps[5], ps[6], ps[7], ps[8])

	}
}

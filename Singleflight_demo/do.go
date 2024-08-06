package main

import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"time"
)

func GetData(i int) string {

	fmt.Printf("查询\n")
	time.Sleep(2 * time.Second)
	return "successful"
}

func main1() {
	g := new(singleflight.Group)

	go func() {
		// 第一次调用
		v, err, shared := g.Do("GetData", func() (interface{}, error) {
			ret := GetData(1)
			return ret, nil
		})
		fmt.Printf("1st call v:%v err:%v shared:%v time:%v\n", v, err, shared, time.Now().Unix())
	}()

	// 休眠为了保证第一次调用已经开始，再开始第二次调用
	time.Sleep(time.Second)

	// 第二次调用
	v, err, shared := g.Do("getData", func() (interface{}, error) {
		ret := GetData(1)
		return ret, nil
	})
	fmt.Printf("2nd call v:%v err:%v shared:%v time:%v\n", v, err, shared, time.Now().Unix())
}

package main

import (
	"golang.org/x/sync/singleflight"
	"time"
)

func main() {

	g := new(singleflight.Group)
	// 从0s开始第一次执行，执行后，停5s
	go func() {
		g.Do("GetData", func() (interface{}, error) {
			v := GetData(1)
			time.Sleep(5 * time.Second)
			return v, nil
		})
	}()

	go func() {
		//从0s开始停，1s后forget
		time.Sleep(time.Second)
		g.Forget("GetData")
	}()

	//从0s开始停，保证上面是第一次执行
	time.Sleep(2 * time.Second)

	//从第2s开始执行，此时第一次执行还没结束，在第1s时已经forget了，所以可以执行
	g.Do("GetData", func() (interface{}, error) {
		v := GetData(1)
		return v, nil
	})

}

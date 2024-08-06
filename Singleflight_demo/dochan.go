package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/singleflight"
	"time"
)

// 使用DoChan 避免第一次请求就异常，导致后面的所有请求都异常

func getDataWithChannel(g *singleflight.Group, ctx context.Context) (string, error) {
	ch := g.DoChan("GetData", func() (interface{}, error) {
		ret := GetData(1)
		return ret, nil
	})

	// 通过select,返回超时就结束
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case ret := <-ch:
		return ret.Val.(string), nil
	}
}

func main2() {

	g := new(singleflight.Group)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	go func() {
		v, err := getDataWithChannel(g, ctx)
		fmt.Printf("1st call v:%v err:%v time:%v\n", v, err, time.Now().Unix())
	}()
	time.Sleep(time.Second)
	v, err := getDataWithChannel(g, ctx)
	fmt.Printf("2nd call v:%v err:%v time:%v\n", v, err, time.Now().Unix())
}

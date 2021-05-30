package main

import (
	"context"
	"fmt"
	"github.com/pjoc-team/healthz/registry"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))

	ctx, cancelFunc := context.WithCancel(context.TODO())
	defer cancelFunc()

	// 模拟kubelet循环打印health status
	tick := time.Tick(1 * time.Second)
	go func() {
		defer func() {
			status, err := registry.HealthStatus()
			if err != nil {
				panic(err.Error())
			}
			fmt.Println(status.String())
		}()
		for {
			select {
			case <-tick:
				status, err := registry.HealthStatus()
				if err != nil {
					panic(err.Error())
				}
				fmt.Println(status.String())
			case <-ctx.Done():
				return
			}
		}
	}()

	// 使用方调用component
	component, err := NewExampleComponent()
	if err != nil {
		time.Sleep(2 * time.Second)
		panic(err.Error())
	}
	component.Do()

}

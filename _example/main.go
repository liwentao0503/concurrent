package main

import (
	"context"
	"fmt"

	"github.com/wt1i/concurrent"
)

func main() {
	len := 10
	handle := make([]concurrent.ContextHandle, 0, len)
	resuit := make([]string, len)

	var funcTemp = func(ctx context.Context, i int) error {
		resuit[i] = fmt.Sprintf("value: %v", i)
		if i%3 == 0 {
			return fmt.Errorf("test err: %v", i)
		}
		return nil
	}

	for i := 0; i < len; i++ {
		handle = append(handle, funcTemp)
	}

	errList := concurrent.GoAndWait(context.Background(), handle)
	fmt.Println(resuit)
	fmt.Println(errList)
}

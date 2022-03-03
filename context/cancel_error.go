package main

import (
	"context"
	"fmt"
	"time"
)

const FormatLayout = "2006-01-02 15:04:05"

// cancel due to error
func main() {
	fmt.Println("main start", time.Now().Format(FormatLayout))

	// parent context
	ctx1, cancel1 := context.WithCancel(context.Background())

	go func() {
		err := action1(ctx1)
		if err != nil {
			fmt.Println("cancel1...", time.Now().Format(FormatLayout))
			cancel1()
		} else {
			fmt.Println("action1 done", time.Now().Format(FormatLayout))
		}
	}()

	go func() {
		// sub context
		ctx2, cancel2 := context.WithCancel(ctx1)
		err := action2(ctx2)
		if err != nil {
			cancel2()
		}
	}()

	action3(ctx1)
}

func action1(ctx context.Context) error {
	time.Sleep(time.Second)
	//return errors.New("action1 failed")
	return nil
}

func action2(ctx context.Context) error {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("action2 done", time.Now().Format(FormatLayout))
	case <-ctx.Done():
		fmt.Println("cancel action2", time.Now().Format(FormatLayout))
	}
	return nil
}

func action3(ctx context.Context) error {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("action3 done", time.Now().Format(FormatLayout))
	case <-ctx.Done():
		fmt.Println("cancel action3", time.Now().Format(FormatLayout))
	}
	return nil
}

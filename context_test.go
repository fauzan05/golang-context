package golangcontext

import (
	"context"
	"fmt"
	"runtime"
	// "sync"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	// contextA := context.Background()

	// contextB := context.WithValue(contextA, "b", "B")
	// contextC := context.WithValue(contextA, "c", "C")

	// contextD := context.WithValue(contextB, "d", "D")
	// contextE := context.WithValue(contextB, "e", "E")

	// contextF := context.WithValue(contextC, "f", "F")

	// fmt.Println(contextA)
	// fmt.Println(contextB)
	// fmt.Println(contextC)
	// fmt.Println(contextD)
	// fmt.Println(contextE)
	// fmt.Println(contextF)

	type favContextKey string // membuat alias dari string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k)                                         // return Go
	f(ctx, favContextKey("language"))                 // return Go juga
	fmt.Println(ctx.Value(favContextKey("language"))) // return go
}

// func Checking(ctx context.Context, destination chan int) bool {
// 	select {
// 	case <-ctx.Done():
// 		return false
// 	}
// }

func CreateCounter(ctx context.Context) chan int {
	// defer group.Done()
	destination := make(chan int)
	// group.Add(1)
	go func() {
		defer close(destination)
		counter := 1
		for {
			// check := Checking(ctx, destination)
			// if !check {
			// 	break
			// } else {
			// 	destination <- counter
			// 	counter++
			// }
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
			}

		}
	}()
	return destination
}

func TestContextWithCancel(t *testing.T) {
	// group := &sync.WaitGroup{}
	fmt.Println("Total Goroutine : ", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)
	for value := range destination {
		fmt.Println("Counter ke : ", value)
		if value == 10 {
			break
		}
		// group.Wait()
	}
	cancel()
	time.Sleep(3 * time.Second)
	fmt.Println("Total Goroutine : ", runtime.NumGoroutine())
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println(runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5 * time.Second)
	defer cancel()

	destination := CreateCounter(ctx)
	for value := range destination {
		fmt.Println("Counter ke : ", value)
	}

	fmt.Println(runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println(runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(2 * time.Second))
	defer cancel()

	destination := CreateCounter(ctx)
	for value := range destination {
		fmt.Println("Counter Ke : ",value)
	}
	fmt.Println(runtime.NumGoroutine())
}
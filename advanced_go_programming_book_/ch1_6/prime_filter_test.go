package ch1_6

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func GenerateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func PrimeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func TestPrimeFilter(t *testing.T) {
	ch := GenerateNatural()
	for i := 0; i < 100; i++ {
		prime := <-ch
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilter(ch, prime)
	}
}

func SafeGenerateNatural(ctx context.Context) chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("context done!")
				return
			case ch <- i:
			}
		}
	}()
	return ch
}

func SafePrimeFilter(ctx context.Context, in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				select {
				case <-ctx.Done():
					fmt.Println("context done!")
					return
				case out <- i:
				}
			}
		}
	}()
	return out
}

func TestSafePrimeFilter(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // avoid goroutine's leakage
	ch := SafeGenerateNatural(ctx)
	for i := 0; i < 100; i++ {
		prime := <-ch
		fmt.Println(prime)
		ch = SafePrimeFilter(ctx, ch, prime)
	}
}

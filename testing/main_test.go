package main

import "testing"

func TestSum(t *testing.T) {
	table := []struct {
		x int
		y int
		r int
	}{
		{2, 4, 6},
		{3, 5, 8},
		{4439, 5, 4444},
	}
	for _, item := range table {
		total := Sum(item.x, item.y)
		if total != item.r {
			t.Errorf("Got %d expected %d", total, item.r)
		}
	}
}

func TestGetMax(t *testing.T) {
	table := []struct {
		x int
		y int
		r int
	}{
		{4, 3, 4},
		{5, 1, 5},
		{4439, 5, 4439},
	}
	for _, item := range table {
		max := GetMax(item.x, item.y)
		if max != item.r {
			t.Errorf("Got %d expected %d", max, item.r)
		}
	}
}

func TestFibonacci(t *testing.T) {
	table := []struct {
		n int
		r int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}
	for _, item := range table {
		fib := Fibonacci(item.n)
		if fib != item.r {
			t.Errorf("Got %d expected %d", fib, item.r)
		}
	}
}

package stack

import "math"

type StockSpanner struct {
	stock []int
}

func StockSpannerConstructor() StockSpanner {
	return StockSpanner{[]int{0, math.MaxInt32}}
}

func (this *StockSpanner) Next(price int) int {
	skipped := 1
	for this.stock[len(this.stock)-1] <= price {
		skipped += this.stock[len(this.stock)-2]
		this.stock = this.stock[:len(this.stock)-2]
	}

	this.stock = append(this.stock, skipped, price)
	return skipped
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */

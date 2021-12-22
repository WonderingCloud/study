package movingaverage

type MovingAverage struct {
	Head int
	Tail int
	Nums []int
	Sum  int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{
		Head: 0,
		Tail: 0,
		Nums: make([]int, size+1),
		Sum:  0,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if this.Head == (this.Tail+1)%len(this.Nums) {
		this.Sum -= this.Nums[this.Head]
		this.Head = (this.Head + 1) % len(this.Nums)
	}
	this.Nums[this.Tail] = val
	this.Sum += val
	this.Tail = (this.Tail + 1) % len(this.Nums)

	cap := this.Tail - this.Head
	if cap < 0 {
		cap = cap + len(this.Nums)
	}
	return float64(this.Sum) / float64(cap)
}

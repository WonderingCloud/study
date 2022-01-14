package minstack

type MinStack struct {
	main []int
	sub  []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (ms *MinStack) Push(val int) {
	ms.main = append(ms.main, val)
	if len(ms.sub) == 0 || val < ms.sub[len(ms.sub)-1] {
		ms.sub = append(ms.sub, val)
	} else {
		ms.sub = append(ms.sub, ms.sub[len(ms.sub)-1])
	}
}

func (ms *MinStack) Pop() {
	ms.main = ms.main[:len(ms.main)-1]
	ms.sub = ms.sub[:len(ms.sub)-1]
}

func (ms *MinStack) Top() int {
	return ms.main[len(ms.main)-1]
}

func (ms *MinStack) GetMin() int {
	return ms.sub[len(ms.sub)-1]
}

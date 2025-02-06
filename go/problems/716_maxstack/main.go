/*
ProblemId: s2296
ProblemTitle: Design a text editor
ProblemLink: https://leetcode.com/problems/merge-two-sorted-lists/
*/
package main

type MaxStack struct {
	maxvals []int
	values  []int
}

func Constructor() MaxStack {
	return MaxStack{}
}

func (this *MaxStack) Push(x int) {
	if len(this.maxvals) == 0 || x >= this.maxvals[len(this.maxvals)-1] {
		this.maxvals = append(this.maxvals, x)
	}
	this.values = append(this.values, x)
}

func (this *MaxStack) Pop() int {
	topval := this.Top()
	mtopval := this.PeekMax()
	this.values = this.values[:len(this.values)-1]
	if topval == mtopval {
		// prop from top stack too
		this.maxvals = this.maxvals[:len(this.maxvals)-1]
	}
	return topval
}

func (this *MaxStack) Top() int {
	return this.values[len(this.values)-1]
}

func (this *MaxStack) PeekMax() int {
	return this.maxvals[len(this.maxvals)-1]
}

func (this *MaxStack) PopMax() int {
	mtop := this.PeekMax()
	for this.Top() != mtop {
		this.Pop()
	}
	// And the max too
	this.Pop()
	return mtop
}

/**
 * Your MaxStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.PeekMax();
 * param_5 := obj.PopMax();
 */

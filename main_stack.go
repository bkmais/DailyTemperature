package main

import "fmt"

func main() {
	Temps := []int{73, 74, 75, 71, 69, 72, 76, 73}

	returnedVals := dailyTemperatures(Temps)

	fmt.Printf("%v", returnedVals)
}

type myStack struct {
	retStack []int
	retInd   []int
}

func Constructor() myStack {
	return myStack{}
}

func (this *myStack) Push(val int, ind int) {
	this.retStack = append(this.retStack, val)
	this.retInd = append(this.retInd, ind)
}

func (this *myStack) Pop() {
	if len(this.retStack) > 0 && len(this.retInd) > 0 {
		this.retStack = this.retStack[:len(this.retStack)-1]
		this.retInd = this.retInd[:len(this.retInd)-1]
	}
}

func (this *myStack) Top() []int {

	retarr := make([]int, 2)
	if len(this.retStack) > 0 && len(this.retInd) > 0 {
		retarr[0] = this.retInd[len(this.retInd)-1]
		retarr[1] = this.retStack[len(this.retStack)-1]
	} else {
		retarr[0] = 0
		retarr[1] = 0
	}
	return retarr

}

func (this *myStack) IsEmpty() bool {
	return len(this.retStack) == 0 && len(this.retInd) == 0
}

func dailyTemperatures(T []int) []int {

	// will hold the values of number of days
	ret := make([]int, len(T))
	// Will hold the largest temp as we iterate from high to low
	stack := new(myStack)

	for i := len(T) - 1; i >= 0; i-- {
		top := stack.Top()
		//While the stack is not empty & current temp is grt thn top in stack
		for !stack.IsEmpty() && T[i] >= top[1] {
			stack.Pop()
			top = stack.Top()
		}
		if stack.IsEmpty() {
			ret[i] = 0
			stack.Push(T[i], i)
		} else {
			ret[i] = top[0] - i
			stack.Push(T[i], i)
		}
	}
	return ret
}

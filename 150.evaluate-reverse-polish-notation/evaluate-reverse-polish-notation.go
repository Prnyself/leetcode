package main

import "strconv"

type myStack struct {
	list []int
}

func Conductor() myStack {
	return myStack{
		list: make([]int, 0),
	}
}

func evalRPN(tokens []string) int {
	tokenMap := map[string]string{
		"+": "+",
		"-": "-",
		"*": "*",
		"/": "/",
	}

	stack := Conductor()

	for _, v := range tokens {
		if _, ok := tokenMap[v]; !ok {
			tmp, _ := strconv.Atoi(v)
			stack.push(tmp)
		} else {
			num2 := stack.pop()
			num1 := stack.pop()
			stack.push(cal(num1, num2, v))
		}
	}

	return stack.pop()
}

func (this *myStack) push(v int) {
	this.list = append(this.list, v)
}

func (this *myStack) pop() int {
	n := len(this.list)
	res := this.list[n-1]
	this.list = this.list[:n-1]
	return res
}

func cal(num1 int, num2 int, sign string) int {
	switch sign {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	default:
		return 0
	}
}

func main() {

}
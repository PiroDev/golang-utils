package calc

import (
	"container/list"
	"errors"
	"log"
	"strconv"
	"strings"
)

type Stack struct {
	data *list.List
}

func NewStack() *Stack {
	return &Stack{data: list.New()}
}

func (stack *Stack) Push(value interface{}) {
	stack.data.PushBack(value)
}

func (stack *Stack) Peek() interface{} {
	if stack.IsEmpty() {
		log.Fatal("stack underflow")
	}

	return stack.data.Back().Value
}

func (stack *Stack) Pop() interface{} {
	defer func() {
		stack.data.Remove(stack.data.Back())
	}()

	return stack.Peek()
}

func (stack *Stack) Len() int {
	return stack.data.Len()
}

func (stack *Stack) IsEmpty() bool {
	return stack.Len() == 0
}

type token struct {
	Value interface{}
	IsNum bool
}

func isOper(symbol byte) bool {
	return strings.Contains("+-*/()", string(symbol))
}

func splitToTokens(expr string) ([]token, error) {
	expr = strings.ReplaceAll(expr, " ", "")
	tokens := make([]token, 0)

	handle := func(err error) ([]token, error) {
		return nil, errors.New("error while parsing expression")
	}

	start := 0
	end := 0

	for i := 0; i < len(expr); i++ {
		if isOper(expr[i]) {
			if start != end {
				num, err := strconv.ParseFloat(expr[start:end], 64)
				if err != nil {
					handle(err)
				}

				tokens = append(tokens, token{Value: num, IsNum: true})
			}
			tokens = append(tokens, token{Value: expr[i], IsNum: false})
			start = end + 1
			end = start
		} else {
			end++
		}
	}

	if start != end {
		num, err := strconv.ParseFloat(expr[start:end], 64)

		if err != nil {
			handle(err)
		}

		tokens = append(tokens, token{Value: num, IsNum: true})
	}

	return tokens, nil
}

func processOper(nums, opers *Stack) {
	rightNum := nums.Pop().(float64)
	leftNum := nums.Pop().(float64)

	var resNum float64
	currentOper := opers.Pop().(byte)

	switch currentOper {
	case '+':
		resNum = leftNum + rightNum
	case '-':
		resNum = leftNum - rightNum
	case '*':
		resNum = leftNum * rightNum
	case '/':
		resNum = leftNum / rightNum
	}

	nums.Push(resNum)
}

func Calculate(expr string) (float64, error) {
	tokens, err := splitToTokens(expr)

	if err != nil {
		return 0, err
	}

	opers := NewStack()
	nums := NewStack()

	priority := map[byte]int{
		'(': 0,
		')': 0,
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
	}

	for _, token := range tokens {
		if token.IsNum {
			num := token.Value.(float64)
			nums.Push(num)
		} else {
			newOper := token.Value.(byte)

			var currentOper byte
			if !opers.IsEmpty() {
				currentOper = opers.Peek().(byte)
			}

			if opers.IsEmpty() || priority[newOper] > priority[currentOper] {
				opers.Push(newOper)
			} else {
				if newOper == '(' {
					opers.Push(newOper)
				} else if newOper == ')' {
					for currentOper != '(' {
						processOper(nums, opers)
						currentOper = opers.Peek().(byte)
					}
					opers.Pop()
				} else {
					processOper(nums, opers)
					opers.Push(newOper)
				}
			}
		}

	}

	for !opers.IsEmpty() {
		if oper := opers.Peek().(byte); oper == '(' {
			opers.Pop()
		} else {
			processOper(nums, opers)
		}
	}
	resNum := nums.Pop().(float64)

	return resNum, nil
}

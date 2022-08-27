package sol

import "fmt"

func RunTest(commands []string, payload [][]int) []string {
	length := len(payload)
	result := []string{"null"}
	var minStack MinStack = Constructor()
	for idx := 1; idx < length; idx++ {
		command := commands[idx]
		switch command {
		case "push":
			data := payload[idx][0]
			minStack.Push(data)
			result = append(result, "null")
		case "pop":
			minStack.Pop()
			result = append(result, "null")
		case "top":
			top := minStack.Top()
			result = append(result, fmt.Sprintf("%d", top))
		case "getMin":
			min := minStack.GetMin()
			result = append(result, fmt.Sprintf("%d", min))
		}
	}
	return result
}

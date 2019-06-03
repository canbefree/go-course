package pipe

type Stack func(int) StackPossible
type StackPossible interface{}

type PItem func(int, Stack) StackPossible

//Carry 返回值作为下次调用的参数
//@params stack
//@params line
func Carry(stack Stack, p PItem) StackPossible {
	return func(num int) StackPossible {
		return p(num, stack)
	}
}

func add(num int, stack Stack) StackPossible {
	return stack(num)
}

func minutes(num int, stack Stack) StackPossible {
	return stack(num)
}

func initStack() Stack {
	return func(num int) StackPossible {
		return num
	}
}

func init() {
	stack := initStack()
	Carry(stack, add)

	f := Carry(Carry(stack, add), minutes)
	f(50)
}

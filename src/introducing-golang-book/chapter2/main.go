package main

import "fmt"

func main() {
	math_operations_examples()
	strings_operations_examples()
	booleans_operations_examples()
	chapter2_exercise1()
	chapter2_exercise2()
	chapter2_exercise3()
	chapter2_exercise4()
	chapter2_exercise5()
}

func math_operations_examples() {
	// Math operations in golang
	fmt.Println("Math operations:\n")
	fmt.Println("\t1 + 1 =", 1+1)
	fmt.Println("\t1 + 1 =", 1+1.0)
	fmt.Println("\t1 + 1 =", 1.0+1.0)
	fmt.Println("\t1 - 1 =", 1-1)
	fmt.Println("\t1 * 1 =", 1*1)
	fmt.Println("\t1 / 1 =", 1/1)
	fmt.Println("\t1 % 1 =", 1%1)
}

func strings_operations_examples() {
	// Strings in golang
	fmt.Println("\nStrings:\n")
	fmt.Println("\tString:", "String")
	fmt.Println("\tString len:", len("String"))
	// This show the value that correspond to the caracter
	fmt.Println("\tString first caracter:", "String"[0])
	fmt.Println("\tString 'sum':", "String1"+"String2")
}

func booleans_operations_examples() {
	// Booleans in golang
	fmt.Println("\nBooleans logic:\n")
	// AND OPERATOR
	fmt.Println("\n\tAND OPERATOR:")
	fmt.Println("\ttrue && false:", true && false)
	fmt.Println("\ttrue && true:", true && true)
	// OR OPERATOR
	fmt.Println("\n\tOR OPERATOR:")
	fmt.Println("\tfalse || false:", false || false)
	fmt.Println("\ttrue || false:", true || false)
	fmt.Println("\ttrue || true:", true || true)
	// NOT OPERATOR
	fmt.Println("\n\tNOT OPERATOR:")
	fmt.Println("\t!false:", !false)
	fmt.Println("\t!true:", !true)
}

func chapter2_exercise1() {
	// Questions
	fmt.Println("\n1. How are integers stored on a computer?")
	// Answer
	fmt.Println("Ans.: Integers are stored in a computer in binary. Although, is important to notice that integers have types, lens and other caracteristcs.")
}

func chapter2_exercise2() {
	// Questions
	fmt.Println("\n2. We know that (in base 10) the largest one-digit number is 9 and the largest two-digit number is 99. Given that in binary the largest two-digit number is 11 (3), the largest three-digit number is 111 (7) and the largest four-digit number is 1111 (15), what's the largest eight-digit number? (Hint: 101-1 = 9 and 102-1 = 99).")
	// Answer
	fmt.Println("Ans.: The largest eight-digit number is 11111111. And the method to know this binary number in 10 base is: \n2⁸-1 = 1*2⁷ + 1*2⁶ + 1*2⁵ + 1*2⁴ + 1*2³ + 1*2² + 1*2¹ + 2⁰ =", 1*128+1*64+1*32+1*16+1*8+1*4+1*2+1*1)
}

func chapter2_exercise3() {
	// Questions
	fmt.Println("\n3. Although overpowered for the task, you can use Go as a calculator. Write a pro-gram that computes 32,132 x 42,452 and prints it to the terminal (use the * operator for multiplication).")
	// Answer
	fmt.Println("Ans.: 32,132 * 42,452 =", 32132*42452)
}

func chapter2_exercise4() {
	// Questions
	fmt.Println("\n4. What is a string? How do you find its length?")
	// Answer
	fmt.Println("Ans.: String is a sequence of characters. To know the lenght of a string the method len(string) is needed. For example: lenght = len('string') =", len("string"))
}

func chapter2_exercise5() {
	// Questions
	fmt.Println("\n5. What's the value of the expression (true && false) || (false && true) || ! (false && false)?")
	// Answer
	fmt.Println("Ans.: The value of the expression is:", (true && false) || (false && true) || !(false && false))
}

// package declaration - must be in the beginning
package main

// to include code from other packages - ALWAYS BETWEEN DOUBLE COTES
import "fmt"

// function declaration
func main() {
	// print line = Hello World
	fmt.Println("Hello World")
	chapter1_exercise1()
	chapter1_exercise2()
	chapter1_exercise3()
	chapter1_exercise4()
	chapter1_exercise5()
}

/*
The name main is special because it’s the function
that gets called when you execute the program
*/

func chapter1_exercise1() {
	// Questions
	fmt.Println("\n1. What is whitespace?")
	// Answer
	fmt.Println("Ans.: Whitespaces are the represetation of newlines, spaces and tabs. Which we cannot see.")
}

func chapter1_exercise2() {
	// Questions
	fmt.Println("\n2. What is a comment? What are the two ways of writing a comment?")
	// Answer
	fmt.Println("Ans.: Comment is a way to make the command be ignored by the program. The two ways to write a comment are: //SINGLE LINE COMMENT /* MULTI LINE COMMENT */")
}

func chapter1_exercise3() {
	// Questions
	fmt.Println("\n3. Our program began with package main. What would the files in the fmt package begin with?")
	// Answer
	fmt.Println("Ans.: Once the files in the fmt package use Golang, they must begin with package main.")
}

func chapter1_exercise4() {
	// Questions
	fmt.Println("\n4. We used the Println function defined in the fmt package. If you wanted to use the Exit function from the os package, what would you need to do?")
	// Answer
	fmt.Println("Ans.: I would need to import the os package, then use the command os.Exit()")
}

func chapter1_exercise5() {
	// Questions
	fmt.Println("\n5. Modify the program we wrote so that instead of printing Hello, World it prints Hello, my name is followed by your name.")
	// Answer
	fmt.Println("Ans.: Hello, my name is Jean Carlo")
}

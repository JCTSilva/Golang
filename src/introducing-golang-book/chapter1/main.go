/*
Create a new executable program that references the fmt library
and contains one function called main. That function takes no
arguments and doesn’t return anything. It accesses the Println
function contained inside of the fmt package and invokes it using
one argument—the string Hello, World.
*/

// package declaration - must be in the beginning
package main

// to include code from other packages - ALWAYS BETWEEN DOUBLE COTES
import "fmt"

// function declaration
func main() {
	// print line = Hello World
	fmt.Println("Hello World")
}

/*
The name main is special because it’s the function
that gets called when you execute the program
*/

// EXERCISES

/*
1. What is whitespace?
R. Whitespaces are the represetation of newlines,
spaces and tabs. Which we cannot see.
*/

/*
2. What is a comment? What are the two ways of writing a comment??
R. Comment is a way to make the command be ignored by the program.
The two ways to write a comment are: //COMMENT /* COMMENT */

/*
3. Our program began with package main. What would the files in the
fmt package begin with?
R. The files in the package must begin with package main too.
*/

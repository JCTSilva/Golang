/*
4. We used the Println function defined in the fmt package. If you
wanted to use the Exit function from the os package, what would you
need to do?
R.
import os
Exit(0)
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	// print my name
	fmt.Println("Hello, my name is Jean")
	// exit status
	os.Exit(0)
}

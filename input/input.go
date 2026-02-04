package input

import (
	"fmt"
	"bufio"
	"os"
)

//Function for user input
func Input() string {
	fmt.Println("Input text:")
    reader := bufio.NewReader(os.Stdin)
    line, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Mistake")
    }
    return line
}

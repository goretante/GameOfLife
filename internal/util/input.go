package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetIntFromUser(prompt string) int {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if val, err := strconv.Atoi(input); err == nil {
			return val
		} else {
			fmt.Println("Invalid input. Please enter an integer.")
		}
	}
}

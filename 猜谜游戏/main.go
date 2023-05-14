package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// 生成随机数
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	// fmt.Println("The secret number is", secretNumber)

	// 读取用户输入
	fmt.Println("Please input your guess")
	reader := bufio.NewReader(os.Stdin)
	// 实现循环
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An err occured while reading input. Please try again", err)
			continue
		}
		input = strings.TrimSuffix(input, "\n")

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer Value")
			continue
		}
		fmt.Println("You guess is", guess)

		// 逻辑判断
		if guess > secretNumber {
			fmt.Println("Your guess is bigger rhan the secret number. Please try again")
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number. Please try again")
		} else {
			fmt.Println("Correct, you Legend!")
			break
		}
	}
}

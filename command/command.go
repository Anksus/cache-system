package command

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func StartCacheSystem(ctx *cli.Context) error {

	fmt.Println(" ------------- starting a cache system -------------")
	fmt.Println("please select a eviction policy")
	fmt.Println("1. LRU")
	fmt.Println("2. LFU")

	reader := bufio.NewReader(os.Stdin)

	selectedPolicy, err := reader.ReadString('\n')
	selectedPolicy = selectedPolicy[:len(selectedPolicy)-1]
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("please enter maximum limit of caching")

	limit, err := reader.ReadString('\n')
	limit = limit[:len(limit)-1]

	// 							limit = limit[:len(limit)-1]

	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("------------------------------------------------------")
	fmt.Println(fmt.Sprintf("selectedPolicy: %s, limit: %s", selectedPolicy, limit))
	fmt.Println("------------------------------------------------------")
	fmt.Println()
	fmt.Println()
	for {
		fmt.Println("Enter key and value like this key:value")
		input, err := reader.ReadString('\n')
		input = input[:len(input)-1]
		if err != nil {
			fmt.Println(err)
			return err
		}
		parts := strings.Split(input, ":")
		if len(parts) != 2 {
			fmt.Println("Invalid input format. Expected 'key:value'.")
			return errors.New("invalid input format. Expected 'key:value'")
		}
		key := parts[0]
		value := parts[1]

	}
}

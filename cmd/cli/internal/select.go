package internal

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetUserSelection(prompt string, max int) (int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)

	number, err := reader.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("failed to read input: %w", err)
	}

	selection, err := strconv.Atoi(strings.TrimSpace(number))
	if err != nil || selection < 1 || selection > max {
		return 0, fmt.Errorf("invalid selection, please rerun the command")
	}

	return selection, nil
}

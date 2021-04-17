package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	operation := scanner.Text()

	valores := strings.Split(operation, "+")
	operation_int, _ := strconv.Atoi(valores[0])
	operation_int_2, _ := strconv.Atoi(valores[1])
	fmt.Println(operation_int + operation_int_2)
}

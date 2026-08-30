package main

import (
	"fmt"

	"github.com/mrruke12/lms/internal/domains/attempt"
)

func main() {
	a := attempt.Attempt{}

	fmt.Println(a.SetStatus(attempt.Status.Active))
	fmt.Println(a.SetStatus("123"))
}

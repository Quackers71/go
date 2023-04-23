package main

import (
	"fmt"
	"github.com/Quackers71/go/go-crash-course-travmed/03_packages/strutil"
	"math"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))

	fmt.Println(strutil.Reverse("olleH"))
}

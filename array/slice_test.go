package array

import (
	"fmt"
	"strings"
	"testing"
)

func TestSlice2(t *testing.T) {
	scores := []int{1,2,3,4,5}
	slice := scores[2:4]
	slice[0] = 999
	fmt.Println(scores)
	fmt.Println(slice)

}

func TestSlice3(t *testing.T) {
	haystack := "the spice must flow";
	fmt.Print(strings.Index(haystack[5:], "p"))
}




package scanf

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func TestScan(t *testing.T) {


	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()
		t.Log(strings.ToUpper(input))
	}


}

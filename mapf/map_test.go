package mapf

import (
	"fmt"
	"sort"
	"testing"
)

func TestMapFunction(t *testing.T) {

	m := map[string]func(string, string){

	}
	t.Log(m)

	f := make(map[int]func())

	f[1] = func() {
		t.Log("f1")
	}

	f[1]()

	t.Log(f)

	i := make(map[int]interface{})

	i[1] = func(a, b, c int) {
		t.Log(a, b, c)
	}

	i[1].(func(a, b, c int))(1, 2, 3)

	i[2] = 3
	t.Log(i[2])

	i[3] = ""

	t.Log(i[3])
	i[4] = `
"a":"valueA",
"b":"valueB",
"c":"valueC"
}`

	t.Log(i[4])

}

func Test3(t *testing.T) {
	m := map[int]string{
		1:    "1231",
		1231: "1232",
		65:   "1233",
		13:   "1234",
	}
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", m[k])
	}
}

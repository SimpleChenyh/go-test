package reflectionf

import (
	"reflect"
	"testing"
)

type User struct {
	Name string
	Age  int
}

func TestRef(t *testing.T) {
	user := User{Name: "haha", Age: 123}
	PrintFieldName(user)
}

func PrintFieldName(i interface{}) {

	of := reflect.TypeOf(i)
	//log.Print(of)
	kind := of.Kind()
	//log.Print(kind)
	var strt reflect.Type
	switch kind {
	case reflect.Ptr:
		strt = of.Elem()
	default:
		strt = of
	}
	fieldCount := strt.NumField()
	for i := 0; i < fieldCount; i++ {
		name := strt.Field(i).Name
		_, _ = strt.FieldByName(name)

	}
}

func BenchmarkRef(b *testing.B) {

	for i := 0; i < b.N; i++ {
		PrintFieldName(User{"123", 1})
	}


}

func BenchmarkDirection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		user := User{"123", 1}
		_ = user.Age
		_ = user.Name
	}

}

func LoadByName()  {
		
}
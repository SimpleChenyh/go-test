package interfacef

import (
	"fmt"
	"log"
	"testing"
)

type Live interface {
	Eat() string
}
type Cat struct {
  Name string
}

type Dog struct {
	Name string
}

type AnimalEat interface {
	Live
}






func (animal *Cat) Eat() string{
	return fmt.Sprintf("The cat called %s likes to eat fish", animal.Name)
}

func (animal *Dog) Eat() string{
	return fmt.Sprintf("The dog called %s likes to eat fish", animal.Name)
}

func TestInterfaceEat(t *testing.T) {


	//Live(Cat{"Tom"})

	c := &Cat{Name: "Tom"}
	prepareFood(c)



	d := &Dog{Name: "Feeee"}
	prepareFood(d)

}


func prepareFood(live Live){
	log.Printf(live.Eat())
}



package funcf

import (
	"log"
	"testing"
	"time"
)

// Pass a func as param
func logFunc(a func(i int)) func(i int) {
	log.Print("start ...")
	a(1)
	log.Print("end ...")
	return a

}

func sleepf(i int) {
	time.Sleep(time.Duration(i) * time.Second)
}

func TestFuncParam(t *testing.T) {

	logFunc(sleepf)

	log.Print("-----------------------------")
	b := func(i int) {
		log.Print("************** b func")
	}
	logFunc(b)
}

func logTimeConsuming(f func(str string)) func(str string) {
	return func(str string) {
		start := time.Now()
		f(str)
		log.Print("consuming : ", time.Since(start).Seconds(), " Seconds")
	}
}

func TestTimeConsuming(t *testing.T) {

	f1 := func(str string) {
		log.Print("Hi this is the func param and return value test case", str)
		time.Sleep(time.Duration(len(str)) * time.Second)
	}

	consumingF1 := logTimeConsuming(f1)

	consumingF1("Ace never die")

}

func Sum(x ...int) int {
	result := 0
	for _, v := range x {
		result += v
	}
	return result

}

func TestMultiParam(t *testing.T) {

	t.Log(Sum(1, 2, 3, 4, 54, 3, 234, 1))
	t.Log(Sum(1, 2, 3, 4, 54, 3, 234, 12313131))
}

func deferFunc(a int) {

	defer log.Print(Sum(1, 2, 3, 4, 54, 3, 234, 1))
	defer log.Print("---------_+++++++______+_")
	log.Print(2)
	log.Print(a)
	panic("Exit ....")
}

func TestDefer(t *testing.T) {

	deferFunc(123)

}

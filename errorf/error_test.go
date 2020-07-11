package errorf

import (
	"errors"
	"fmt"
	"testing"
)

func Febonacci(n int) ([]int, error) {

	if n < 0 {
		return nil, errors.New("n can not less than 0")
	}

	rlt := []int{1, 1}

	for i := 2; i < n; i++ {
		rlt = append(rlt, rlt[i-1]+rlt[i-2])
	}

	return rlt, nil

}




func TestErr(t *testing.T) {

	var e error

	if v, e := Febonacci(-100); e != nil {
		t.Log("Error -->", e)
	} else {
		t.Log("result -->", v)
	}

	if nil != e {
		t.Log("Error is not nil")
	} else {
		t.Log("Error is nil")
	}

}


func threeReturnVal() (a int,b int,c int,d int){
	return 1231,b,123123,d
}


func threeReturnVal2() (a int,b int,c int,d int){
	d = 1111
	a = 4444
	return d,b,123123,a
}

func TestThreeReVal(t *testing.T) {
	t.Log(threeReturnVal())
}


func TestThreeReVal2(t *testing.T) {
	t.Log(threeReturnVal2())

}



var CommonError = errors.New("Common error")
var NotLogin = errors.New("Not login error")


func TestError(t *testing.T) {

	a := 100


	if(a < 10){
		panic(CommonError)
	}
	panic(NotLogin)

fmt.Print("hah")
}












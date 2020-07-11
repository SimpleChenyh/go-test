package condition

import "testing"

func TestSwitch1(t *testing.T) {

	for i := 0; i < 111; i++ {

		switch i {
		case 0,1, 2, 3, 4:
			t.Log("lte 5")
		case 5,6,7,8:
			t.Log("gte 5 and lte 8")
		default:
			t.Log("gt 8")
		}

	}

}



func TestSwitch2(t *testing.T){

	for i := 0;i < 10; i++ {

		switch {
		case i%2 == 0:
			t.Log("Odd")
		default:
			t.Log("Even")



		}


	}
}
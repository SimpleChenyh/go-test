package objectf

import (
	"log"
	"testing"
	"unsafe"
)

type User struct {
	id string
	b  string
	c  string
}

func TestObject(t *testing.T) {


	u1 := User{"a", "v", "c"}
	t.Log(u1)

	u2 := User{id: "a", c: "c",b: "v"}

	t.Log(u2)

	t.Log(u1 == u2)



	u3 := new(User)

	u3.id = "123"
	u3.b = "a"
	u3.c = "c"
	t.Log(u3)

	t.Logf("u1 type %[1]T %[1]T %[1]T ",u1)
	t.Logf("u2 type %T",u2)
	t.Logf("u3 type %T",u3)
}


func (user *User) doRightThings(){
	log.Printf("%s %s %s",user.id,user.b,user.c)
	log.Printf("func -- %x \n",unsafe.Pointer(&user.id))
	user.b = "Change value"
}


func TestPointerFuncParam(t *testing.T) {

	u1 := User{"a", "v", "c"}
	t.Logf("%x \n",unsafe.Pointer(&u1.id))
	u1.doRightThings()
	t.Logf("%x \n",unsafe.Pointer(&u1.id))
	u1.doRightThings()
}


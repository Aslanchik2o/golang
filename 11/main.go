// структуры
package main

import (
	"fmt"
)

type User struct {
	name string
	age int
	sex int
	weight string
	height int
}

func (u User) Printuserinfo(){
	fmt.Println(u.name, u.age, u.sex, u.weight,  u.height)
	}

func NewUser(name, sex, string, age, weight, height int) User {
	  return User{
		name: name,
		sex: sex,
		age: age,
		weight: weight,
		height: height,

	}
}

func main() {
	user1 := NewUser( "Vasya",  "Male", 23, 75, 186)
	user2 := User{ "Petya", 31, "Male",  79, 197}


user1.Printuserinfo()
user2.Printuserinfo()	



		
	
	//fmt.Printf("%+v\n",user1)
	//fmt.Printf("%+v\n",user2)
	//fmt.Println(user1.name, user1.age)
	//fmt.Println(user2.name, user2)
	
}






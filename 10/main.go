package main

import "fmt"

// тема:Мапы
 func main(){
	users := map[string] int{
		"Vana": 15,
		"Yana": 23, 
		"Loxxxx": 48,
	}


//	delete(users, "Sergay")
	
//	users["Sergay"] = 21

fmt.Println(len(users))
//	for key, value := range users{
//		fmt.Println(key,value)
//	}
	



	//age, exists := users["Vana"]
	//fmt.Println(age,exists)

	//age, exsits := users["Vana"]
  //if exsits
	//	fmt.Println("Vana",age)
//} else{
//		fmt.Println("Vana нет в списке")
//	}

 }
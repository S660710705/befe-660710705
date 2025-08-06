package main
import ("fmt")

//var email string ="Nakfon_n@silpakorn.edu" ประกาศ global ใช้ <<<<<<

func main (){
	//var name string = "Napapat"
	var age int = 20

	email :="Nakfon_n@silpakorn.edu"
	gpa := 3.50

	firstname, lastname := "Napapat", "Nakfon"

	fmt.Printf("Name %s %s, age %d, email %s, gpa %.2f\n", firstname, lastname, age ,email, gpa)


}
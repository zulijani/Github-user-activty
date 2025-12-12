package main

import (
	"fmt"
	"os"
)

func main(){
	var argument []string
	argument = append(argument, os.Args...)
	fmt.Println(argument)
	
	
	if len(argument) == 2 {
		fmt.Println(argument[1])
		get_github_api(argument[1])
	} else {
		fmt.Println("program expects one argument")
	}
	
}

func get_github_api(username string){

}
package main

import "fmt"

func main(){
	m := make(map[string]string)

	//setting elment

	m['name'] = "golang"

	//get an elemt

	fmt.Println(m["name"])
}
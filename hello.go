package main

import "fmt"

func main(){
	fmt.Println(Hello("Martin"))
}

func Hello(name string) string {
	return "Hello " + name
}
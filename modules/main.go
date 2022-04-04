package main

import (
	"github.com/donvito/hellomod"
	hellomod2 "github.com/donvito/hellomod/v2"
)

func main() {
	//hellomod.Println("Hello")
	hellomod.SayHello()
	hellomod2.SayHello("World..!!\n")
}

package main

// for test
import "fmt"
import "time"

type singleton struct {
	Name string
}

var obj *singleton = nil

func GetSingleInstance() *singleton {
	if obj == nil {
		obj = new(singleton)
	}
	return obj
}

// for test

func main() {
	singleton := GetSingleInstance()
	singleton.Name = "www"
	go func() {
		singleton := GetSingleInstance()
		fmt.Println(singleton)
	}()
	time.Sleep(9999)
	fmt.Println(singleton)
}

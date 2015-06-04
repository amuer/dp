package main

import "fmt"
import "reflect"

type SimpleFactory struct {
	FactoryName string
}

type ProductA struct {
	Name string
}

type ProductB struct {
	Name string
}

func (p *ProductA) Create() *ProductA {
	p.Name = "i'm A"
	return p
}

func (p *ProductB) Create() *ProductB {
	p.Name = "i'm B"
	return p
}

func (SimpleFactory) CreateProduct(id int) interface{} {
	switch id {
	case 1:
		return new(ProductA).Create()
	case 2:
		return new(ProductB).Create()
	}
	return nil
}

func main() {
	f := SimpleFactory{}
	p1 := f.CreateProduct(1)
	p2 := f.CreateProduct(2)
	fmt.Println("p1", p1, "p2", p2)
	fmt.Println("p1", reflect.TypeOf(p1), "p2", reflect.TypeOf(p2))
}

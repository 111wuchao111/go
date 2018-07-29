package main

import "fmt"

type P struct {
	name string
}

func (p *P) Shu() {
	fmt.Println("Parent--")
}

func (p *P) H() {
	fmt.Println("H--")
}

type Z struct {
	P
	name string
}

func (z *Z) Shu() {
	fmt.Println("Zi--")
}
func main() {
	z := new(Z)
	z.P.Shu()
	z.Shu()
	//语法糖 调用z.H()相当与z.P.H()
	z.H()
}

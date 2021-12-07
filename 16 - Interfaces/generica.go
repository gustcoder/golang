package main

import "fmt"

type any interface{} //minha implementacao de tipo "any" pra ver se funciona, e funcionou!!

func generica(iface interface{}) {
	fmt.Println(iface)
}

func generica2(iface any) {
	fmt.Println(iface)
}

func main() {
	generica("String")
	generica(1)
	generica(true)

	generica2("String")
	generica2(1)
	generica2(true)
}

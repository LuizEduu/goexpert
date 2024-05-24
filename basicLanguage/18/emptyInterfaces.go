package main

import "fmt"

//type x interface{} // interfaces vazias servem para indicar que todo mundo implementa ele, já que não tem especificações de metodos

func main() {
	//posso passar qualquer tipo de valor, porque a interface vazia implementa todo mundo
	var x interface{} = 10
	var y interface{} = "Hello World"

	showType(x)
	showType(y)

}

func showType(t interface{}) {
	fmt.Printf("O tipo da variável é %T e o valor é %v\n", t, t)
}

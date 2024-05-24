package main

func main() {
	var x interface{} = "Luiz Eduardo"

	result, err := x.(string)

	if !err {
		println(err)
	} else {
		println(result)
	}

}

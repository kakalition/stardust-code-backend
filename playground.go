package main

func maina() {
	var temp interface{}
	temp = nil
	val, _ := temp.(*string)

	println(val)
}

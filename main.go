package main

import ("fmt")

func main() {
	x := make([]string, 0)
	x = append(x, "status")
	output, err := fail2banRequest(x)

	fmt.Println(output)
	fmt.Println(err)

}

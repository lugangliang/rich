package main

func funcA() {
	funcB()

}

func funcB() {
	funcA()
}

func main() {

	funcA()

	return
}

package main

func main() {
	a := App{}
	a.Initializer(
		"michal",
		"michal",
		"school",
	)
	a.Run(":5432")
}

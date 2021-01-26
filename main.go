package main

func main() {
	a := App{}
	a.Initializer(
		"michal",
		"michal",
		"school",
	)
	a.Run("localhost:5432")
}

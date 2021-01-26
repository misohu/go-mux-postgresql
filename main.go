package main

func main() {
	a:= App{}
	a.Initializer(
		"michal",
		"michal", 
		"school",
	)
	a.Runner(":5432")
}
package main

import ("fmt"; "regexp")

func main() {
	m1, _ := regexp.MatchString("p[a-z]+ch", "apple")
	m2, _ := regexp.MatchString("p[a-z]+ch", "peach")
	fmt.Println(m1)
	fmt.Println(m2)

	r1, _ := regexp.Compile("p[a-z]+ch")
	fmt.Println(r1.MatchString("apple"))
	fmt.Println(r1.MatchString("peach"))
}

// todo more

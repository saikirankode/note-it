package main

import "fmt"

func greet(fname string, lname string) string {
	return fmt.Sprint(fname, lname)
}

func wish(finame string, laname string) (x string) {
	x = fmt.Sprint(finame, laname)
	return
}

func welcome(firname string, lasname string) (string, string) {

	return fmt.Sprint(firname, lasname), fmt.Sprint(lasname, firname)
}

func main() {
	fmt.Println(greet("vamsi", " gorapalli"))
	fmt.Println(wish("naveen", " chinta"))
	fmt.Println(welcome(" Arun", " chandra"))
}

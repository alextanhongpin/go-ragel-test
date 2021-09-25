package main

import "fmt"

//go:generate ragel -Z -G2 -o email.go email.rl

func main() {
	emails := FindEmails([]byte(`
	john.doe@mail.com  john.appleseed@mail.com      hello@world.com`))
	//emails := FindEmails([]byte(`john.doe@mail.com`))
	for i, email := range emails {
		fmt.Printf("%d. %q\n", i+1, email)
	}
}

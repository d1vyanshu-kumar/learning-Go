package main // here is the entry point of the program

import (
	

    "github.com/fatih/color"
	"github.com/d1vyanshu-kumar/learning_Go/auth"
	"github.com/d1vyanshu-kumar/learning_Go/user"
)

// here is the convention for the packages in go (Terminal) : go mod init <module_name>/ git repo URl

// now look what if you want to import outside packagege here is the convention: go get "package e.g.-->   github.com/fatih/color"

//  if you want to fixed every thing by own like there is a package which is used  but it is showing unused you can run: "go mod tidy"

// look exactly what this is mean is creating  a package for which we are going to reuse that code/package again and again which will save writing
// more code...
func main() {

	// now look if you want to use the auth package here is how to do it...

	auth.LoginWithCreds("d1vyanshu", "password123") // here is the function call from the auth package...

	sess := auth.CreateSession() // here is the function call from the auth package...
	println("Session created:", sess)

	user := user.User{ // here is the struct from the user package...
		Username: "d1vyanshu",
		Email:    "d1vyanshu@example.com",
		Password: "password123",
	}

	// println("User created:", user.Username, user.Email, user.Password)

	color.Green("Username:", user.Username)

    // Print email in cyan
    color.Cyan("Email: %s", user.Email)
	color.Red("Username:", user.Username)

    // Or just plain fmt
    
 
}






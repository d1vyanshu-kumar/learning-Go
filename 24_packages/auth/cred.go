package auth // package name should be same as the folder name

import "fmt"

// NOW LOOK HERE IS SOME ISSUE WHATEVER THE CONFIG WE ARE GOING TO WRITE HERE IS GOING TO BE USED WITH THIS AUT PACKAGE ONLY SO  IF YOU ARE SEEING THIS loginWithCreds FUNCTION IS NOT going to used outside of this package and this is called scope.

// woooow that is intersting right... so if you want to use this function outside of this package then we have to make it public by capitalizing the first letter of the function name... seriously dude that is so cool...
func LoginWithCreds(username, password string){

	fmt.Println("Logging in with username:", username, "and password:", password)
}

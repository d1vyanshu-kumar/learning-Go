package user


type User struct { // lol as we can see here i am using captial letter for the struct name so that it can be used outside of this package...
	Username string      // now look if i want to use this struct outside of this package then i have to make the fields public by capitalizing the first letter of the field name... this is now getting weird...
	Email    string
	Password string
}
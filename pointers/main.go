package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i

	fmt.Println("i:", p, "j:", &j)
	fmt.Println(*p)

}

// type User struct {
// 	email    string
// 	username string
// 	age      int
// }

// func (u User) Email() string {
// 	return u.email
// }

// func Email(user *User) string {
// 	return user.email
// }

// func (u *User) updateEmail(email string) {
// 	u.email = email
// }

// func main() {

// 	user := User{
// 		email: "agg@foo.com",
// 	}

// 	user.updateEmail("foo@bar.com")
// 	fmt.Println(user.Email())
// }

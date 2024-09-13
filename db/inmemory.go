package db

import (
    "golang.org/x/exp/slices"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}
 
var count int = 3
var users []User

func init() {
    // Initialize the in-memory database with some sample data
    users = []User{
        {ID: 1, Name: "User 1"},
        {ID: 2, Name: "User 2"},
        {ID: 3, Name: "User 3"},
    }
}

func GetUsers() []User {
	return users
}

func AddUser(user User) (id int) {
	count++
	user.ID = count

	users = append(users,user)
	
	return count   
}



func GetUser(id int) User {
	var user User

	for _, user := range users{
		if user.ID == id {
			return user
		}
	}

	return user

}

func DeleteUser(id int) bool {
      
    for i, user := range users{
		if user.ID == id {
			//users = append(users[:i], users[i+1:]...)
            users = slices.Delete(users, i, i+1)
            return true
		}
	}
    return false
	
}

func PutUser(updatedUser User, id int) bool {
    for i, user := range users {
        if user.ID == id {
            users[i].Name = updatedUser.Name
            return true
        }
    }
    return false
}
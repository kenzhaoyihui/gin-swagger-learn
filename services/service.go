package services

import (
	dbb "ginlearn/db"
	user "ginlearn/models"
)

func InsertUser(id int, name string, age int) {
	u := &user.User{
		Id:   id,
		Name: name,
		Age:  age,
	}

	dbb.InsertUser(u)
}

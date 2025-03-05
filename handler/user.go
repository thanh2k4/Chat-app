package handler

import (
	"fmt"
	"net/http"
)

type User struct {
}

func (user *User) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create user")
}

func (user *User) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List user")
}

func (user *User) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get user by ID")
}

func (user *User) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update user by ID")
}

func (user *User) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete user by ID")
}

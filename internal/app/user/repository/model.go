package repository

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	id       primitive.ObjectID `bson:"_id,omitempty"`
	username string             `bson:"username"`
	password string             `bson:"password"`
	email    string             `bson:"email"`
}

func NewUser() *User {
	return &User{}
}

func (u *User) GetId() primitive.ObjectID {
	return u.id
}

func (u *User) SetId(id primitive.ObjectID) *User {
	u.id = id
	return u
}

func (u *User) GetUsername() string {
	return u.username
}

func (u *User) SetUsername(username string) *User {
	u.username = username
	return u
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) SetPassword(password string) *User {
	u.password = password
	return u
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) SetEmail(email string) *User {
	u.email = email
	return u
}

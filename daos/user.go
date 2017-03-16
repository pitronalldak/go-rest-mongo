package daos

import (
	"../models"
	"github.com/ant0ine/go-json-rest/rest"
)

// UserDAO persists user data in database
type UserDAO struct{}

// NewUserDAO creates a new UserDAO
func NewUserDAO() *UserDAO {
	return &UserDAO{}
}

// Get reads the user with the specified ID from the database.
func (dao *UserDAO) Get(r *rest.Request, id int) (*models.User, error) {
	var user models.User
	c := session.DB("test").C("people")

	err := c.Find(query).One(&result)
	return &user, err
}
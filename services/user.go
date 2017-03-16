package services

import (
	"../models"
	"github.com/ant0ine/go-json-rest/rest"
)

// userDAO specifies the interface of the user DAO needed by UserService.
type userDAO interface {
	// Get returns the artist with the specified the artist ID.
	Get(w rest.ResponseWriter, r *rest.Request) (*models.User, error)
	Post(w rest.ResponseWriter, r *rest.Request) (*models.User, error)
}

// UserService provides services related with user.
type UserService struct {
	dao userDAO
}

// NewUserService creates a new UserService with the given user DAO.
func NewUserService(dao userDAO) *UserService {
	return &UserService{dao}
}

func (s *UserService) GetUser(r *rest.Request) (*models.User, error) {
	return s.dao.GetUser(r, 1)
}

func (s *UserService) PostUser(r *rest.Request) (*models.User, error) {
	return s.dao.PostUser(r, 1)
}
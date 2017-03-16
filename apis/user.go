package apis

import (
	"strconv"
	"github.com/ant0ine/go-json-rest/rest"
	"../services"
	"../models"
	"log"
)

type (
	// гыукService specifies the interface for the user service needed by artistResource.
	userService interface {
		Get(w rest.ResponseWriter, r *rest.Request) (*models.User, error)
		Create(w rest.ResponseWriter, r *rest.Request, model *models.User) (*models.User, error)
	}

	// artistResource defines the handlers for the CRUD APIs.
	userResource struct {
		service userService
	}
)

// ServeArtist sets up the routing of artist endpoints and the corresponding handlers.
func ServeUserResource(r *rest, service userService) {
	res := &userResource{service}
	router, err := rest.MakeRouter(
		rest.Get("/user", res.get),
		rest.Post("/user", res.create),
	)

	if err != nil {
		log.Fatal(err)
	}

	return router
}


func (r *userResource) get(w rest.ResponseWriter, r *rest.Request) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	response, err := r.service.Get(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}

	return w.WriteJson(response)
}

func (r *userResource) create(w rest.ResponseWriter, r *rest.Request) error {
	var model models.Artist
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Create(app.GetRequestScope(c), &model)
	if err != nil {
		return err
	}

	return w.WriteJson(response)
}

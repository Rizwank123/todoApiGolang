package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/rizwank123/models"
	"github.com/rizwank123/service"
)

type UserController struct {
	us service.UserService
}

func NewUserController(us service.UserService) UserController {
	return UserController{us: us}
}

// CreateUser godoc
//
//	@Summary		Create a user
//	@Description	Get a new user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			u	body		models.User	true	"create user"
//	@Success		201	{object}	models.User
//	@Router			/users [post]
func (a *UserController) CreateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		fmt.Printf("error in binding data%v\n", err)
		return err
	}
	d, err := a.us.CreateUser(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Not created")
	}
	return c.JSON(http.StatusCreated, d)
}

// getUser 	godoc
//
//	@Summary		get user bu Id
//	@Description	get user by id
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"User ID"
//	@Success		200	{object}	models.User
//	@Router			/users/{id} [get]
func (a *UserController) FindByID(c echo.Context) error {
	Id := c.Param("id")
	uid, err := strconv.Atoi(Id)
	if err != nil {
		fmt.Printf("Error parsing to into int %v\n", err)
		return c.JSON(http.StatusBadRequest, "invalid userId")
	}
	usr, err := a.us.FindByID(uid)

	if err != nil {
		fmt.Printf("error to in getting data %v\n", err)
		return c.JSON(http.StatusNotFound, "No Record found")
	}
	if usr.Id == 0 && usr.Name == "" && usr.Email == "" {
		return c.JSON(http.StatusNotFound, "User Not Found")
	}
	return c.JSON(http.StatusOK, usr)

}

// update	godoc
//
//	@Summary		update user
//	@Description	update user by id
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string			true	"User ID"
//	@Param			user	body		models.User		true	"User Data"
//	@Success		200		{string}	successfully	updated
//	@Router			/users/{id} [put]
func (a *UserController) Update(c echo.Context) error {
	var user models.User
	Id := c.Param("id")
	if err := c.Bind(&user); err != nil {
		return err
	}
	userId, err := strconv.Atoi(Id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)

	}
	err = a.us.Update(&user, userId)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusAccepted, "Update successfully")
}

// delete	godoc
//
//	@Summary		delete user
//	@Description	delete user by id
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string			true	"User ID"
//	@Success		200	{string}	successfully	Deleted
//	@Router			/users/{id} [delete]
func (a *UserController) Delete(c echo.Context) error {
	id := c.Param("id")
	uid, err := strconv.Atoi(id)
	if err != nil {
		c.Logger().Fatal(fmt.Printf("%v\n", err))
	}
	err = a.us.Delete(uid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Row Not deleted")
	}
	return c.JSON(http.StatusOK, "Row Deleted Successfully")
}

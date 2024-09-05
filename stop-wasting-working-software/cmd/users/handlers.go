package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *UserService) createUserHandler(c echo.Context) error {
	// Parse request body to get user data
	var user User
	if err := c.Bind(&user); err != nil {
		return err
	}

	// Create user using the UserService
	err := s.CreateUser(c.Request().Context(), user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}

func (s *UserService) getUserByIdHandler(c echo.Context) error {
	userID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	user, err := s.FetchUserByID(c.Request().Context(), userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	return c.JSON(http.StatusOK, user)
}

func (s *UserService) deleteUserHandler(c echo.Context) error {
	userID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	err = s.DeleteUser(c.Request().Context(), userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (s *UserService) updateUserHandler(c echo.Context) error {
	userID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	var updatedUser User
	if err := c.Bind(&updatedUser); err != nil {
		return err
	}

	err = s.UpdateUser(c.Request().Context(), userID, updatedUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, updatedUser)
}

func (s *UserService) loginHandler(c echo.Context) error {
	// Parse request body to get username and password
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.Bind(&credentials); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	token, err := s.authenticateUser(credentials.Username, credentials.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

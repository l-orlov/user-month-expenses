package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	ierrors "github.com/l-orlov/user-month-expenses/internal/errors"
	"github.com/l-orlov/user-month-expenses/internal/models"
	"github.com/pkg/errors"
)

var ErrNotValidIDParameter = errors.New("not valid id parameter")

func (h *Handler) CreateUser(c *gin.Context) {
	setHandlerNameToLogEntry(c, "CreateUser")

	var (
		user models.UserToCreate
		err  error
	)
	if err = c.BindJSON(&user); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	id, err := h.svc.User.CreateUser(c, user)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) GetUserByID(c *gin.Context) {
	setHandlerNameToLogEntry(c, "GetUserByID")

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		h.newErrorResponse(
			c, http.StatusBadRequest, ierrors.NewBusiness(ErrNotValidIDParameter, ""),
		)
		return
	}

	user, err := h.svc.User.GetUserByID(c, id)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	if user == nil {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	setHandlerNameToLogEntry(c, "UpdateUser")

	var user models.User
	var err error
	if err = c.BindJSON(&user); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err = h.svc.User.UpdateUser(c, user); err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) GetAllUsers(c *gin.Context) {
	setHandlerNameToLogEntry(c, "GetAllUsers")

	users, err := h.svc.User.GetAllUsers(c)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	if users == nil {
		c.JSON(http.StatusOK, []struct{}{})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *Handler) GetUsersWithParameters(c *gin.Context) {
	var params models.UserParams
	if err := c.BindJSON(&params); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	users, err := h.svc.User.GetUsersWithParameters(c, params)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	if users == nil {
		c.JSON(http.StatusOK, []struct{}{})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	setHandlerNameToLogEntry(c, "DeleteUser")

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		h.newErrorResponse(
			c, http.StatusBadRequest, ierrors.NewBusiness(ErrNotValidIDParameter, ""),
		)
		return
	}

	if err = h.svc.User.DeleteUser(c, id); err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

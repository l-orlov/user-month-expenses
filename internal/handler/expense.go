package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	ierrors "github.com/l-orlov/user-month-expenses/internal/errors"
	"github.com/l-orlov/user-month-expenses/internal/models"
	"github.com/pkg/errors"
)

var ErrNotValidSizeParameter = errors.New("not valid size parameter")

func (h *Handler) CreateUserExpense(c *gin.Context) {
	setHandlerNameToLogEntry(c, "CreateUser")

	var (
		expense models.UserExpense
		err     error
	)
	if err = c.BindJSON(&expense); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := h.svc.UserExpense.CreateUserExpense(c, expense); err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) GetUserExpenseByUserID(c *gin.Context) {
	setHandlerNameToLogEntry(c, "GetUserByID")

	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		h.newErrorResponse(
			c, http.StatusBadRequest, ierrors.NewBusiness(ErrNotValidIDParameter, ""),
		)
		return
	}

	expense, err := h.svc.UserExpense.GetUserExpenseByID(c, userID)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	if expense == nil {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, expense)
}

func (h *Handler) UpdateUserExpense(c *gin.Context) {
	setHandlerNameToLogEntry(c, "UpdateUser")

	var expense models.UserExpense
	var err error
	if err = c.BindJSON(&expense); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err = h.svc.UserExpense.UpdateUserExpense(c, expense); err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) GetAllUserExpenses(c *gin.Context) {
	setHandlerNameToLogEntry(c, "GetAllUsers")

	expenses, err := h.svc.UserExpense.GetAllUserExpenses(c)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	if expenses == nil {
		c.JSON(http.StatusOK, []struct{}{})
		return
	}

	c.JSON(http.StatusOK, expenses)
}

func (h *Handler) GetUserExpensesWithParameters(c *gin.Context) {
	var params models.UserExpenseParams
	if err := c.BindJSON(&params); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	expenses, err := h.svc.UserExpense.GetUserExpensesWithParameters(c, params)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	if expenses == nil {
		c.JSON(http.StatusOK, []struct{}{})
		return
	}

	c.JSON(http.StatusOK, expenses)
}

func (h *Handler) GetUserExpensesByCategories(c *gin.Context) {
	sizeStr, ok := c.GetQuery("size")
	if !ok {
		h.newErrorResponse(
			c, http.StatusBadRequest, ierrors.NewBusiness(ErrNotValidSizeParameter, ""),
		)
		return
	}

	size, err := strconv.ParseUint(sizeStr, 10, 64)
	if err != nil {
		h.newErrorResponse(
			c, http.StatusBadRequest, ierrors.NewBusiness(ErrNotValidSizeParameter, ""),
		)
		return
	}

	expenses, err := h.svc.UserExpense.GetUserExpensesByCategories(c, uint16(size))
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	if expenses == nil {
		c.JSON(http.StatusOK, []struct{}{})
		return
	}

	c.JSON(http.StatusOK, expenses)
}

func (h *Handler) GetUserExpensesByUserIDAndCategories(c *gin.Context) {
	// ToDo: use query parameters and schema package may be
	userID, err := strconv.ParseUint(c.Param("userID"), 10, 64)
	if err != nil {
		h.newErrorResponse(
			c, http.StatusBadRequest, ierrors.NewBusiness(ErrNotValidIDParameter, ""),
		)
		return
	}

	sizeStr, ok := c.GetQuery("size")
	if !ok {
		h.newErrorResponse(
			c, http.StatusBadRequest, ierrors.NewBusiness(ErrNotValidSizeParameter, ""),
		)
		return
	}

	size, err := strconv.ParseUint(sizeStr, 10, 64)
	if err != nil {
		h.newErrorResponse(
			c, http.StatusBadRequest, ierrors.NewBusiness(ErrNotValidSizeParameter, ""),
		)
		return
	}

	expenses, err := h.svc.UserExpense.GetUserExpensesByUserIDAndCategories(c, userID, uint16(size))
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	if expenses == nil {
		c.JSON(http.StatusOK, []struct{}{})
		return
	}

	c.JSON(http.StatusOK, expenses)
}

func (h *Handler) DeleteUserExpense(c *gin.Context) {
	setHandlerNameToLogEntry(c, "DeleteUser")

	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		h.newErrorResponse(
			c, http.StatusBadRequest, ierrors.NewBusiness(ErrNotValidIDParameter, ""),
		)
		return
	}

	if err = h.svc.UserExpense.DeleteUserExpense(c, userID); err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

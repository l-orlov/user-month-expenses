package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ierrors "github.com/l-orlov/user-month-expenses/internal/errors"
	"github.com/sirupsen/logrus"
)

const (
	ctxLogEntry = "log-entry"
)

type errorResponse struct {
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

func (h *Handler) newErrorResponse(c *gin.Context, statusCode int, err error) {
	logEntry := h.getLogEntry(c)

	if customErr, ok := err.(*ierrors.Error); ok {
		handleCustomError(c, logEntry, customErr)
		return
	}

	handleDefaultError(c, logEntry, err, statusCode)
}

func handleCustomError(c *gin.Context, logEntry *logrus.Entry, err *ierrors.Error) {
	var statusCode int

	if err.Level == ierrors.Business {
		logEntry.Debug(err)
		statusCode = http.StatusBadRequest
	} else {
		logEntry.Error(err)
		statusCode = http.StatusInternalServerError
	}

	c.AbortWithStatusJSON(statusCode, &errorResponse{
		Message: err.Error(),
		Detail:  err.Detail,
	})
}

func handleDefaultError(c *gin.Context, logEntry *logrus.Entry, err error, statusCode int) {
	errResp := &errorResponse{
		Message: err.Error(),
	}
	if statusCode >= http.StatusBadRequest && statusCode < http.StatusInternalServerError {
		logEntry.Debug(err)
		errResp.Detail = ierrors.DetailBusiness
	} else {
		logEntry.Error(err)
		errResp.Detail = ierrors.DetailServer
	}

	c.AbortWithStatusJSON(statusCode, errResp)
}

func (h *Handler) getLogEntry(c *gin.Context) *logrus.Entry {
	logEntryValue, ok := c.Get(ctxLogEntry)
	if !ok {
		return logrus.NewEntry(h.log)
	}

	logEntry, ok := logEntryValue.(*logrus.Entry)
	if !ok {
		return logrus.NewEntry(h.log)
	}

	return logEntry
}

func setHandlerNameToLogEntry(c *gin.Context, handlerName string) {
	logEntryValue, ok := c.Get(ctxLogEntry)
	if !ok {
		return
	}

	logEntry, ok := logEntryValue.(*logrus.Entry)
	if !ok {
		return
	}

	logEntry.WithField("method", handlerName)
	c.Set(ctxLogEntry, logEntry)
}

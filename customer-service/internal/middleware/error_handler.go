package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paulfernandosr/order-processor-golang/customer-service/internal/model"
)

func ErrorHandler(context *gin.Context) {
	context.Next()

	if len(context.Errors) == 0 {
		return
	}

	err := context.Errors.Last().Err
	var errModel *model.Error

	if ok := errors.As(err, &errModel); ok {
		context.AbortWithStatusJSON(errModel.Status, gin.H{"error": errModel.Error()})
		return
	}

	context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

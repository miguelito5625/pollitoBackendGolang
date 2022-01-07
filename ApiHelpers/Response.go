package ApiHelpers

import (
	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Status int
	// Meta     interface{}
	Message string
	Data    interface{}
}

func RespondJSON(w *gin.Context, status int, payload interface{}, message string) {
	// fmt.Println(message, payload)
	var res ResponseData
	res.Status = status
	//res.Meta = utils.ResponseMessage(status)
	res.Data = payload
	res.Message = message
	w.JSON(status, res)
}

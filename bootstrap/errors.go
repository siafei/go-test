package bootstrap

import(
	"net/http"
	"github.com/gin-gonic/gin"
)



type Error struct{
	StatusCode  int 	`json:"_"`
	Code 		int 	`json:"code"`
	Msg			string	`json:"msg"` 
}


var (
	Success     = NewError(http.StatusOK, 0, "success")
	ServerError = NewError(500, 500, "系统异常，请稍后重试!")
	NotFound    = NewError(http.StatusNotFound, 404, "你找的页面不存在")
	NotAuthorized = NewError(403, 403, "操作未经授权")
)


func (e *Error) Error() string {
	return e.Msg
}


func OptionError(msg string) *Error{
	return &Error{
		StatusCode:	400,
		Code:0,
		Msg:msg,
	}
}

func NewError(statuscode,code int ,msg string) *Error  {
	return &Error{
		StatusCode:	statuscode,
		Code:code,
		Msg:msg,
	}
}

func HandleNotFound(c *gin.Context) {
	err := NotFound
	c.JSON(err.StatusCode,err)
	return
}
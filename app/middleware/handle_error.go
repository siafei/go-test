package middleware

import(
	"github.com/gin-gonic/gin"
	."go-test/bootstrap"
)


func ErrHandler() gin.HandlerFunc  {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var Err *Error
				if e,ok := err.(*Error); ok {
					Err = e
				}else if e, ok := err.(error); ok {
					Err = OptionError(e.Error())
				}else{
					Err = ServerError
				}
				// 记录一个错误的日志
				c.JSON(Err.StatusCode,Err)
				return
			}
		}()
		c.Next()
	}
}
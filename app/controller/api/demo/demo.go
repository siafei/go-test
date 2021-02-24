package demo

import(
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"go-test/app/models"
	"go-test/bootstrap"
)


func Index(c *gin.Context)  {
	panic(bootstrap.OptionError("犯个错看看"))
	// c.JSON(200, gin.H{
	// 	"message": "api home",
	// })
}

func Login(c *gin.Context)  {
	pass := c.PostForm("password")
	right := string(c.PostForm("right"))
	hash, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	err := bcrypt.CompareHashAndPassword([]byte(right), []byte(pass))
	res := "ok"
	if err != nil {
        res = "worng"
    }
	c.JSON(200, gin.H{
		"errs":err,
		"right":right,
		"hash":string(hash),
		"err": res,
	})
}

func Register(c *gin.Context)  {
	user := models.User{}
	user_create := map[string]interface{}{
		"uname"		:	string(c.PostForm("uname")),
		"password"	:	string(c.PostForm("password")),
		"mobile"	: 	string(c.PostForm("mobile")),
	}
	err := user.CreateUser(user_create)
	message := "添加成功"
	if err != nil {
		message = "添加失败"
	}
	c.JSON(200, gin.H{
		"message":message,
	})
}
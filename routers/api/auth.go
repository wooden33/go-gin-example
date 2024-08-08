package api

import (
	"gin-blog/models"
	"gin-blog/pkg/e"
	"gin-blog/pkg/logging"
	"gin-blog/pkg/util"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username, password := c.Query("username"), c.Query("password")

	valid := validation.Validation{}
	a := auth{username, password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]any)
	code := e.INVALID_PARAMS

	if !ok {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	} else {
		if models.CheckAuth(username, password) {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

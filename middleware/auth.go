package  middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func IsLogin() gin.HandlerFunc {
	return func(c *gin.Context) {

		session := sessions.Default(c)
		adminId := session.Get("AdminId")

		if adminId == nil {
			c.HTML(401, "login_logout/login.html", gin.H{})
			c.Abort()
		} else {
			c.Set("AdminId", adminId)
			c.Next()	
		}
	}
}
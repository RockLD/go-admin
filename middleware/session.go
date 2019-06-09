package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"net/http"
)

var Store = sessions.NewCookieStore([]byte("go-admin"))
func SessionAuth(c *gin.Context){
	session,err := Store.Get(c.Request,"admin_id")
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error_msg":"session error",
			"error_code":-1,
		})
	}
	if session.Values["admin_id"] == nil{
		c.Redirect(http.StatusFound,"/ad/login")
	}else{
		c.Next()
	}



	//session := sessions.Default(c)
	//adminId := session.Get("admin_id")
	//
	//if adminId == nil {
	//	c.Redirect(http.StatusFound,"/ad/login")
	//}else {
	//	c.Next()
	//}

}

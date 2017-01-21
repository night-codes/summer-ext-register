package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"time"
)

func setCookie(c *gin.Context, name string, value string) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    name,
		Value:   value,
		Path:    "/",
		MaxAge:  32000000,
		Expires: time.Now().AddDate(1, 0, 0),
	})
}

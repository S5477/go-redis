package controller

import (
	"net/http"

	"github.com/S5477/go-redis/models"
	"github.com/S5477/go-redis/store"

	"github.com/gin-gonic/gin"
)

const KEY = "key"

func Add(context *gin.Context) {
	var v models.Add

	if err := context.ShouldBindJSON(&v); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	store.SetString(v)

	context.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func Get(context *gin.Context) {
	var v models.Get

	if err := context.ShouldBindQuery(&v); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	key := getParam(context, KEY)

	res := store.GetString(key)

	context.JSON(http.StatusOK, gin.H{"data": res})
}

func HAdd(context *gin.Context) {

}

func HGet(context *gin.Context) {

}

func getParam(context *gin.Context, param string) string {
	var s models.Get

	context.ShouldBindQuery(&s)

	return s.Key
}

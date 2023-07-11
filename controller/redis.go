package controller

import (
	"net/http"

	"github.com/S5477/go-redis/store"
	"github.com/S5477/go-redis/validation"
	"github.com/gin-gonic/gin"
)

const VALUE = "value"
const KEY = "key"

func Add(context *gin.Context) {
	var v validation.AddRequested

	if err := context.ShouldBind(&v); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	value := getParam(context, VALUE)

	key := getParam(context, KEY)

	store.SetString(key, value)

	context.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func Get(context *gin.Context) {
	var v validation.GetRequested

	if err := context.ShouldBind(&v); err != nil {
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
	param, issetVal := context.GetQuery(param)

	if !issetVal {
		panic("no " + param)
	}

	return param
}

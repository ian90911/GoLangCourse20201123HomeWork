package controllers

import (
	"hw/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Data *[]models.Role

// 取得全部資料
func GetAllRole(c *gin.Context) {
	c.JSON(http.StatusOK, Data)
}

// 取得單一筆資料
func GetOneRole(c *gin.Context) {

}

// 新增資料
func PostRole(c *gin.Context) {

}

// 更新資料, 更新角色名稱與介紹
func PutRole(c *gin.Context) {

}

// 刪除資料
func DeleteRole(c *gin.Context) {

}

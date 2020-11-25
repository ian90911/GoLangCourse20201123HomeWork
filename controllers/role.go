package controllers

import (
	"hw/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Data *[]models.Role
var GetLatestRoleId func() *uint
var GetLatestSkillId func() *uint

// 取得全部資料
func GetAllRole(c *gin.Context) {
	c.JSON(http.StatusOK, Data)
}

// 取得單一筆資料
func GetOneRole(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		log.Fatal(err)
	}

	for i := range *Data {
		if (*Data)[i].ID == uint(id) {
			c.JSON(http.StatusOK, (*Data)[i])
			return
		}
	}
	c.AbortWithStatus(http.StatusNotFound)
}

// 新增資料
func PostRole(c *gin.Context) {
	role := models.Role{}
	err := c.BindJSON(&role)
	if err != nil {
		log.Fatalln(err)
	}

	role.ID = *GetLatestRoleId()
	for i := range role.Skills {
		role.Skills[i].ID = *GetLatestSkillId()
	}
	roleSlice := make([]models.Role, 1, 1)
	roleSlice = append(roleSlice, role)
	*Data = append(*Data, roleSlice...)

	c.JSON(http.StatusOK, role)
}

// 更新資料, 更新角色名稱與介紹
func PutRole(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		log.Fatal(err)
	}

	role := models.RolePutVM{}
	err = c.BindJSON(&role)
	if err != nil {
		log.Fatalln(err)
	}

	for i := range *Data {
		if (*Data)[i].ID == uint(id) {
			(*Data)[i].Name = role.Name
			(*Data)[i].Summary = role.Summary
			c.JSON(http.StatusOK, (*Data)[i])
			return
		}
	}
	c.AbortWithStatus(http.StatusNotFound)
}

// 刪除資料
func DeleteRole(c *gin.Context) {

}

package main

import (
	"hw/controllers"
	"hw/models"

	"golang.org/x/net/context"

	"github.com/gin-gonic/gin"
)

var Data []models.Role
var LatestRoleId uint
var LatestSkillId uint

func main() {
	context.Background()
	router := gin.Default()
	initRoutes(router)
	router.Run(":8080")
}

func init() {
	Data = []models.Role{
		{
			ID:      1,
			Name:    "阿修羅",
			Summary: "死國魖族最強者，在歷經多場戰役統一死國，直接挑戰天者權威，不服其領導，要求讓三族和平共處，因此被天者以增加資源為名，前往打通連接火宅佛獄、死國及苦境的莫汗走廊，未料在工程半途被楓岫主人所乘駕的隕石撞毀，阿修羅在天者刻意算計下意外身亡，身葬苦境與死國間的異次元。",
			Skills: []models.RoleSkill{
				{ID: 1, Type: models.MartialArts, Name: "天之爆"},
				{ID: 2, Type: models.MartialArts, Name: "魔之狂"},
				{ID: 3, Type: models.MartialArts, Name: "天之渦"},
				{ID: 4, Type: models.MartialArts, Name: "闇之爆"},
				{ID: 5, Type: models.Magic, Name: "山河凝元·天地共引"},
				{ID: 6, Type: models.Magic, Name: "地之火·九天滅絕"},
			},
		},
		{
			ID:      2,
			Name:    "白塵子",
			Summary: "火宅佛獄凱旋侯的副體之一，本名黑枒君，臥底中原武林，向佛獄通風報信，最後被素還真所殺並冒充其身份一探佛獄之秘。",
			Skills: []models.RoleSkill{
				{ID: 7, Type: models.MartialArts, Name: "凝宇化空"},
				{ID: 8, Type: models.MartialArts, Name: "反神源"},
			},
		},
	}
	LatestRoleId = Data[len(Data)-1].ID
	LatestSkillId = Data[len(Data)-1].Skills[len(Data[len(Data)-1].Skills)-1].ID

	controllers.Data = &Data
	controllers.GetLatestRoleId = newRoleId
	controllers.GetLatestSkillId = newSkillId
}

func newRoleId() *uint {
	LatestRoleId++
	return &LatestRoleId
}

func newSkillId() *uint {
	LatestSkillId++
	return &LatestSkillId
}

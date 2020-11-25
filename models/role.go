package models

// 布袋戲角色
type Role struct {
	ID      uint        `json:"id"`      // Key
	Name    string      `json:"name"`    // 角色名稱
	Summary string      `json:"summary"` // 介紹
	Skills  []RoleSkill `json:"skills"`  // 會使用的招式
}

// RoleSkill 用來記錄招式名稱
type RoleSkill struct {
	ID   uint      `json:"id"`   //key
	Type SkillType `json:"type"` // 招式分類
	Name string    `json:"name"` // 招式名稱
}

type SkillType string

const (
	MartialArts SkillType = "武學"
	Magic       SkillType = "法術"
)

type RoleVM struct {
	ID      uint   `json:"id"`      // Key
	Name    string `json:"name"`    // 角色名稱
	Summary string `json:"summary"` // 介紹
}

type RolePutVM struct {
	Name    string `json:"name"`    // 角色名稱
	Summary string `json:"summary"` // 介紹
}

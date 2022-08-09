// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Cert is the golang structure for table cert.
type Cert struct {
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	IsDeleted int         `json:"isDeleted" description:""`
	Number    string      `json:"number"    description:"证书 ID"`
	DevId     string      `json:"devId"     description:"开发者 ID"`
	DevOrg    string      `json:"devOrg"    description:"开发者组织"`
	CertPass  string      `json:"certPass"  description:"证书密码"`
	BeginTime string      `json:"beginTime" description:"颁发时间"`
	EndTime   string      `json:"endTime"   description:"过期时间"`
	Content   string      `json:"content"   description:"证书内容"`
	Id        int         `json:"id"        description:""`
}
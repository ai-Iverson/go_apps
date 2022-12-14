// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// File is the golang structure for table file.
type File struct {
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	IsDeleted int         `json:"isDeleted" description:""`
	Md5       string      `json:"md5"       description:"文件MD5"`
	Sha1      string      `json:"sha1"      description:"文件 SHA1"`
	Name      string      `json:"name"      description:"文件名称"`
	DownUrl   string      `json:"downUrl"   description:"下载地址"`
	DownCount int         `json:"downCount" description:"下载次数"`
	VersionId string      `json:"versionId" description:"版本 ID"`
	CertId    string      `json:"certId"    description:"证书 ID"`
	Id        int         `json:"id"        description:""`
}

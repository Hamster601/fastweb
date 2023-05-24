package role

import "time"

// Role 角色权限表
//go:generate gormgen -structs Role -input .
type Role struct {
	Id              int32     // 主键
	Rolename        string    // 角色名
	RoleDescription string    // 角色描述
	API             string    // 拥有的API权限
	IsDeleted       int32     // 是否删除 1:是  -1:否
	CreatedAt       time.Time `gorm:"time"` // 创建时间
	CreatedUser     string    // 创建人
	UpdatedAt       time.Time `gorm:"time"` // 更新时间
	UpdatedUser     string    // 更新人
}

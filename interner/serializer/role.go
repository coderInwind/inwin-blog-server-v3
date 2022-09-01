package serializer

import (
	"inwind-blog-server-v3/interner/model"
)

type RoleOption struct {
	Role  string `json:"role"`
	Id    uint   `json:"id"`
	Power int    `json:"power"`
}

func BuildRoleOptions(items []model.Role) (roleOptions []RoleOption) {
	for _, v := range items {
		option := BuildRoleOption(v)
		roleOptions = append(roleOptions, option)
	}
	return roleOptions
}

func BuildRoleOption(item model.Role) RoleOption {
	return RoleOption{
		Role:  item.Role,
		Id:    item.ID,
		Power: item.Power,
	}
}

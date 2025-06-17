package authmodels

type Role struct {
	Id          uint
	Name        string        `gorm:"uniqueIndex;not null"`
	Permissions []*Permission `gorm:"many2many:role_permissions"`
}

type Permission struct {
	Id    uint
	Name  string  `gorm:"uniqueIndex;not null"`
	Roles []*Role `gorm:"many2many:role_permissions"`
}

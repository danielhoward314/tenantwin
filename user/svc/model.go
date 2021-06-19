package svc

import (
	"time"
)

type UserTypeEnum int64

const (
	Tenant UserTypeEnum = iota
	Owner
	BuildingStaff
	PortalAdmin
	PortalAdminSuper
)

type PortalAccessStatusEnum int64

const (
	Pending PortalAccessStatusEnum = iota
	Active
	Revoked
)

type User struct {
	ID                 string                 `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"-"`
	CreatedAt          time.Time              `json:"-"`
	UpdatedAt          time.Time              `json:"-"`
	DeletedAt          *time.Time             `sql:"index" json:"-"`
	Phone              string                 `json:"phone" gorm:"size:50;not null; unique"`
	Email              string                 `json:"email" gorm:"size:50;not null; unique"`
	Password           string                 `json:"-" gorm:"size:25;not null; unique"`
	UserType           UserTypeEnum           `json:"user_type"`
	PortalAccessStatus PortalAccessStatusEnum `json:"portal_access_status"`
}

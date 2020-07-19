package models

type Job struct {
	JobID string `gorm:"primary_key"`
	JobName string `gorm:"type:varchar(256);not null;"`
	WorkDir string `gorm:"type:varchar(1024);not null;"`
	Command string `gorm:"type:varchar(1024);not null;"`
}

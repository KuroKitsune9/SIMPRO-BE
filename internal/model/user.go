package model

type User struct {
	ID         int    `gorm:"column:id;primaryKey;autoIncrement"`
	Name       string `gorm:"column:name"`
	Username   string `gorm:"column:username"`
	Password   string `gorm:"column:password"`
	Role       string `gorm:"column:role"`
	BranchCode string `gorm:"column:branch_code"`
}

func (User) TablenName() string {
	return "users"
}

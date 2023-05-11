package user

type User struct {
	PK   int64  `gorm:"pk"`
	Name string `gorm:"name"`
}

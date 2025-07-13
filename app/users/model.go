package users

type UserModel struct {
	ID       string `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	DOB      string `gorm:"not null"`
	Email    string `gorm:"not null;uniqueIndex"`
	Password string `gorm:"not null"`
	Phone    string `gorm:"not null;unique"`
	Role     string `gorm:"not null"`
}

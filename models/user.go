package models

type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	NAME     string `gorm:"not null" json:"name"`
	LASTNAME string `gorm:"not null" json:"lastname"`
	AGE      string `gorm:"not null" json:"age"`
}

func (u *User) GetId() uint {
	return u.ID
}

func (u *User) GetName() string {
	return u.NAME
}

func (u *User) GetLastname() string {

	return u.LASTNAME
}

func (u *User) GetAge() string {
	return u.AGE
}

func (u *User) SetId(id uint) {
	u.ID = id
}

func (u *User) SetName(name string) {
	u.NAME = name
}

func (u *User) SetLastname(lastname string) {
	u.LASTNAME = lastname
}

func (u *User) SetAge(age string) {
	u.AGE = age
}

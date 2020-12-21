package users_store

type UserStore interface {
	Get(id string) (*User, error)
	Create(user *User) (*User, error)
	Delete(id string) error
	GetByEmail(email string) (*User, error)
	GetByPhone(phone string) (*User, error)
}

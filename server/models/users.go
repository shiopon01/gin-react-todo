package models

// Users ...
type Users struct {
	ID       int    `json:"id" xorm:"'id'"`
	Username string `json:"username" xorm:"'username'"`
	Name     string `json:"name" xorm:"'name'"`
	Password string `json:"password" xorm:"'password'"`
}

// NewUser ...
func NewUser(username string, name string, password string) Users {
	return Users{
		Username: username,
		Name:     name,
		Password: password,
	}
}

// UserRepository ...
type UserRepository struct {
}

// NewUserRepository ...
// Run this method to get UserRepository
func NewUserRepository() UserRepository {
	return UserRepository{}
}

// GetByID ...
// UserRepository's method
func (m UserRepository) GetByID(id int) *Users {
	user := Users{ID: id}
	has, _ := engine.Get(&user)

	if has {
		return &user
	}

	return nil
}

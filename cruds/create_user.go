package cruds

import "snsback/db"

func CreateUserByJSON(name, userID, email, password string) db.User {
	user := db.User{
		Name:     name,
		UserID:   userID,
		Email:    email,
		Password: password,
	}
	db.DB.Create(&user)
	return user
}

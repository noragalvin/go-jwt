package models

type User struct {
	ID        string `json:"id"`
	UserName  string `json:"username"`
	Password  string `json:"password"`
	Token     string `json:"token"`
	ExpiredAt string `json:"expied_at`
}

func UserShow() User {
	var user = User{ID: "1", UserName: "minhnora98"}
	// mapstructure.Decode(decoded.(jwt.MapClaims), &user)
	return user
}

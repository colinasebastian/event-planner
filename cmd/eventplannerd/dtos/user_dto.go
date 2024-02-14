package dtos

type UserData struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

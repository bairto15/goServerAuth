package structures

type User struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Login       string `json:"login" db:"login"`
	Password    string `json:"password" db:"password"`
	OldPassword string `json:"old_password" db:"old_password"`
	Root        int    `json:"root" db:"root"`
	Role        string `json:"role" db:"role"`
}

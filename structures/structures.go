package structures

type User struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Root     int    `json:"root"`
	Сhilds   []int  `json:"childs"`
}

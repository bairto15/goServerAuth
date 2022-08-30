package repository

import (
	"goServerAuth/structures"
)

//Получить список дочерних пользователей
func (r *AuthPostgres) GetUsers(id int) ([]structures.User, error) {
	var users []structures.User

	query := "SELECT * FROM users WHERE root=$1"
	err := r.db.Select(&users, query, id)

	return users, err
}

//Получить данные пользователя по id
func (r *AuthPostgres) GetUser(id int) (structures.User, error) {
	var user structures.User

	query := "SELECT * FROM users WHERE id=$1"
	err := r.db.Get(&user, query, id)

	return user, err
}

//Изменить данные пользователя
func (r *AuthPostgres) EditUser(user structures.User) error {
	query := "UPDATE users SET name=$1, password=$2 WHERE id=$3 AND login=$4"
	_, err := r.db.Exec(query, user.Name, user.Password, user.Id, user.Login)

	return err
}

//Изменить данные Админа
func (r *AuthPostgres) EditAdmin(user structures.User) error {
	query := "UPDATE admins SET name=$1, password=$2 WHERE id=$3 AND login=$4"
	_, err := r.db.Exec(query, user.Name, user.Password, user.Id, user.Login)

	return err
}

//Удалить пользователя по id
func (r *AuthPostgres) DeleteUser(idUser int, idAdmin int) error {
	query := "DELETE FROM users WHERE id=$1 AND root=$2"
	_, err := r.db.Exec(query, idUser, idAdmin)

	return err
}

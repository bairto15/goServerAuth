package repository

import (
	"fmt"
	"goServerAuth/structures"
)

//Получить список дочерних пользователей
func (r *AuthPostgres) GetUsers(id int) ([]structures.User, error) {
	var users []structures.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE root=$1", "users")
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

//Изменить данные пользователя по id
func (r *AuthPostgres) EditUser(user structures.User) error {
	//query := "UPDATE name, password FROM users SET name=$1, password=$2 WHERE id=$3 AND root=$4"
	query := "UPDATE users SET name=? password=? WHERE id=? IN (?)"
	_, err := r.db.Exec(query, user.Name, user.Password, user.Id, user.Root)

	return err
}

//Удалить пользователя по id
func (r *AuthPostgres) DeleteUser(idUser int, idAdmin int) error {
	query := "DELETE FROM users WHERE id=$1 AND root=$2"
	_, err := r.db.Exec(query, idUser, idAdmin)

	return err
}

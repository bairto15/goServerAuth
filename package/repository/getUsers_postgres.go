package repository

import (
	"fmt"
	"goServerAuth/structures"
)

func (r *AuthPostgres) GetUser(login, password string) (structures.User, error) {
	var user structures.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE login=$1 AND password=$2", "users")
	err := r.db.Get(&user, query, login, password)
	
	return user, err
}

func (r *AuthPostgres) GetUsers(id int) ([]structures.User, error) {
	// var users []structures.User
	// query := fmt.Sprintf("SELECT * FROM %s WHERE login=$1 AND password=$2", "users")
	return nil, nil
}
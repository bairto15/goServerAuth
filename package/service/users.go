package service

import "goServerAuth/structures"

//Создания админа
func (s *AuthService) CreateAdmin(user structures.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateAdmin(user)
}

//Создание пользователя
func (s *AuthService) CreateUser(user structures.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

//Список дочерних пользователей
func (s *AuthService) GetUsers(id int) ([]structures.User, error) {
	return s.repo.GetUsers(id)
}

//Получить данные пользователя по id
func (s *AuthService) GetUser(id int) (structures.User, error) {
	return s.repo.GetUser(id)
}

//Изменить данные пользователя
func (s *AuthService) EditUser(user structures.User) error {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.EditUser(user)
}

//Изменить данные Админа
func (s *AuthService) EditAdmin(user structures.User) error {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.EditAdmin(user)
}

//Удалить пользователя по id
func (s *AuthService) DeleteUser(idUser int, idAmin int) error {
	return s.repo.DeleteUser(idUser, idAmin)
}

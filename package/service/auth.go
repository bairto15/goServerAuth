package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"goServerAuth/package/repository"
	"goServerAuth/structures"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	salt = "kj8932jfgj74thjg78p"
	signingKey = "jkwe823hu095tjioge904tr"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Autorization
}

func NewAuthService(repo repository.Autorization) *AuthService{
	return &AuthService{repo: repo}
}

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

//Изменить данные пользователя по id
func (s *AuthService) EditUser(user structures.User) (error) {
	return s.repo.EditUser(user)
}

//Удалить пользователя по id
func (s *AuthService) DeleteUser(idUser int, idAmin int) (error) {
	return s.repo.DeleteUser(idUser, idAmin)
}

//Генерация токена и проверка логина и пароля
func (s *AuthService) GenerateToken(login, password string) (string, error) {
	user, err := s.repo.Auth(login, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
		ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
		IssuedAt: time.Now().Unix(),
		},
		user.Id,
		
	})
	
	return token.SignedString([]byte(signingKey))
}

//Парсинг токена в хэдарах
func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type")
	}

	return claims.UserId, nil
}

//хэширования пароля
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
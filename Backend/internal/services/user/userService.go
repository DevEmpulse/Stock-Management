package services

import (
	"errors"
	"time"

	users "github.com/DevEmpulse/Stock-Management.git/internal/models/users"
	repositories "github.com/DevEmpulse/Stock-Management.git/internal/repositories/user"
	"github.com/dgrijalva/jwt-go" // Importa jwt-go
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("secreto") // Cambia el secreto por uno más seguro

type AuthService struct {
	UserRepo *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{UserRepo: userRepo}
}

// Registrar nuevo usuario
func (s *AuthService) Register(user *users.Users) error {
	user.Nickname = user.Email[:len(user.Email)-len("@example.com")]
	// Encriptar la contraseña con bcrypt
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Contrasena), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Contrasena = string(passwordHash)

	return s.UserRepo.CreateUser(user)
}

// Iniciar sesión con email y contraseña
func (s *AuthService) Login(email, password string) (*users.Users, string, error) {
	// Buscar usuario por email
	user, err := s.UserRepo.FindUserByEmail(email)
	if err != nil {
		return nil, "", errors.New("credenciales inválidas")
	}

	// Comparar la contraseña ingresada con la almacenada
	if err := bcrypt.CompareHashAndPassword([]byte(user.Contrasena), []byte(password)); err != nil {
		return nil, "", errors.New("credenciales inválidas")
	}

	// Crear token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID_users,
		"exp": time.Now().Add(time.Hour * 72).Unix(), // El token expirará en 72 horas
	})

	// Firmar el token con el secreto
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return nil, "", err
	}

	// Retornar el usuario y el token
	return user, tokenString, nil
}

// Eliminar usuario
func (s *AuthService) DeleteUser(id uint) error {
	return s.UserRepo.DeleteUser(id)
}

// Actualizar usuario
func (s *AuthService) UpdateUser(id uint, updatedUser *users.Users) error {
	return s.UserRepo.UpdateUser(id, updatedUser)
}

package service

import (
	"log"

	"github.com/ibnu-a/Golang-RESTful-API/dto"
	"github.com/ibnu-a/Golang-RESTful-API/entity"
	"github.com/ibnu-a/Golang-RESTful-API/repository"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email, password string) interface{}
	CreateUser(user dto.RegisterDTO) entity.User
	FindByEmail(email string) entity.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

// wajib melakukan implemntasi pada AuthService
func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

func (service *authService) VerifyCredential(email, password string) interface{} {
	res := service.userRepository.VerifyCredential(email, password)
	if v, ok := res.(entity.User); ok {
		comparePwd := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparePwd {
			return res
		}
		return false
	}
	return false
}

func (service *authService) CreateUser(user dto.RegisterDTO) entity.User {

	// userToCreate := entity.User{
	// 	Name: user.Name,
	// 	Password: user.Password,
	// 	Email: user.Email,
	// }

	userToCreate := entity.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.userRepository.InsertUser(userToCreate)
	return res
}

func (service *authService) FindByEmail(email string) entity.User {
	return service.userRepository.FindByEmail(email)
}

func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

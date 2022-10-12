package usecase

import (
	"assignment-golang-backend/customerrors.go"
	"assignment-golang-backend/entity"
	"assignment-golang-backend/repository"
	"assignment-golang-backend/utils"
)

type UserUsecase interface {
	Login(email, password string) (*entity.Token, error)
	Register(e *entity.User) (*entity.Token, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (u *userUsecase) Login(email, password string) (*entity.Token, error) {
	user, rowsAffected, err := u.userRepository.GetUserByEmail(email)

	//validation 1: if user not exist
	if rowsAffected == 0 {
		return nil, &customerrors.NoDataFoundError{}
	}
	//validation 2: if password not match
	if !utils.ComparePassword(user.Password, password) {
		return nil, &customerrors.WrongPasswordError{}
	}

	if err != nil {
		return nil, err
	}

	userToken := &entity.UserToken{
		ID:    user.WalletID,
		Name:  user.Name,
		Email: user.Email,
	}

	token, err := utils.GenerateJWT(userToken)
	if err != nil {
		return nil, err
	}

	return &entity.Token{TokenID: token}, nil

}

func (u *userUsecase) Register(e *entity.User) (*entity.Token, error) {
	var err error
	if e.Name == "" {
		return nil, &customerrors.InputEmptyError{}
	}
	if e.Email == "" {
		return nil, &customerrors.InputEmptyError{}
	}
	if e.Password == "" {
		return nil, &customerrors.InputEmptyError{}
	}

	isExist, err := u.userRepository.IsUserExist(e.Email)
	if err != nil {
		return nil, err
	}
	if isExist {
		return nil, &customerrors.DataAlreadyExistError{}
	}

	e.Password, err = utils.HashAndSalt(e.Password)
	if err != nil {
		return nil, err
	}

	newUser, rowsAffected, err := u.userRepository.CreateUser(e)
	if rowsAffected == 0 {
		return nil, &customerrors.NoDataUpdatedError{}
	}
	if err != nil {
		return nil, err
	}

	userToken := &entity.UserToken{
		ID:    newUser.WalletID,
		Name:  newUser.Name,
		Email: newUser.Email,
	}

	token, err := utils.GenerateJWT(userToken)
	if err != nil {
		return nil, err
	}

	return &entity.Token{TokenID: token}, nil
}

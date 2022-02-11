package usermodule

import (
	"context"
	"hienviluong125/trello-clone-be/common"
	"hienviluong125/trello-clone-be/component"
	"hienviluong125/trello-clone-be/errorhandler"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Signup(ctx context.Context, userCreate *UserCreate) error
	Login(ctx context.Context, userLogin *UserLogin) (*string, *string, error)
	RefreshCredentials(ctx context.Context, rfToken string) (*string, *string, error)
}

type UserDefaultService struct {
	repo       UserRepo
	appContext component.AppContext
}

func NewUserDefaultService(repo UserRepo, appContext component.AppContext) *UserDefaultService {
	return &UserDefaultService{repo: repo, appContext: appContext}
}

func (service *UserDefaultService) Signup(ctx context.Context, userCreate *UserCreate) error {
	existedUser, err := service.repo.FindByCondition(ctx, map[string]interface{}{"email": userCreate.Email})

	if existedUser != nil {
		return errorhandler.ErrRecordExisted("User", err)
	}

	if err := userCreate.Validate(); err != nil {
		return errorhandler.ErrInvalidRecord("User", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userCreate.Password), service.appContext.GetBcryptCost())

	if err != nil {
		return errorhandler.ErrInternal(err)
	}

	userCreate.Password = string(hashedPassword)
	userCreate.Status = true
	userCreate.Role = "member"

	if err := service.repo.Create(ctx, userCreate); err != nil {
		return err
	}

	return nil
}

func (service *UserDefaultService) Login(ctx context.Context, userLogin *UserLogin) (*string, *string, error) {
	if err := userLogin.Validate(); err != nil {
		return nil, nil, errorhandler.ErrUnauthorized(err)
	}

	user, err := service.repo.FindByCondition(ctx, map[string]interface{}{"email": userLogin.Email})

	if err != nil {
		return nil, nil, errorhandler.ErrUnauthorized(err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(userLogin.Password)); err != nil {
		return nil, nil, errorhandler.ErrUnauthorized(err)
	}

	var defaultTokenProvider common.TokenProvider = common.NewDefaultTokenProvider(service.appContext.GetSecretKey())

	accessToken, err := defaultTokenProvider.GenAccessToken(user.Id, 15)

	if err != nil {
		return nil, nil, err
	}

	refreshToken, err := defaultTokenProvider.GenRefreshToken()

	if err != nil {
		return nil, nil, err
	}

	err = service.repo.UpdateById(ctx, user.Id, &UserUpdate{RefreshToken: refreshToken})

	if err != nil {
		return nil, nil, err
	}

	return &accessToken, &refreshToken, nil
}

func (service *UserDefaultService) RefreshCredentials(ctx context.Context, rfToken string) (*string, *string, error) {
	user, err := service.repo.FindByCondition(ctx, map[string]interface{}{"refresh_token": rfToken})

	if err != nil {
		return nil, nil, err
	}

	var defaultTokenProvider common.TokenProvider = common.NewDefaultTokenProvider(service.appContext.GetSecretKey())

	accessToken, err := defaultTokenProvider.GenAccessToken(user.Id, 15)

	if err != nil {
		return nil, nil, err
	}

	refreshToken, err := defaultTokenProvider.GenRefreshToken()

	if err != nil {
		return nil, nil, err
	}

	err = service.repo.UpdateById(ctx, user.Id, &UserUpdate{RefreshToken: refreshToken})

	if err != nil {
		return nil, nil, err
	}

	return &accessToken, &refreshToken, nil
}

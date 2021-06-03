package controller

import (
	"encoding/json"
	"fmt"
	"github.com/IkezawaYuki/bookshelf-go/src/domain/model"
	"github.com/IkezawaYuki/bookshelf-go/src/infrastructure/redis"
	"github.com/IkezawaYuki/bookshelf-go/src/usecase/inputport"
	"github.com/IkezawaYuki/bookshelf-go/src/usecase/outputport"
	"github.com/google/uuid"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"net/http"
	"os"
)

type AuthController struct {
	userInteractor inputport.UserInputPort
}

func NewAuthController(userInteractor inputport.UserInputPort) AuthController {
	return AuthController{userInteractor: userInteractor}
}

const Google = "google"

func (a *AuthController) Login(c outputport.Context) error {
	config := oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Endpoint:     google.Endpoint,
		RedirectURL:  os.Getenv("REDIRECT_URL"),
		Scopes: []string{
			"https://www.googleapis.com/auth/spreadsheets",
			"https://www.googleapis.com/auth/userinfo.email",
		},
	}
	loginURL := config.AuthCodeURL(
		os.Getenv("SECURITY_KEY"),
		oauth2.AccessTypeOffline,
		oauth2.ApprovalForce)
	return c.JSON(http.StatusOK, loginURL)
}

func (a *AuthController) Callback(c outputport.Context) error {
	provider, err := gomniauth.Provider(Google)
	if err != nil {
		_ = c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}

	code := c.QueryParam("code")
	credentials, err := provider.CompleteAuth(objx.MustFromURLQuery("code=" + code))
	if err != nil {
		_ = c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}

	user, err := provider.GetUser(credentials)
	if err != nil {
		_ = c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}

	bookshelfUser, err := a.userInteractor.FindUserByEmail(user.Email())
	if err != nil {
		_ = c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}

	if bookshelfUser == nil {
		err := fmt.Errorf("permission denied email: %v", user.Email())
		_ = c.JSON(http.StatusForbidden, err.Error())
		return err
	}

	uuID, err := uuid.NewRandom()
	if err != nil {
		_ = c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}

	key := uuID.String()
	token := model.Token{
		AccessToken:  fmt.Sprint(credentials.Get("access_token")),
		RefreshToken: fmt.Sprint(credentials.Get("refresh_token")),
		Email:        bookshelfUser.Email,
		UserID:       bookshelfUser.ID,
	}

	bytes, err := json.Marshal(token)
	if err != nil {
		_ = c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}

	if err := redis.Handler.Set(key, bytes); err != nil {
		_ = c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}

	return c.JSON(http.StatusOK, nil)
}

func (a *AuthController) Logout(c outputport.Context) error {
	if err := redis.Handler.Delete(c.Get("key").(string)); err != nil {
		_ = c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}
	return c.JSON(http.StatusOK, nil)
}

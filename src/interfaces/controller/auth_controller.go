package controller

import (
	"github.com/IkezawaYuki/bookshelf-go/src/usecase/inputport"
	"github.com/IkezawaYuki/bookshelf-go/src/usecase/outputport"
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
	return c.String(http.StatusOK, loginURL)
}

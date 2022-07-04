package account

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"portfolio/db"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func newConf(url string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("clientId"),
		ClientSecret: os.Getenv("clientSecret"),
		RedirectURL:  url,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

func Route(router *gin.Engine) {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Printf(".envを読み込めませんでした: %v", err)
		return
	}

	aRouter := router.Group("/auth")
	aRouter.Use(login)
	aRouter.GET("/", index)
}
func login(ctx *gin.Context) {
	code := ctx.Query("code")
	_, err := getSession(ctx)
	if err != nil {
		return
	}
	if code == "" && err != nil {
		conf := newConf("https://suzakutakumi.mydns.jp:8080/auth")
		url := conf.AuthCodeURL(os.Getenv("state"), oauth2.AccessTypeOffline)
		ctx.HTML(http.StatusOK, "login.html", gin.H{"url": url})
		return
	}
	ctx.Next()
}
func index(ctx *gin.Context) {
	session, err := getSession(ctx)
	var client *http.Client
	if err != nil {
		code := ctx.Query("code")
		conf := newConf("https://suzakutakumi.mydns.jp:8080/auth")
		tok, err := conf.Exchange(ctx, code)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}
		u, err := uuid.NewRandom()
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}
		db.Push("insert into account values(?,?,?)", u.String(), tok.RefreshToken, tok.AccessToken)
		setSession(ctx, u.String())
		client = conf.Client(ctx, tok)
	} else {
		conf := newConf("https://suzakutakumi.mydns.jp:8080/auth")
		var tokens []Session
		db.Select(&tokens, "select * from account where session=?", session)
		token := &oauth2.Token{
			AccessToken:  tokens[0].AccessToken,
			TokenType:    "bearer",
			RefreshToken: tokens[0].RefreshToken,
			Expiry:       time.Now().Add(time.Hour * 24 * 30 * 2),
		}
		tokenSource := conf.TokenSource(ctx, token)
		oauth2.ReuseTokenSource()
		tok, err := conf.Exchange(ctx, session)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}
		client = conf.Client(ctx, tok)
	}

	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.String(http.StatusOK, string(body))
}

func setSession(ctx *gin.Context, session string) {
	ctx.SetCookie("session", session, 3600*24*30, "/", "suzakutakumi.mydns.jp:8080", false, true)
}
func getSession(ctx *gin.Context) (string, error) {
	session, err := ctx.Cookie("session")
	if err != nil {
		return "", err
	}
	return session, nil
}

type Session struct {
	SessionNum   int    `db:"session"`
	RefreshToken string `db:"refresh"`
	AccessToken  string `db:"access"`
}

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/translate"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

type greetings struct {
	Name    string `query:"name" json:"name"`
	Lang    string `query:"lang" json:"lang"`
	Message string `json:"message"`
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", greet)
	e.Logger.Fatal(e.Start(":8080"))
}

func greet(c echo.Context) error {
	var g greetings
	if err := c.Bind(&g); err != nil {
		return err
	}

	tl, err := language.Parse(g.Lang)

	if err != nil {
		log.Fatal(err)
	}

	bgCtx := context.Background()
	ctx, cancel := context.WithTimeout(bgCtx, 10*time.Second)
	defer cancel()

	apiKey := os.Getenv("API_KEY")
	client, err := translate.NewClient(ctx, option.WithAPIKey(apiKey))

	if err != nil {
		log.Fatal(err)
	}

	m := fmt.Sprintf("Thank you from GIDS2023 %s!", g.Name)
	res, err := client.Translate(ctx,
		[]string{m},
		tl,
		&translate.Options{
			Source: language.English,
			Format: translate.Text,
		})

	if err != nil {
		log.Fatal(err)
	}

	if len(res) > 0 {
		g.Message = res[0].Text
		return c.JSON(http.StatusOK, g)
	} else {
		return fmt.Errorf("unable to translate %s", m)
	}
}

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humaecho"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/go-co-op/gocron/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"api-go/environment"
	"api-go/mutations"
	"api-go/routes/auth_api"
	"api-go/routes/home_api"
	"api-go/routes/users_api"
	"api-go/scripts/places"
	"api-go/utils/auth"
	"api-go/utils/cache"
	"api-go/utils/db"
	"api-go/utils/files"
)

type OptionsCLI struct {
	Port   int
	Env    string
	Server string
}

func main() {
	cli := humacli.New(func(hooks humacli.Hooks, options *OptionsCLI) {

		cfg, err := environment.SetEnvironment(options.Env)
		if err != nil {
			panic("Cannot load environment")
		}

		db.SetEntDB(cfg.Database)
		db.SetBunDB(cfg.Database)
		cache.SetCache(cfg.VALKEY)
		auth.SetSecretJWT(cfg.Secret)
		places.SetMapsClient(cfg.Google_Maps_API_Key)
		files.SetStorageGCP(cfg.Google_Credentials, cfg.Google_Project, cfg.Google_Bucket)

		mutations.UserHook()

		if options.Server == "api" {
			runApi(hooks, options)
		} else if options.Server == "cron" {
			runCron()
		}

	})
	cli.Run()
}

func runApi(hooks humacli.Hooks, options *OptionsCLI) {
	router := echo.New()
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	config := huma.DefaultConfig("Spectra", "1.0.0")
	if options.Env == "prod" {
		config.OpenAPIPath = ""
		config.DocsPath = ""
	}

	api := humaecho.New(router, config)

	auth_api.Register(api)
	home_api.Register(api)
	users_api.Register(api)

	hooks.OnStart(func() {
		options.Port = 8080
		http.ListenAndServe(fmt.Sprintf(":%d", options.Port), router)
	})
}

func runCron() {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	_, err = scheduler.NewJob(
		gocron.DurationJob(
			30*time.Second,
		),
		gocron.NewTask(
			func(a string, b int) {
				fmt.Println("Cron Job")
				fmt.Println(a)
				fmt.Println(b)
			},
			"Test",
			1,
		),
	)
	if err != nil {
		panic(err)
	}

	scheduler.Start()
}

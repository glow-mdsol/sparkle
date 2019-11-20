package main

import (
	"github.com/glow-mdsol/sparkle/models"
	"github.com/jinzhu/gorm"
	"github.com/nu7hatch/gouuid"
	log "github.com/sirupsen/logrus"
	"net/http"

	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr"
	"github.com/spf13/viper"
	"os"
)

// App struct.
// There is no singleton anti-pattern,
// all variables defined locally inside
// this struct.
type App struct {
	Engine *gin.Engine
	Conf   *viper.Viper
	React  *React
	API    *API
	DB     *gorm.DB
}

// NewApp returns initialized struct
// of main server application.
func NewApp(opts ...AppOptions) *App {
	options := AppOptions{}
	for _, i := range opts {
		options = i
		break
	}

	options.init()
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)

	envVar := os.Getenv("ENV")

	if envVar == "" {
		envVar = "development"
	}
	viper.SetEnvPrefix("spk")
	viper.AutomaticEnv()
	viper.Set("ENV", envVar)

	viper.SetConfigName("config-" + envVar)
	viper.AddConfigPath("server/config/")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// Make an engine
	engine := gin.New()

	// Logger for each request
	engine.Use(gin.Logger())

	// Recovery for any panics issue in system
	engine.Use(gin.Recovery())

	engine.GET("/favicon.ico", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/static/images/favicon.ico")
	})

	engine.LoadHTMLGlob("server/data/templates/*")

	// Initialise the database
	datastoreType := viper.GetString("db_type")
	datastoreConnection := viper.GetString("db_connection")
	db := models.Database(datastoreType, datastoreConnection)

	// Initialize the application
	app := &App{
		Conf:   viper.GetViper(),
		Engine: engine,
		API:    &API{SPARQLServerName: viper.GetString("sparql_endpoint")},
		React: NewReact(
			viper.GetString("duktape.path"),
			viper.GetBool("debug"),
			engine,
		),
		DB: db,
	}

	// Map app and uuid for every requests
	app.Engine.Use(func(c *gin.Context) {
		c.Set("app", app)
		id, _ := uuid.NewV4()
		c.Set("uuid", id)
		c.Next()
	})

	// Bind api handling for URL api.prefix
	app.API.Bind(
		app.Engine.Group(
			app.Conf.GetString("api.prefix"),
		),
	)

	box := packr.NewBox("./data")

	// Create file http server from bindata
	fileServerHandler := http.FileServer(box)

	//Serve static via bindata and handle via react app
	//in case when static file was not found
	engine.NoRoute(func(c *gin.Context) {

		if _, err := box.MustBytes(c.Request.URL.Path[1:]); err == nil {
			log.Info("Handling ",c.Request.URL," through BOX")
			fileServerHandler.ServeHTTP(
				c.Writer,
				c.Request)
			return
		}
		log.Info("Handling ", c.Request.URL, " through REACT")
		app.React.Handle(c)
	})

	return app
}

// Run runs the app
func (app *App) Run() {
	Must(app.Engine.Run(":" + app.Conf.GetString("port")))
}

// AppOptions is options struct
type AppOptions struct{}

func (ao *AppOptions) init() { /* write your own*/ }

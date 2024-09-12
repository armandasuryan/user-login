package apps

import (
	db "auth/backend/config"
	"auth/backend/handler"
	"auth/backend/middleware"
	"auth/backend/repository"
	"auth/backend/routes"
	"auth/backend/services"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func StartApps() {
	app := fiber.New()

	// setup handle panic
	app.Use(middleware.CustomRecoverMiddleware)

	// setup logger
	app.Use(logger.New())
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:  "2006/01/02 15:04:05",
		DisableTimestamp: false,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "@level",
			logrus.FieldKeyMsg:   "@message",
			logrus.FieldKeyFunc:  "@caller",
		},
	})
	log := logrus.New()
	log.SetOutput(os.Stdout)

	// setup corse
	app.Use(cors.New())
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Set("Access-Control-Allow-Credentials", "true") // Optional
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusOK)
		}
		return c.Next()
	})

	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// setup db
	mysql := setupMySQLConnection()

	authSetting := setupAuth(mysql, log)
	authRouteConfig := routes.AuthRoute{
		App:         app,
		AuthHandler: authSetting,
	}
	authRouteConfig.SetupAuthRoute()

	errApp := app.Listen(":8080")
	if errApp != nil {
		log.Fatalf("Error starting Fiber app: %v", errApp)
	}
}

func setupMySQLConnection() *gorm.DB {
	hostMysql := os.Getenv("DB_HOST")
	usernameMysql := os.Getenv("DB_USERNAME")
	passwordMysql := os.Getenv("DB_PASSWORD")
	dbMysql := os.Getenv("DB_NAME")

	return db.MysqlConnect(hostMysql, usernameMysql, passwordMysql, dbMysql)
}

func setupAuth(mysql *gorm.DB, log *logrus.Logger) *handler.AuthHandlerMethod {
	repo := repository.AuthRepo(mysql, log)
	svc := services.AuthService(repo, log)
	return handler.AuthHandler(svc, log)
}

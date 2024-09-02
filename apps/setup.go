package apps

import (
	"os"
	db "user-service/backend/config"
	"user-service/backend/handler"
	"user-service/backend/middleware"
	"user-service/backend/repository"
	"user-service/backend/routes"
	"user-service/backend/services"

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

	userSetting := setupUser(mysql, log)
	userRouteConfig := routes.UserRoute{
		App:         app,
		UserHandler: userSetting,
	}
	userRouteConfig.SetupUserRoute()

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

func setupUser(mysql *gorm.DB, log *logrus.Logger) *handler.UserHandlerMethod {
	repo := repository.UserRepo(mysql, log)
	svc := services.UserService(repo, log)
	return handler.UserHandler(svc, log)
}

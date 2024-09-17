package db

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func MysqlConnect(host string, username string, password string, db_name string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=UTC",
		username, password, host, db_name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // This instructs GORM to not pluralize table names
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("Failed to connect to the database")
	} else {
		connectFigure := figure.NewColorFigure("Connect to Mysql", "", "yellow", true)
		connectFigure.Print()
	}

	return db
}

// func PostgresConnect(host string, username string, password string, db_name string, port string) *gorm.DB {
// 	// Format DSN untuk PostgreSQL
// 	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
// 		host, username, password, db_name, port)

// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
// 		NamingStrategy: schema.NamingStrategy{
// 			SingularTable: true, // Menginstruksikan GORM untuk tidak membuat nama tabel menjadi jamak
// 		},
// 		Logger: logger.Default.LogMode(logger.Info),
// 	})
// 	if err != nil {
// 		fmt.Println("Failed to connect to the database")
// 	} else {
// 		connectFigure := figure.NewColorFigure("Connect to PostgreSQL", "", "green", true)
//      connectFigure.Print()
// 	}

// 	return db
// }

// func MongoDBConnect(uri string) *mongo.Client {
// 	// setting timeout
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	// create client option
// 	clientOptions := options.Client().ApplyURI(uri)

// 	// connect to mongo
// 	client, err := mongo.Connect(ctx, clientOptions)
// 	if err != nil {
// 		fmt.Println("Failed to connect to MongoDB:", err)
// 		return nil
// 	}

// 	// test ping to mongo db
// 	err = client.Ping(ctx, readpref.Primary())
// 	if err != nil {
// 		fmt.Println("Failed to ping MongoDB:", err)
// 		return nil
// 	}

// 	fmt.Println("Connected to MongoDB")
// 	return client
// }

package database

import (
	"app/internal/config"
	"app/pkg/database/ent"
	_ "app/pkg/database/ent/runtime"
	rd "app/pkg/database/redis"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"entgo.io/ent/dialect"
	esql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/redis/go-redis/v9"
)

type DBClient struct {
	Ent   *ent.Client
	Redis *redis.Client
}

var ctx = context.Background()

func newDBClient() (*ent.Client, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Tối ưu connection pool
	db.SetMaxOpenConns(25)                 // Max connections
	db.SetMaxIdleConns(10)                 // Max idle connections
	db.SetConnMaxLifetime(5 * time.Minute) // Connection lifetime
	db.SetConnMaxIdleTime(2 * time.Minute) // Idle timeout

	drv := esql.OpenDB(dialect.MySQL, db)
	return ent.NewClient(ent.Driver(drv)), err
}

func DBConnect() *DBClient {
	client, err := newDBClient()
	rdb := rd.NewRedisClient()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Connected to MySQL database: %s:%d\n", config.DBHost, config.DBPort)
	fmt.Printf("Connected to Redis database: %s:%d\n", config.RedisHost, config.RedisPort)
	// check if the Redis client is connected
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatal(err)
	}

	return &DBClient{
		Ent:   client,
		Redis: rdb,
	}
}

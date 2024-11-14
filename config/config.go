package config

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/go-redis/redis/v8"
    "context"
)

var (
    DB  *gorm.DB
    RDB *redis.Client
)

func Initialize() {
    // Load environment variables
    if err := godotenv.Load("/app/.env"); err != nil {
        log.Fatal("Error loading .env file")
    }

    // Build PostgreSQL DSN
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        dbHost, dbUser, dbPassword, dbName, dbPort)

    // Initialize GORM with PostgreSQL
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }
    fmt.Println("Connected to PostgreSQL with GORM")

    // Initialize Redis connection
    redisURL := os.Getenv("REDIS_URL")
    RDB = redis.NewClient(&redis.Options{
        Addr: redisURL,
    })

    // Test Redis connection
    if _, err := RDB.Ping(context.Background()).Result(); err != nil {
        log.Fatalf("Failed to connect to Redis: %v", err)
    }
    fmt.Println("Connected to Redis")
}

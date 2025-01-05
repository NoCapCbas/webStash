package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8" // Redis client
	_ "github.com/lib/pq"          // PostgreSQL driver

	"github.com/NoCapCbas/webStash/internal/common"
	"github.com/NoCapCbas/webStash/users"
)

// Configuration variables, temporarily hard coded
const (
	postgresHost     = "postgres"
	postgresPort     = 5432
	postgresUser     = "postgres"
	postgresPassword = "postgres"
	postgresDB       = "postgres"

	redisAddr = "redis:6379"
	redisPass = "" // Add Redis password if needed
	redisDB   = 0  // Redis database index
)

func main() {

	// PostgreSQL connection
	pgConn, err := connectPostgres()
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer pgConn.Close()
	fmt.Println("Connected to PostgreSQL!")

	// Redis connection
	redisClient := connectRedis()
	defer redisClient.Close()
	fmt.Println("Connected to Redis!")

	// Initialize publisher
	publisher := common.NewPublisher(redisClient)

	// Initialize subscriber
	subscriber := common.NewSubscriber(redisClient)

	// Initialize the repo
	userRepo := users.NewUserRepository(pgConn)

	// Initialize the service
	userService := users.NewUserService(userRepo, publisher)

	// Initialize the http handler
	userHandler := users.NewUserHandler(userService)

	// Initialize the event handler
	eventHandler := users.NewEventHandler(userService)

	// Subscribe to the user.created channel
	err = subscriber.Subscribe(context.Background(), "users")
	if err != nil {
		log.Fatalf("Failed to subscribe to users channel: %v", err)
	}

	// Start a goroutine to listen for messages
	subscriber.Start(context.Background(), eventHandler)

	// Set up event routes /{service}/events
	http.HandleFunc("/users/events", userHandler.EventTypesHandler)

	// Set up general user routes /{service}/{event}
	http.HandleFunc("/users/signup", userHandler.SignUpUserHandler)

	// Set up user specific routes /{service}/{event}/{user_id}
	http.HandleFunc("users/login/{id}", userHandler.LoginUserHandler)
	http.HandleFunc("users/update/{id}", userHandler.UpdateUserHandler)
	http.HandleFunc("users/deactivate/{id}", userHandler.DeactivateUserHandler)
	http.HandleFunc("users/reactivate/{id}", userHandler.ReactivateUserHandler)

	// Start the server
	log.Println("Starting server on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// connectPostgres establishes a connection to the PostgreSQL database.
func connectPostgres() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		postgresHost, postgresPort, postgresUser, postgresPassword, postgresDB)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening PostgreSQL connection: %w", err)
	}

	// Verify the connection
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging PostgreSQL: %w", err)
	}

	return db, nil
}

// connectRedis establishes a connection to the Redis server.
func connectRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPass, // no password set
		DB:       redisDB,   // use default DB
	})

	return client
}

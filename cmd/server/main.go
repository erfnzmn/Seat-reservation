package main

import (
	"fmt"
	"log"
	"time"
	"net/http"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"seat-reservation/pkg/rabbitmq"
	"seat-reservation/pkg/redisclient"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

type MySQLConfig struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	Database  string `mapstructure:"database"`
	Charset   string `mapstructure:"charset"`
	ParseTime bool   `mapstructure:"parseTime"`
	Loc       string `mapstructure:"loc"`
}

type RedisConfig struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type RabbitConfig struct {
	URL string `mapstructure:"url"`
}

type LoggingConfig struct {
	Level string `mapstructure:"level"`
}

type Config struct {
	Server   ServerConfig `mapstructure:"server"`
	MySQL    MySQLConfig  `mapstructure:"mysql"`
	Redis    RedisConfig  `mapstructure:"redis"`
	RabbitMQ RabbitConfig `mapstructure:"rabbitmq"`
	Logging  LoggingConfig `mapstructure:"logging"`
}

func loadConfig() (*Config, error) {
	v := viper.New()

	v.SetConfigName("configs")
	v.SetConfigType("yaml")
	v.AddConfigPath("configs") // folder

	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %w", err)
	}

	return &cfg, nil
}
func connectMySQL(cfg *Config) (*gorm.DB, error) {
    // ساخت DSN
    dsn := fmt.Sprintf(
        "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=%s",
        cfg.MySQL.User,
        cfg.MySQL.Password,
        cfg.MySQL.Host,
        cfg.MySQL.Port,
        cfg.MySQL.Database,
        cfg.MySQL.Loc,
    )

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, fmt.Errorf("gorm open error: %w", err)
    }

    sqlDB, err := db.DB()
    if err != nil {
        return nil, fmt.Errorf("sql db unwrap error: %w", err)
    }

    sqlDB.SetMaxOpenConns(25)
    sqlDB.SetMaxIdleConns(10)
    sqlDB.SetConnMaxLifetime(30 * time.Minute)

    if err := sqlDB.Ping(); err != nil {
        return nil, fmt.Errorf("mysql ping error: %w", err)
    }

    return db, nil
}

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	fmt.Println("Server will run on port:", cfg.Server.Port)
	fmt.Println("MySQL host:", cfg.MySQL.Host)

	db, err := connectMySQL(cfg)
	if err != nil {
		log.Fatalf("MySQL connection failed: %v", err)
	}
	_ = db
	log.Println("MySQL connected ✔")

	// rdb, err := redisclient.New(redisclient.Config{
	// 	Enabled:  true,
	// 	Addr:     cfg.Redis.Address,
	// 	Password: cfg.Redis.Password,
	// 	DB:       cfg.Redis.DB,
	// })
	// if err != nil {
	// 	log.Fatalf("Redis error: %v", err)
	// }
	// if rdb != nil {
	// 	log.Println("Redis connected ✔")
	// }

	// rb, err := rabbitmq.NewRabbitMQ(cfg.RabbitMQ.URL)
	// if err != nil {
	// 	log.Fatalf("RabbitMQ error: %v", err)
	// }
	// log.Println("Rabbitmq Connected!")
	// defer rb.Close()

    e := echo.New()
    e.HideBanner = true

    e.Use(middleware.Recover())
    e.Use(middleware.Logger())
    e.Use(middleware.CORS())
    e.Use(middleware.Secure())
	// ---------------------------
// API Versioning
// ---------------------------
api := e.Group("/api")
v1 := api.Group("/v1")

// Empty modules for now
showsGroup := v1.Group("/shows")
seatsGroup := v1.Group("/seats")
reservationsGroup := v1.Group("/reservations")
waitingGroup := v1.Group("/waiting")

log.Println("API routes initialized.")


    // Health Check
	e.GET("/healthz", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{
			"status": "ok",
			"time":   time.Now().UTC(),
		})
	})

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("Server is running on %s ...", addr)

	if err := e.Start(addr); err != nil {
		log.Fatalf("Server error: %v", err)
	}
	

}

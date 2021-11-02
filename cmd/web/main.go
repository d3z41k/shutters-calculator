package main

import (
	"context"
	"flag"
	"github.com/d3z41k/shutters-calculator/pkg/models/mysql"
	ms "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type application struct {
	errorLog   *log.Logger
	infoLog    *log.Logger
	categories *mysql.CategoryModel
	fabrics    *mysql.FabricsModel
	profiles   *mysql.ProfilesModel
}

func main() {
	addr := flag.String("addr", ":8080", "HTTP network address")
	dsn := flag.String("dsn", "web:pass@tcp(127.0.0.1:3306)/shutters_calculator?charset=utf8mb4&parseTime=True&loc=Local", "MySQL data source name")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		errorLog:   errorLog,
		infoLog:    infoLog,
		categories: &mysql.CategoryModel{DB: db},
		fabrics:    &mysql.FabricsModel{DB: db},
		profiles:   &mysql.ProfilesModel{DB: db},
	}

	srv := &http.Server{
		Addr:         *addr,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on %s", *addr)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errorLog.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}

func openDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(ms.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

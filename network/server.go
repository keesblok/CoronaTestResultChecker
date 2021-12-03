package network

import (
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"log"
	"net/http"
	"os"
	"time"
)

// StartServer Starts the server for the heroku webpage
//This server is needed to let heroku think that this is a webserver
//If the server is accessed within an hour, it will stay online. This is done with the use of kaffeine: https://kaffeine.herokuapp.com/
func StartServer(stopServer chan bool) {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()
	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Thanks for keeping me alive!")
	})

	srv := &http.Server{
		Addr: ":" + port,
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	select {
	case <- stopServer:
		// Restore default behavior on the interrupt signal and notify user of shutdown.
		log.Println("shutting down gracefully, press Ctrl+C again to force")

		// The context is used to inform the server it has 5 seconds to finish
		// the request it is currently handling
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("Server forced to shutdown: ", err)
		}

		log.Println("Server exiting")
		return
	}
}
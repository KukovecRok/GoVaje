package main

// obvezno package main!!

import (
	"context"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
	"todorokvaja1/API"
	"todorokvaja1/DB/MongoDB"
	"todorokvaja1/Logic"
)

func main() {

	err := sentry.Init(sentry.ClientOptions{
		Dsn: getEnv("SENTRY", "error"),
	})
	if err != nil {
		sentry.CaptureException(err)
		log.Fatalf("sentry.Init: %s", err)
	}

	defer sentry.Flush(2 * time.Second)

	sentry.CaptureMessage("Go vaja1 zalaufala") //Ob vklopu pošlji info "error"

	port, err := strconv.Atoi(getEnv("MONGO_PORT", "27017"))
	if err != nil {
		sentry.CaptureException(err)
		log.Printf("Sentry.init %s", err)
		return
	}
	//Kreiramo DB objekt, ga inicializiramo in z njim naredimo objekt Logic (Kreiramo MariaDB ampak ga v Logic vstavimo kot tip DB - interface)
	db := &MongoDB.MongoDB{
		User:          getEnv("MONGO_USER", "PridobljenoIzEnv"),
		Pass:          getEnv("MONGO_PASS", "PridobljenoIzEnv"),
		IP:            getEnv("MONGO_IP", "PridobljenoIzEnv"),
		Port:          port,
		Database:      getEnv("MONGO_DB", "PridobljenoIzEnv"),
		AuthDB:        getEnv("MONGO_AUTH_DB", "PridobljenoIzEnv"),
		AuthMechanism: getEnv("MONGO_AUTH_MECHANISM", "PridobljenoIzEnv"),
	}

	db.Init(context.Background())

	logic := Logic.NewController(db, []byte("i87i7tzfjhvkbjhlkizutizrfhvjb"))

	//Kreiramo naš router objekt
	var router Router
	router.engine = gin.Default()
	router.api = API.NewController(logic)
	router.engine.Use(gin.Logger())
	router.engine.Use(gin.Recovery())
	router.secret = []byte("i87i7tzfjhvkbjhlkizutizrfhvjb")

	//Registriramo HTTP REST API povezave
	err = router.registerRoutes()
	if err != nil {
		sentry.CaptureException(err)
		fmt.Println(err.Error())
		return
	}

	//Naredimo 2 kanala in enega od njih povežemo na sistemski exit signal
	quit := make(chan os.Signal, 0)
	done := make(chan bool, 0)
	signal.Notify(quit, os.Interrupt)

	//Definiramo port, handler in timeoute za HTTP server
	srv := &http.Server{
		Addr:         ":8000",
		Handler:      router.engine,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	//V ločenem threadu čakamo na exit signal in nato izklopimo http server
	go func() {

		<-quit

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			sentry.CaptureException(err)
			fmt.Println(err.Error())
		}
		close(done)
	}()

	//V ločenem threadu zaženemo HTTP server
	go func() {

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			sentry.CaptureException(err)
			fmt.Println(err.Error())
		}
		os.Exit(1)
	}()

	//Čakamo na konec izvajanja. Vsi deli programa so sedaj zagnani v ločenih threadih.
	//Dalje od tod pridemo samo če program izklopimo ali crashne
	<-done
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")
		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()

		log.Println(status)
	}
}

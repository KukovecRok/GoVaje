package Vaja1

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
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
	port, err := strconv.Atoi(getEnv("MONGO_PORT", "27017"))
	if err != nil {
		return
	}
	//Kreiramo DB objekt, ga inicializiramo in z njim naredimo objekt Logic (Kreiramo MariaDB ampak ga v Logic vstavimo kot tip DB - interface)
	db := &MongoDB.MongoDB{
		User:          getEnv("MONGO_USER", "root"),
		Pass:          getEnv("MONGO_PASS", "root"),
		IP:            getEnv("MONGO_IP", "localhost"),
		Port:          port,
		Database:      getEnv("MONGO_DB", "pts"),
		AuthDB:        getEnv("MONGO_AUTH_DB", "admin"),
		AuthMechanism: getEnv("MONGO_AUTH_MECHANISM", "SCRAM-SHA-256"),
	}

	db.Init(context.Background())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	logic := Logic.NewController(db)

	//Kreiramo naš router objekt
	var router Router
	router.engine = gin.Default()
	router.api = API.NewController(logic)

	//Registriramo HTTP REST API povezave
	err = router.registerRoutes()
	if err != nil {
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
			fmt.Println(err.Error())
		}
		close(done)

	}()

	//V ločenem threadu zaženemo HTTP server
	go func() {

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
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

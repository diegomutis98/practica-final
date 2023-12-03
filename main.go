package main

import (
	"log"
	"net/http"

	"github.com/diegomutis98/practica-final/controllers"
	myhandlers "github.com/diegomutis98/practica-final/handlers"
	"github.com/diegomutis98/practica-final/models"
	repositorio "github.com/diegomutis98/practica-final/repository"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

func ConectarDB(url, driver string) (*sqlx.DB, error) {
	pgUrl, _ := pq.ParseURL(url)
	db, err := sqlx.Connect(driver, pgUrl)
	if err != nil {
		log.Printf("fallo la conexion a PostgreSQL, error: %s", err.Error())
		return nil, err
	}

	log.Printf("Nos conectamos bien a la base de datos db: %#v", db)
	return db, nil
}

func main() {

	db, err := ConectarDB("postgres://ilhwtjwn:MFGblZOQftrdWnJWE7LiRBocjzsjq49n@batyr.db.elephantsql.com/ilhwtjwn", "postgres")
	if err != nil {
		log.Fatalln("error conectando a la base de datos", err.Error())
		return
	}

	repo, err := repositorio.NewRepository[models.Estudiante](db)
	if err != nil {
		log.Fatalln("fallo al crear una instancia de repositorio", err.Error())
		return
	}

	controller, err := controllers.NewController(repo)
	if err != nil {
		log.Fatalln("fallo al crear una instancia de controller", err.Error())
		return
	}

	handler, err := myhandlers.NewHandler(controller)
	if err != nil {
		log.Fatalln("fallo al crear una instancia de handler", err.Error())
		return
	}
	router := mux.NewRouter()

	router.Handle("/estudiantes", http.HandlerFunc(handler.LeerEstudiantes)).Methods(http.MethodGet)
	router.Handle("/estudiantes", http.HandlerFunc(handler.CrearEstudiante)).Methods(http.MethodPost)
	router.Handle("/estudiantes/{id}", http.HandlerFunc(handler.LeerUnEstudiante)).Methods(http.MethodGet)
	router.Handle("/estudiantes/{id}", http.HandlerFunc(handler.ActualizarEstudiante)).Methods(http.MethodPatch)
	router.Handle("/estudiantes/{id}", http.HandlerFunc(handler.EliminarEstudiante)).Methods(http.MethodDelete)

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router))

}

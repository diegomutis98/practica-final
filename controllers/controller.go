package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/diegomutis98/practica-final/models"
	repositorio "github.com/diegomutis98/practica-final/repository"
)

var (
	updateQuery = "UPDATE estudiantes SET %s WHERE id=:id;"
	deleteQuery = "DELETE FROM estudiantes WHERE id=$1;"
	selectQuery = "SELECT id, usuario, nombre, identidad, programa, semestre, situacion, creditos, nivel FROM estudiantes WHERE id=$1;"
	listQuery   = "SELECT id, usuario, nombre, identidad, programa, semestre, situacion, creditos, nivel FROM estudiantes LIMIT $1 OFFSET $2;"
	createQuery = "INSERT INTO estudiantes (usuario, nombre, identidad, programa, semestre, situacion, creditos, nivel) VALUES (:usuario, :nombre, :identidad, :programa, :semestre, :situacion, :creditos, :nivel) returning id;"
)

type Controller struct {
	repo repositorio.Repository[models.Estudiante]
}

func NewController(repo repositorio.Repository[models.Estudiante]) (*Controller, error) {
	if repo == nil {
		return nil, fmt.Errorf("se necesita un repositorio no nulo para instanciar un controlador")
	}
	return &Controller{
		repo: repo,
	}, nil
}

func (c *Controller) ActualizarEstudiante(reqBody []byte, id string) error {
	nuevosValoresEstudiante := make(map[string]interface{})
	err := json.Unmarshal(reqBody, &nuevosValoresEstudiante)
	if err != nil {
		log.Printf("fallo al actualizar a un estudiante, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar a un estudiante, con error: %s", err.Error())
	}

	if len(nuevosValoresEstudiante) == 0 {
		log.Printf("no se proporcionaron valores para actualizar")
		return fmt.Errorf("no se proporcionaron valores para actualizar")
	}
	query := construirUpdateQuery(nuevosValoresEstudiante)
	nuevosValoresEstudiante["id"] = id
	err = c.repo.Update(context.TODO(), query, nuevosValoresEstudiante)
	if err != nil {
		log.Printf("fallo al actualizar un estudiante, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un estudiante, con error: %s", err.Error())
	}
	return nil
}

func construirUpdateQuery(nuevosValores map[string]interface{}) string {
	columns := []string{}
	for key := range nuevosValores {
		if key != "id" {
			columns = append(columns, fmt.Sprintf("%s=:%s", key, key))
		}
	}
	columnsString := strings.Join(columns, ",")
	return fmt.Sprintf(updateQuery, columnsString)
}

func (c *Controller) EliminarEstudiante(id string) error {
	err := c.repo.Delete(context.TODO(), deleteQuery, id)
	if err != nil {
		log.Printf("fallo al eliminar un estudiante, con error: %s", err.Error())
		return fmt.Errorf("fallo al eliminar un estudiante, con error: %s", err.Error())
	}
	return nil
}

func (c *Controller) LeerUnEstudiante(id string) ([]byte, error) {
	estudiante, err := c.repo.Read(context.TODO(), selectQuery, id)
	if err != nil {
		log.Printf("fallo al leer un estudiante, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer un estudiante, con error: %s", err.Error())
	}

	estudiantesJson, err := json.Marshal(estudiante)
	if err != nil {
		log.Printf("fallo al leer un estudiante, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer un estudiante, con error: %s", err.Error())
	}
	return estudiantesJson, nil
}

func (c *Controller) LeerEstudiantes(limit, offset int) ([]byte, error) {
	estudiantes, _, err := c.repo.List(context.TODO(), listQuery, limit, offset)
	if err != nil {
		log.Printf("fallo al leer a los estudiantes, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer a los estudiantes, con error: %s", err.Error())
	}

	jsonEstudiantes, err := json.Marshal(estudiantes)
	if err != nil {
		log.Printf("fallo al leer a los estudinates, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer a los estudiantes, con error: %s", err.Error())
	}
	return jsonEstudiantes, nil
}

func (c *Controller) CrearEstudiante(reqBody []byte) (int64, error) {
	nuevoEstudiante := &models.Estudiante{}
	err := json.Unmarshal(reqBody, nuevoEstudiante)

	log.Println(err)

	if err != nil {
		log.Printf("fallo al crear un nuevo estudiante, con error: %s", err.Error())
		return 0, fmt.Errorf("fallo al crear un nuevo estdiante, con error: %s", err.Error())
	}
	valoresColumnasNuevoEstudiante := map[string]any{

		"usuario":   nuevoEstudiante.Usuario,
		"nombre":    nuevoEstudiante.Nombre,
		"identidad": nuevoEstudiante.Identidad,
		"programa":  nuevoEstudiante.Programa,
		"semestre":  nuevoEstudiante.Semestre,
		"situacion": nuevoEstudiante.Situacion,
		"creditos":  nuevoEstudiante.Creditos,
		"nivel":     nuevoEstudiante.Nivel,
	}

	nuevoID, err := c.repo.Create(context.TODO(), createQuery, valoresColumnasNuevoEstudiante)
	if err != nil {
		log.Printf("fallo al crear un nuevo estudiante, con error: %s", err.Error())
		return 0, fmt.Errorf("fallo al crear un nuevo estudiante, con error: %s", err.Error())
	}

	return nuevoID, nil

}

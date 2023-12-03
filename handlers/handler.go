package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/diegomutis98/practica-final/controllers"

	"github.com/gorilla/mux"
)

type Handler struct {
	controller *controllers.Controller
}

func NewHandler(controller *controllers.Controller) (*Handler, error) {
	if controller == nil {
		return nil, fmt.Errorf("se requiere un controlador no nulo para instanciar un manejador")
	}
	return &Handler{
		controller: controller,
	}, nil
}

func (h *Handler) ActualizarEstudiante(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("fallo al actualizar un estudiante, con error: %s", err.Error())
		http.Error(writer, fmt.Sprintf("fallo al actualizar un estudiante, con error: %s", err.Error()), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	err = h.controller.ActualizarEstudiante(body, id)
	if err != nil {
		log.Printf("fallo al actualizar un estudiante, con error: %s", err.Error())
		http.Error(writer, fmt.Sprintf("fallo al actualizar un estudiante, con error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (h *Handler) EliminarEstudiante(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	err := h.controller.EliminarEstudiante(id)
	if err != nil {
		log.Printf("fallo al eliminar un estudiante, con error: %s", err.Error())
		http.Error(writer, fmt.Sprintf("fallo al eliminar un estudiante, con error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (h *Handler) LeerUnEstudiante(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	estudiante, err := h.controller.LeerUnEstudiante(id)
	if err != nil {
		log.Printf("fallo al leer un estudiante, con error: %s", err.Error())
		http.Error(writer, fmt.Sprintf("fallo al leer un estudiante, con error: %s", err.Error()), http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(estudiante)
}

func (h *Handler) LeerEstudiantes(writer http.ResponseWriter, req *http.Request) {
	estudiantes, err := h.controller.LeerEstudiantes(100, 0)
	if err != nil {
		log.Printf("fallo al leer a los estudiantes, con error: %s", err.Error())
		http.Error(writer, fmt.Sprintf("fallo al leer a los estudiantes, con error: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(estudiantes)
}

func (h *Handler) CrearEstudiante(writer http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("fallo al leer el cuerpo de la solicitud, con error: %s", err.Error())
		http.Error(writer, fmt.Sprintf("fallo al leer el cuerpo de la solicitud, con error: %s", err.Error()), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	nuevoID, err := h.controller.CrearEstudiante(body)
	if err != nil {
		log.Printf("fallo al crear un estudiante, con error: %s", err.Error())
		http.Error(writer, fmt.Sprintf("fallo al crear un empleado 10000, con error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(fmt.Sprintf("ID del nuevo estudiante: %d", nuevoID)))
}

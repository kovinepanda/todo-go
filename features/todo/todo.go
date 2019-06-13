package todo

import (
	"github.com/kovinepanda/todo-go/internal/config"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Config struct {
	*config.Config
}

func New(configuration *config.Config) *Config {
	return &Config{configuration}
}

type Todo struct {
	Slug string `json:"slug"`
	Title string `json:"title"`
	Body string `json:"body"`
}

func (config *Config) Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{todoID}", GetATodo)
	router.Delete("/{todoID}", DeleteTodo)
	router.Post("/", CreateTodo)
	router.Get("/", GetAllTodos)
	return router
}

func GetATodo(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	todos := Todo{
		Slug: todoID,
		Title: "Hello todo",
		Body: "Hello todo body",
	}
	render.JSON(w, r, todos)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Deleted TODO successfully"
	render.JSON(w, r, response)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Created TODO successfully"
	render.JSON(w, r, response)
}

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos := []Todo{
		{
			Slug: "slug",
			Title: "Hello todo",
			Body: "Hello todo body",
		},
	}
	render.JSON(w, r, todos)
}


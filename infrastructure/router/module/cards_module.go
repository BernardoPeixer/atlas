package module

import (
	"atlas/domain"
	"atlas/domain/entities"
	"atlas/infrastructure/router"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type moduleCards struct {
	config  entities.Config
	useCase domain.CardsUseCase
	name    string
	path    string
	host    string
}

func NewModuleCards(useCase domain.CardsUseCase, config entities.Config) router.Module {
	return moduleCards{
		useCase: useCase,
		name:    "Cards module",
		path:    "/cards",
		config:  config,
	}
}

func (m moduleCards) Name() string { return m.name }

func (m moduleCards) Path() string { return m.path }

func (m moduleCards) Setup(r *mux.Router) *mux.Router {
	handlers := []router.ModuleHandler{
		{
			Handler: m.listAllCards,
			Path:    "/list",
			Label:   "List all cards",
			Methods: []string{http.MethodGet},
		},
	}

	for _, h := range handlers {
		r.HandleFunc(h.Path, h.Handler).Methods(h.Methods...)
	}

	return r
}

func (m moduleCards) listAllCards(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := m.useCase.ListAllCards(ctx)
	if err != nil {
		log.Printf("Error in [ListAllCards]: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

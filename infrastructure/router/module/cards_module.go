package module

import (
	"atlas/domain"
	"atlas/domain/entities"
	"atlas/infrastructure/router"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
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
		{
			Handler: m.listAllCryptoType,
			Path:    "/cryptoType/list",
			Label:   "List all crypto type",
			Methods: []string{http.MethodGet},
		},
		{
			Handler: m.registerCard,
			Path:    "/register",
			Label:   "Register a card",
			Methods: []string{http.MethodPost},
		},
		{
			Handler: m.finishTransactionCard,
			Path:    "/finish/{cardID}",
			Label:   "Finalizes a transaction for the specified card ID.",
			Methods: []string{http.MethodPost},
		},
	}

	for _, h := range handlers {
		r.HandleFunc(h.Path, h.Handler).Methods(h.Methods...)
	}

	return r
}

func (m moduleCards) listAllCards(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	list, err := m.useCase.ListAllCards(ctx)
	if err != nil {
		log.Printf("Error in [ListAllCards]: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(list)
	if err != nil {
		log.Printf("Error in [Marshal]: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(response)
}

func (m moduleCards) listAllCryptoType(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	list, err := m.useCase.ListAllCryptoType(ctx)
	if err != nil {
		log.Printf("Error in [ListAllCryptoType]: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(list)
	if err != nil {
		log.Printf("Error in [Marshal]: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(response)
}

func (m moduleCards) registerCard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error in [ReadAll]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var festivalCard entities.FestivalCard
	err = json.Unmarshal(body, &festivalCard)
	if err != nil {
		log.Printf("Error in [Unmarshal]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = m.useCase.RegisterCard(ctx, festivalCard)
	if err != nil {
		log.Printf("Error in [RegisterCard]: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (m moduleCards) finishTransactionCard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)

	cardIDString := vars["cardID"]
	cardID, err := strconv.Atoi(cardIDString)
	if err != nil {
		log.Printf("Error in [Atoi]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = m.useCase.FinishTransactionCard(ctx, int64(cardID))
	if err != nil {
		log.Printf("Error in [FinishTransactionCard]: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

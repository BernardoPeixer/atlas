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
)

type moduleUser struct {
	config  entities.Config
	useCase domain.UserUseCase
	name    string
	path    string
	host    string
}

func NewModuleUser(useCase domain.UserUseCase, config entities.Config) router.Module {
	return moduleUser{
		useCase: useCase,
		name:    "User module",
		path:    "/user",
		config:  config,
	}
}

func (m moduleUser) Name() string { return m.name }

func (m moduleUser) Path() string { return m.path }

func (m moduleUser) Setup(r *mux.Router) *mux.Router {
	handlers := []router.ModuleHandler{
		{
			Handler: m.registerUser,
			Path:    "/register",
			Label:   "Register user in database",
			Methods: []string{http.MethodPost},
		},
		{
			Handler: m.checkUser,
			Path:    "/checkUser",
			Label:   "Validates if a user exists for the given wallet address",
			Methods: []string{http.MethodPost},
		},
	}

	for _, h := range handlers {
		r.HandleFunc(h.Path, h.Handler).Methods(h.Methods...)
	}

	return r
}

func (m moduleUser) registerUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error in [ReadAll]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user entities.UserInfo
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Printf("Error in [Unmarshal]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = m.useCase.RegisterUser(ctx, user)
	if err != nil {
		log.Printf("Error in [RegisterUser]: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (m moduleUser) checkUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error in [ReadAll]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var request struct {
		WalletAddress string `json:"wallet_address"`
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		log.Printf("Error in [Unmarshal]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	validUser, err := m.useCase.CheckUser(ctx, request.WalletAddress)
	if err != nil {
		log.Printf("Error in [CheckUser]: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseJson := struct {
		ValidUser bool `json:"valid_user"`
	}{
		ValidUser: validUser,
	}

	response, err := json.Marshal(responseJson)
	if err != nil {
		log.Printf("Error in [Marshal]: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(response)
}

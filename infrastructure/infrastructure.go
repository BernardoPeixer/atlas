package infrastructure

import (
	"atlas/domain/entities"
	"atlas/domain/usecases"
	"atlas/infrastructure/datastore"
	"atlas/infrastructure/datastore/repository"
	"atlas/infrastructure/router"
	"atlas/infrastructure/router/module"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func SetupModules(r *mux.Router, cfg entities.Config) func() error {
	log.Println("Start modules setup")

	// Settings repository
	settings := datastore.NewRepositorySettings(cfg)

	// =============================== REPOSITORIES ===============================
	// Cards repository
	cardsRepository := repository.NewCardsRepository(settings, cfg)

	// User repository
	userRepository := repository.NewUserRepository(settings, cfg)

	// ================================ USE CASES ==================================

	// Cards use case
	cardsUseCase := usecases.NewCardsUseCase(cardsRepository, cfg)

	// User use case
	userUseCase := usecases.NewUserUseCase(userRepository, cfg)

	// ================================= MODULES ====================================

	// Cards module
	cardsModule := module.NewModuleCards(cardsUseCase, cfg)

	// User module
	userModule := module.NewModuleUser(userUseCase, cfg)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		dt, err := settings.ServerTime(ctx)
		if err != nil {
			log.Printf("Error in [ServerTime]")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tm := time.Time(*dt)
		_, err = fmt.Fprintf(w, "%v", tm.UTC().Unix())

		if err != nil {
			log.Println(err)
		}
	})

	modules := []router.Module{
		cardsModule,
		userModule,
	}

	mainRouter := r.PathPrefix("/api").Subrouter()
	// mainRouter.Use(middleware) TODO: Implements middleware

	for _, m := range modules {
		moduleSubRouter := mainRouter.PathPrefix(m.Path()).Subrouter()
		_ = m.Setup(moduleSubRouter)
	}

	log.Printf("All modules up and running")
	return func() error {
		return settings.Dismount()
	}
}

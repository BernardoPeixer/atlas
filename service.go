package main

import (
	"atlas/domain/entities"
	"atlas/infrastructure"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/kardianos/service"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type program struct {
	server *http.Server
	cfg    entities.Config
}

func (p *program) Start(s service.Service) error {
	log.Println("received call to program#start")
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}

func (p *program) Stop(s service.Service) error {
	log.Println("received call to program#stop")

	// Stop should not block. Return with a few seconds
	return nil
}

func (p *program) run() {
	corsOptions := handlers.AllowedOriginValidator(func(s string) bool { return true })

	r := mux.NewRouter()
	_ = infrastructure.SetupModules(r, p.cfg)
	p.server = &http.Server{
		Addr: fmt.Sprintf(":%d", p.cfg.Port),
		Handler: handlers.CORS(
			corsOptions,
			handlers.AllowedHeaders([]string{"Authorization", "Content-Type", "Accept"}),
			handlers.AllowedMethods([]string{"GET", "PUT", "POST", "DELETE"}),
			handlers.AllowCredentials(),
		)(r),
		ReadTimeout:       time.Second * 120,
		ReadHeaderTimeout: time.Second * 15,
		WriteTimeout:      time.Minute * 7,
		IdleTimeout:       time.Second * 60,
	}

	log.Printf("Starting server on address: %s", p.server.Addr)
	if err := p.server.ListenAndServe(); err != nil {
		log.Printf("Error: %v", err)
	}
}

func loadFlags() (string, string, bool) {
	var cfgPath string
	var action string
	var terminal bool

	flag.StringVar(&cfgPath, "configs", "", "the path to the application config file")
	flag.StringVar(&action, "action", "", "the action to execute")
	flag.BoolVar(&terminal, "terminal", false, "display the logs in the terminal")

	flag.Parse()

	fmt.Printf("Config path recebido: %s\n", cfgPath)

	if cfgPath == "" {
		cfgPath = os.Getenv("configs")
	}

	if cfgPath == "" {
		panic("[configs] arguments not found")
	}

	fmt.Printf(cfgPath)
	fmt.Printf(action)

	return cfgPath, action, terminal
}

func readCFGFile(cfgPath string) *entities.Config {

	fmt.Printf(cfgPath)

	file, err := os.Open(cfgPath)
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo de configuração: %v", err)
	}

	read, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = file.Close()
	if err != nil {
		panic(err)
	}

	var cfg entities.Config
	_, err = toml.Decode(string(read), &cfg)
	if err != nil {
		log.Fatalf("Erro ao decodificar TOML: %v", err)
	}

	return &cfg
}

func configureOutput(logFolder string) (*os.File, error) {
	if logFolder == "" {
		return nil, nil
	}

	now := time.Now()
	logName := fmt.Sprintf("%s/%s.log", logFolder, now.Format("20060102150405"))

	file, err := os.OpenFile(logName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return nil, err
	}

	log.SetOutput(file)
	return file, nil
}

func newService(config entities.Config) (service.Service, error) {
	log.Println("Creating service")

	var args []string

	if len(os.Args) > 1 {
		for _, arg := range os.Args {
			if strings.Contains(arg, "configs") {
				args = append(args, arg)
			}
		}
	}

	serviceConfig := &service.Config{
		Name:        "atlas",
		DisplayName: "Webservice FestChain - Production",
		Description: "Server used for the production of the FestChain application",
		Arguments:   args,
	}

	programSvc := &program{
		cfg: config,
	}

	newService, err := service.New(programSvc, serviceConfig)
	if err != nil {
		return nil, err
	}

	return newService, nil
}

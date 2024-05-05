package main

import (
	"fmt"

	"context"
	"log"
	"net/http"

	"strconv"
	"time"

	"github.com/shubhamgoyal1402/hpe-golang-workflow/project/adapters/cadenceAdapter"
	"github.com/shubhamgoyal1402/hpe-golang-workflow/project/config"

	"github.com/shubhamgoyal1402/hpe-golang-workflow/project/worker/workflows"
	"go.uber.org/cadence/client"
	"go.uber.org/cadence/worker"
	"go.uber.org/zap"
)

const cadenceCLIImage = "ubercadence/cli:master"
const cadenceAddress = "host.docker.internal:7933"
const domain = "day56-domain"
const taskList = "Service_process"
const workflowType = "github.com/shubhamgoyal1402/hpe-golang-workflow/project/worker/workflows.CustomerWorkflow"

const (
	address = "localhost:50051"
)

type Service struct {
	cadenceAdapter *cadenceAdapter.CadenceAdapter
	logger         *zap.Logger
}

func (h *Service) formHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Parse form error : %v", err)
		return
	}

	fmt.Fprintf(w, "Request For %s Service Submitted\n", r.FormValue("name"))

	Name := r.FormValue("name")
	id := r.FormValue("service_id")
	requestid, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	if requestid < 1 || requestid > 6 {
		fmt.Fprintf(w, "Service not available for ID %s", id)
	} else {

		wo := client.StartWorkflowOptions{
			TaskList:                     taskList,
			ExecutionStartToCloseTimeout: time.Hour * 24,
		}

		_, err := h.cadenceAdapter.CadenceClient.StartWorkflow(context.Background(), wo, workflows.CustomerWorkflow, 1)

		if err != nil {
			http.Error(w, "Error starting workflow!", http.StatusBadRequest)
			return

		}
	}

	fmt.Fprintf(w, "Service Name = %s\n", Name)
	fmt.Fprintf(w, "Request ID= %s\n", id)
}

func startWorkers(h *cadenceAdapter.CadenceAdapter, taskList string) {
	// Configure worker options.
	workerOptions := worker.Options{
		MetricsScope: h.Scope,
		Logger:       h.Logger,
	}

	cadenceWorker := worker.New(h.ServiceClient, h.Config.Domain, taskList, workerOptions)
	err := cadenceWorker.Start()
	if err != nil {
		h.Logger.Error("Failed to start workers.", zap.Error(err))
		panic("Failed to start workers")
	}
}

func main() {

	fmt.Println("Starting Worker..")
	var appConfig config.AppConfig
	appConfig.Setup()
	var cadenceClient cadenceAdapter.CadenceAdapter
	cadenceClient.Setup(&appConfig.Cadence)

	startWorkers(&cadenceClient, taskList)

	fmt.Println("Cadence worker ready ")

	fileServer := http.FileServer(http.Dir("./static"))
	service := Service{&cadenceClient, appConfig.Logger}
	http.Handle("/", fileServer)

	http.HandleFunc("/form", service.formHandler)

	fmt.Printf("Starting server at 8080 port\n")

	err2 := http.ListenAndServe(":8080", nil)

	if err2 != nil {
		log.Fatal(err2)
	}

}

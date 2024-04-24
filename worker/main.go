package main

import (
	"fmt"

	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/shubhamgoyal1402/hpe-golang-workflow/project/adapters/cadenceAdapter"
	"github.com/shubhamgoyal1402/hpe-golang-workflow/project/config"
	"github.com/shubhamgoyal1402/hpe-golang-workflow/project/worker/workflows"

	"go.uber.org/cadence/worker"
	"go.uber.org/zap"
)

const cadenceCLIImage = "ubercadence/cli:master"
const cadenceAddress = "host.docker.internal:7933"
const domain = "day40-domain"
const taskList = "Service_process"
const workflowType = "github.com/shubhamgoyal1402/hpe-golang-workflow/project/worker/workflows.customerWorkflow"

var task_counter1 = 0
var task_counter2 = 0
var task_counter3 = 0
var c = 0
var a = 0
var b = 0

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

	go worker1()
	go worker2()
	go worker3()

	fmt.Println("All workers are readyy ")

	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)

	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server at 8080 port\n")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}

	// The workers are supposed to be long running process that should not exit.
	time.Sleep(time.Hour * 8)
	//select {}
}
func worker1() {

	for c > -1 {
		if task_counter1 < 5 {
			task_counter1++
			go PrivateCloudEnterpriseServiceProcessing()
		} else {
			time.Sleep(time.Second)
		}

	}
}
func worker2() {

	for a > -1 {
		if task_counter2 < 5 {

			task_counter2++
			go NetworkingServiceProcessing()
		} else {
			time.Sleep(time.Second)
		}

	}

}

func worker3() {

	for b > -1 {
		if task_counter3 < 5 {
			task_counter3++
			go BlockStorageServiceProcessing()
		} else {
			time.Sleep(time.Second)
		}

	}

}

func NetworkingServiceProcessing() {
	response2, _, _, err2 := workflows.Q2.Dequeue()

	if err2 == nil {

		time.Sleep(time.Second * 10)

		workflows.Response_Queue2.Enqueue2(response2)

		fmt.Println("signal sent to Response queue 2")
		time.Sleep(time.Millisecond)

	}
	task_counter2--

}

func BlockStorageServiceProcessing() {

	response3, _, _, err3 := workflows.Q3.Dequeue()

	if err3 == nil {

		time.Sleep(time.Second * 10)

		workflows.Response_Queue3.Enqueue2(response3)

		fmt.Println("signal sent to Response queue 3")
		time.Sleep(time.Millisecond)

	}
	task_counter3--
}

func PrivateCloudEnterpriseServiceProcessing() {

	response1, _, _, err1 := workflows.Q1.Dequeue()

	if err1 == nil {

		time.Sleep(time.Second * 10)

		workflows.Response_Queue1.Enqueue2(response1)

		fmt.Println("signal sent to Response queue 1")
		time.Sleep(time.Millisecond)

	}
	task_counter1--
}
func formHandler(w http.ResponseWriter, r *http.Request) {

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

	dockerCmd := fmt.Sprintf("docker run --rm %s --address %s -do %s workflow start --et 1000 --tl %s --wt %s --input %d", cadenceCLIImage, cadenceAddress, domain, taskList, workflowType, requestid)
	executeCommand(dockerCmd)

	fmt.Fprintf(w, "Service Name = %s\n", Name)
	fmt.Fprintf(w, "Request ID= %s\n", id)

}

func executeCommand(command string) {
	cmd := exec.Command("cmd", "/c", command)
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {

		os.Exit(1)
	}

}

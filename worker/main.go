package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/shubhamgoyal1402/hpe-golang-workflow/project/adapters/cadenceAdapter"
	"github.com/shubhamgoyal1402/hpe-golang-workflow/project/config"
	"github.com/shubhamgoyal1402/hpe-golang-workflow/project/worker/workflows"
	"go.uber.org/cadence/worker"
	"go.uber.org/zap"
)

const numWorkflows = 10
const cadenceCLIImage = "ubercadence/cli:master"
const cadenceAddress = "host.docker.internal:7933"
const domain = "day20-domain"
const taskList = "Service_buy_process"
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

	startWorkers(&cadenceClient, workflows.TaskListName)
	go worker1()
	go worker2()
	go worker3()

	time.Sleep(time.Second * 10)
	go start()

	fmt.Println("All workers are readyy ")
	// The workers are supposed to be long running process that should not exit.
	select {}
}

func start() {

	for i := 1; i <= numWorkflows; i++ {
		// Generate a random workflow input
		randomNumber := rand.Intn(6) + 1
		fmt.Println(randomNumber)
		dockerCmd := fmt.Sprintf("docker run --rm %s --address %s -do %s workflow start --et 1000 --tl %s --wt %s --input %d",
			cadenceCLIImage, cadenceAddress, domain, taskList, workflowType, randomNumber)
		executeCommand(dockerCmd)

		// Execute Docker command
		//	fmt.Println("Executing Docker command:", dockerCmd)

	}
}
func executeCommand(command string) {
	cmd := exec.Command("cmd", "/c", command)
	cmd.Stdout = os.Stdout

	//cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {

		//fmt.Printf("Error executing Docker command: %v\n", err)
		os.Exit(1)
	}

}
func worker2() {

	for a > -1 {
		if task_counter2 < 6 {

			task_counter2++
			go testing2()
		} else {
			time.Sleep(time.Millisecond)
		}

	}

}

func worker3() {

	for b > -1 {
		if task_counter3 < 6 {
			task_counter3++
			go testing3()
		} else {
			time.Sleep(time.Millisecond)
		}

	}

}
func worker1() {

	for c > -1 {
		if task_counter1 < 6 {
			task_counter1++
			go testing1()
		} else {
			time.Sleep(time.Millisecond)
		}

	}
}

func testing2() {

	response2, err2 := workflows.Q2.Dequeue()

	if err2 == nil {

		time.Sleep(time.Second * 10)

		workflows.Response_Queue2.Enqueue2(response2)

		fmt.Println("signal sent to Response queue 1")
		time.Sleep(time.Millisecond * 500)

	}
	task_counter2--
}

func testing3() {

	response3, err3 := workflows.Q3.Dequeue()

	if err3 == nil {

		time.Sleep(time.Second * 10)

		workflows.Response_Queue3.Enqueue2(response3)

		fmt.Println("signal sent to Response queue 3")
		time.Sleep(time.Millisecond * 500)

	}
	task_counter3--
}

func testing1() {

	response1, err1 := workflows.Q1.Dequeue()

	if err1 == nil {

		time.Sleep(time.Second * 10)

		workflows.Response_Queue1.Enqueue2(response1)

		fmt.Println("signal sent to Response queue 1")
		time.Sleep(time.Millisecond * 500)

	}
	task_counter1--
}

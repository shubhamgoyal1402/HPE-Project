package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

const cadenceCLIImage = "ubercadence/cli:master"
const cadenceAddress = "host.docker.internal:7933"
const domain = "day56-domain"
const taskList = "Service_process"
const workflowType = "github.com/shubhamgoyal1402/hpe-golang-workflow/project/worker/workflows.CustomerWorkflow"

type RequestBody struct {
	WorkID string `json:"work_id"`
	RunID  string `json:"run_id"`
}

func handleRequest1(w http.ResponseWriter, r *http.Request) {
	handleRequest(w, r, "Endpoint 1")
}

func handleRequest2(w http.ResponseWriter, r *http.Request) {
	handleRequest(w, r, "Endpoint 2")
}

func handleRequest3(w http.ResponseWriter, r *http.Request) {
	handleRequest(w, r, "Endpoint 3")
}

func handleRequest(w http.ResponseWriter, r *http.Request, endpoint string) {
	// Decode JSON request body into RequestBody struct
	var requestBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Accessing the values from the request body
	fmt.Printf("Received request at %s: WorkflowID=%s, RunID=%s\n", endpoint, requestBody.WorkID, requestBody.RunID)

	jsonSignal1, err1 := json.Marshal(requestBody.WorkID)
	if err1 != nil {
		log.Fatalf("Error marshaling signal to JSON: %v", err1)
	}
	signalcmd := fmt.Sprintf("docker run --rm %s --address %s -do %s workflow signal -w %s -r %s -n %s -i %s", cadenceCLIImage, cadenceAddress, domain, requestBody.WorkID, requestBody.RunID, requestBody.WorkID, string(jsonSignal1))
	go executeCommand(signalcmd)

	// Send a response
	fmt.Fprintf(w, "Received request at %s: WorkflowID=%s, RunID=%s\n", endpoint, requestBody.WorkID, requestBody.RunID)
}

func main() {

	http.HandleFunc("/endpoint1", handleRequest1)
	http.HandleFunc("/endpoint2", handleRequest2)
	http.HandleFunc("/endpoint3", handleRequest3)

	fmt.Println("Server listening on port 8090...")
	log.Fatal(http.ListenAndServe(":8090", nil))
}

func executeCommand(command string) {
	cmd := exec.Command("cmd", "/c", command)
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {

		os.Exit(1)
	}

}

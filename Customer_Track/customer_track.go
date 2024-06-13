package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"text/template"
)

const cadenceCLIImage = "ubercadence/cli:master"
const cadenceAddress = "host.docker.internal:7933"
const domain = "day56-domain"
const state = "current_state"

func main() {

	http.HandleFunc("/progress", formHandler)
	http.HandleFunc("/submit", submitHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Printf("Starting server at 9095 port\n")

	err2 := http.ListenAndServe(":9095", nil)

	if err2 != nil {
		log.Fatal(err2)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/progress.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	workflowID := r.FormValue("workflow_id")
	runID := r.FormValue("run_id")

	dockerCmd := fmt.Sprintf("docker run --rm %s --address %s --domain %s workflow query -w %s -r %s -qt %s", cadenceCLIImage, cadenceAddress, domain, workflowID, runID, state)
	ans := executeCommand(dockerCmd)

	fmt.Fprintf(w, "Current Status %s\n", ans)
}
func executeCommand(command string) string {
	cmd := exec.Command("cmd", "/c", command)

	//cmd.Stderr = os.Stderr

	output, err := cmd.CombinedOutput()
	if err != nil {

		return "INVALID WORKLFOW AUR RUN ID"

	}

	return string(output)

}

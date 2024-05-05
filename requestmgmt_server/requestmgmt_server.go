package main

import (
	"context"
	"log"

	"encoding/json"
	"errors"
	"fmt"

	"net"
	"net/http"
	"time"

	"github.com/shubhamgoyal1402/hpe-golang-workflow/project/Queue"
	pb "github.com/shubhamgoyal1402/hpe-golang-workflow/project/requestmgmt"
	"google.golang.org/grpc"

	"os"
	"os/exec"
)

var task_counter1 = 0
var task_counter2 = 0
var task_counter3 = 0
var c = 0
var a = 0
var b = 0

const cadenceCLIImage = "ubercadence/cli:master"
const cadenceAddress = "host.docker.internal:7933"
const domain = "day56-domain"
const taskList = "Service_process"
const workflowType = "github.com/shubhamgoyal1402/hpe-golang-workflow/project/worker/workflows.CustomerWorkflow"
const (
	port = ":50051"
)

type RequestManagementServer struct {
	pb.UnimplementedRequestManagementServer
}

var Q1 = Queue.Queue{

	Size: 10,
}

var Q2 = Queue.Queue{

	Size: 10,
}
var Q3 = Queue.Queue{

	Size: 10,
}

func sendProgrammaticRequest(url string) {
	client := &http.Client{}

	// Create a new request with the custom header
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creating new request: %s", err)
	}

	// Set the custom header to indicate a programmatic request
	req.Header.Set("Programmatic-Request", "true")

	// Send the request and handle the response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making GET request to %s: %s", url, err)
	}
	defer resp.Body.Close()

}

func (s *RequestManagementServer) CreateRequest(ctx context.Context, in *pb.NewRequest) (*pb.Request, error) {

	log.Printf("received: %v", in.GetWid())
	_, err := rpc1(ctx, in.GetWid(), in.GetRid(), in.GetId())

	if err != nil {
		log.Fatalf("error while enquing %v", err)
	}
	return &pb.Request{Wid: in.GetWid()}, nil

}
func rpc1(ctx context.Context, workflow_id string, rid string, id int32) (string, error) {
	switch id {
	case 1, 4:
		ans, err := rpc_function1(ctx, workflow_id, id, rid, &Q1)
		Q1.SortCustomers()
		return ans, err
	case 2, 5:
		ans, err := rpc_function1(ctx, workflow_id, id, rid, &Q2)
		Q2.SortCustomers()
		return ans, err
	case 3, 6:
		ans, err := rpc_function1(ctx, workflow_id, id, rid, &Q3)
		Q3.SortCustomers()
		return ans, err
	default:
		return "Error Service not available ", errors.ErrUnsupported
	}

}
func rpc_function1(ctx context.Context, workflow_id string, id int32, rid string, q *Queue.Queue) (string, error) {

	for q.GetLength() >= q.Size/2 {
		time.Sleep(time.Second)
	}

	customer1 := Queue.New(workflow_id, rid, ctx, id)
	ans := fmt.Sprintf("Enqueud at %s", time.Now())
	fmt.Println(ans, workflow_id)
	_, err := q.Enqueue(customer1)
	if err != nil {
		panic(err)
	}
	return ans, err
}

func main() {

	go worker1()
	go worker2()
	go worker3()

	log.Println("All workers ready ")
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterRequestManagementServer(s, &RequestManagementServer{})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve at: %v", err)
	}

	select {}
}

func worker1() {

	for c > -1 {
		if task_counter1 < 5 {
			task_counter1++
			go PrivateCloudEnterpriseServiceProcessing()
			time.Sleep(time.Millisecond * 100)
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
			time.Sleep(time.Millisecond * 100)
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
			time.Sleep(time.Millisecond * 100)
		} else {
			time.Sleep(time.Second)
		}

	}

}

func PrivateCloudEnterpriseServiceProcessing() {

	response1, runid1, _, err1 := Q1.Dequeue()

	if err1 == nil {
		sendProgrammaticRequest("http://localhost:8090/endpoint1")
		// time waiting for completion of service
		time.Sleep(time.Second * 10)
		// sending signal to workflow
		jsonSignal1, err1 := json.Marshal(response1)
		if err1 != nil {
			log.Fatalf("Error marshaling signal to JSON: %v", err1)
		}
		signalcmd := fmt.Sprintf("docker run --rm %s --address %s -do %s workflow signal -w %s -r %s -n %s -i %s", cadenceCLIImage, cadenceAddress, domain, response1, runid1, response1, string(jsonSignal1))
		go executeCommand(signalcmd)

		time.Sleep(time.Millisecond)

	}
	task_counter1--
}
func NetworkingServiceProcessing() {

	response2, runid2, _, err2 := Q2.Dequeue()

	if err2 == nil {
		sendProgrammaticRequest("http://localhost:8090/endpoint2")
		time.Sleep(time.Second * 10)

		jsonSignal2, err1 := json.Marshal(response2)
		if err1 != nil {
			log.Fatalf("Error marshaling signal to JSON: %v", err1)
		}
		signalcmd := fmt.Sprintf("docker run --rm %s --address %s -do %s workflow signal -w %s -r %s -n %s -i %s", cadenceCLIImage, cadenceAddress, domain, response2, runid2, response2, string(jsonSignal2))

		go executeCommand(signalcmd)

		time.Sleep(time.Millisecond)

	}
	task_counter2--

}

func BlockStorageServiceProcessing() {

	response3, runid3, _, err3 := Q3.Dequeue()

	if err3 == nil {
		sendProgrammaticRequest("http://localhost:8090/endpoint3")
		time.Sleep(time.Second * 10)

		jsonSignal3, err1 := json.Marshal(response3)
		if err1 != nil {
			log.Fatalf("Error marshaling signal to JSON: %v", err1)
		}
		signalcmd := fmt.Sprintf("docker run --rm %s --address %s -do %s workflow signal -w %s -r %s -n %s -i %s", cadenceCLIImage, cadenceAddress, domain, response3, runid3, response3, string(jsonSignal3))
		go executeCommand(signalcmd)

		time.Sleep(time.Millisecond)

	}
	task_counter3--
}

func executeCommand(command string) {
	cmd := exec.Command("cmd", "/c", command)
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {

		os.Exit(1)
	}

}

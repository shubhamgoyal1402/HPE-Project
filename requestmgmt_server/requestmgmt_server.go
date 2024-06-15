package main

import (
	"bytes"
	"context"
	"log"
	"sync"

	"encoding/json"
	"errors"
	"fmt"

	"net"
	"net/http"
	"time"

	"github.com/shubhamgoyal1402/hpe-golang-workflow/project/Queue"
	pb "github.com/shubhamgoyal1402/hpe-golang-workflow/project/requestmgmt"
	"google.golang.org/grpc"
)

var task_counter1 = 0
var task_counter2 = 0
var task_counter3 = 0
var c = 0
var a = 0
var b = 0

const (
	port = ":50051"
)

type RequestManagementServer struct {
	pb.UnimplementedRequestManagementServer
}

type RequestBody struct {
	WorkID     string `json:"work_id"`
	PriorityID int    `json:"p_id"`
}

var dequeuedTasks []RequestBody
var mu sync.Mutex

func sendRequest(requestBody RequestBody, url string) {
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

}

var Q1 = Queue.Queue{

	Size: 100,
}

var Q2 = Queue.Queue{

	Size: 100,
}
var Q3 = Queue.Queue{

	Size: 100,
}

func (s *RequestManagementServer) CreateRequest(ctx context.Context, in *pb.NewRequest) (*pb.Request, error) {

	//log.Printf("received: %v", in.GetWid())
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
		time.Sleep(time.Microsecond)
	}

	customer1 := Queue.New(workflow_id, rid, ctx, id, time.Now())
	ans := fmt.Sprintf("Enqueud at %s id %d", time.Now(), id)

	// fmt.Printf("WID: %s  Priority: %v\n", workflow_id, id)
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

	go Q1.MoveToFrontIfOverdue(4)
	go Q2.MoveToFrontIfOverdue(5)
	go Q3.MoveToFrontIfOverdue(6)

	log.Println("All workers ready ")
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterRequestManagementServer(s, &RequestManagementServer{})
	log.Printf("server listening at %v", lis.Addr())

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	http.HandleFunc("/dequeuedTasks", dequeuedTasksHandler)
	http.ListenAndServe(":9092", nil)

	fmt.Println("vdfvkjf")
	select {}
}
func dequeuedTasksHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Dequeued Tasks</title>
		<style>
			body {
				font-family: Arial, sans-serif;
				margin: 0;
				padding: 0;
				display: flex;
				justify-content: center;
				align-items: center;
				height: 100vh;
				background-color: #f0f0f0;
			}
			table {
				border-collapse: collapse;
				width: 80%;
				margin: 20px auto;
				background-color: #fff;
				box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
			}
th, td {
				padding: 12px;
				text-align: left;
				border-bottom: 1px solid #ddd;
			}
			th {
				background-color: #4CAF50;
				color: white;
			}
			tr:nth-child(even) {
				background-color: #f2f2f2;
			}
			tr:hover {
				background-color: #ddd;
			}
		</style>
	</head>
	<body>
		<table>
			<tr>
				<th>Work ID</th>
				<th>Priority</th>
			</tr>`)
	for _, task := range dequeuedTasks {
		fmt.Fprintf(w, `
					<tr>
						<td>%s</td>
						<td>%d</td>
					</tr>`, task.WorkID, task.PriorityID)
	}
	fmt.Fprintln(w, `
				</table>
			</body>
			</html>`)
}

func worker1() {

	for c > -10 {
		if task_counter1 < 3 {
			task_counter1++
			go NetworkingServiceProcessing()

			time.Sleep(time.Millisecond * 100)
		} else {
			time.Sleep(time.Second)
		}

	}
}
func worker2() {

	for a > -10 {
		if task_counter2 < 3 {

			task_counter2++
			go PrivateCloudEnterpriseServiceProcessing()
			time.Sleep(time.Millisecond * 100)
		} else {
			time.Sleep(time.Second)
		}

	}

}

func worker3() {

	for b > -10 {
		if task_counter3 < 3 {
			task_counter3++
			go BlockStorageServiceProcessing()
			time.Sleep(time.Millisecond * 100)
		} else {
			time.Sleep(time.Second)
		}

	}

}

func PrivateCloudEnterpriseServiceProcessing() {

	response1, _, _, priority, _, err1 := Q2.Dequeue()

	if err1 == nil {
		request1 := RequestBody{
			WorkID:     response1,
			PriorityID: int(priority),
		}

		mu.Lock()
		dequeuedTasks = append(dequeuedTasks, request1)
		mu.Unlock()

		fmt.Printf("WID: %s  Priority: %v\n", response1, priority)
		// time waiting for completion of service
		time.Sleep(time.Second * 10)

		sendRequest(request1, "http://localhost:8090/endpoint2")

		time.Sleep(time.Millisecond)

	}
	task_counter2--
}

func NetworkingServiceProcessing() {

	response2, _, _, priority, _, err2 := Q1.Dequeue()

	if err2 == nil {
		request1 := RequestBody{
			WorkID:     response2,
			PriorityID: int(priority),
		}
		mu.Lock()
		dequeuedTasks = append(dequeuedTasks, request1)
		mu.Unlock()

		fmt.Printf("WID: %s  Priority: %v\n", response2, priority)
		time.Sleep(time.Second * 10)

		sendRequest(request1, "http://localhost:8090/endpoint1")
		time.Sleep(time.Millisecond)

	}
	task_counter1--

}

func BlockStorageServiceProcessing() {

	response3, _, _, priority, _, err3 := Q3.Dequeue()

	if err3 == nil {

		request1 := RequestBody{
			WorkID:     response3,
			PriorityID: int(priority),
		}
		mu.Lock()
		dequeuedTasks = append(dequeuedTasks, request1)
		mu.Unlock()

		fmt.Printf("WID: %s  Priority: %v\n", response3, priority)
		time.Sleep(time.Second * 10)

		sendRequest(request1, "http://localhost:8090/endpoint3")

		time.Sleep(time.Millisecond)

	}
	task_counter3--
}

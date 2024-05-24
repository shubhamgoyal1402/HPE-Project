package workflows

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	pb "github.com/shubhamgoyal1402/hpe-golang-workflow/project/requestmgmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"go.uber.org/cadence/activity"

	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

func init() {
	// Registering workflow and activtiy

	workflow.Register(CustomerWorkflow2)

	activity.Register(validateUser)

	activity.Register(Activity3)
	activity.Register(subscriptionDetails)
	activity.Register(blockStorage)

}

func CustomerWorkflow2(ctx workflow.Context, id int) error {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute * 60,
		StartToCloseTimeout:    time.Minute * 60,
		HeartbeatTimeout:       time.Minute * 60,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// To get the Workflow ID and run ID
	wid := workflow.GetInfo(ctx).WorkflowExecution.ID
	rid := workflow.GetInfo(ctx).WorkflowExecution.RunID

	logger := workflow.GetLogger(ctx)
	logger.Info("Customer workflow started")
	var Result string

	err2 := workflow.ExecuteActivity(ctx, validateUser, wid).Get(ctx, &Result)
	if err2 != nil {
		logger.Error("Validate User Failed", zap.Error(err2))
		return err2
	}
	err3 := workflow.ExecuteActivity(ctx, subscriptionDetails, wid).Get(ctx, &Result)
	if err3 != nil {
		logger.Error("Subscription Failed", zap.Error(err3))
		return err3
	}

	err := workflow.ExecuteActivity(ctx, Activity3, wid, rid, id).Get(ctx, &Result)
	if err != nil {
		logger.Error("Activity failed.", zap.Error(err))
		return err
	}

	var signalName = wid
	var signalVal string
	signalChan := workflow.GetSignalChannel(ctx, signalName)

	s := workflow.NewSelector(ctx)
	s.AddReceive(signalChan, func(c workflow.Channel, more bool) {
		c.Receive(ctx, &signalVal)
		workflow.GetLogger(ctx).Info("Received signal!", zap.String("signal", signalName), zap.String("value", signalVal))
	})
	s.Select(ctx)

	if len(signalVal) > 0 && signalVal != wid {
		return errors.New(wid)
	}

	err4 := workflow.ExecuteActivity(ctx, blockStorage, wid).Get(ctx, &Result)
	if err4 != nil {
		logger.Error("Activity failed.", zap.Error(err4))
		return err4
	}

	logger.Info("Workflow completed.", zap.String("Result", Result))

	return nil
}

func validateUser(ctx context.Context, workflow_id string) error {

	time.Sleep(time.Second * 5)
	endpoint := "http://localhost:9090/validate"
	resp, err := http.Post(endpoint, "application/x-www-form-urlencoded", bytes.NewBufferString(""))
	if err != nil {
		fmt.Printf("Error posting to %s: %v\n", endpoint, err)
		return err

	}

	defer resp.Body.Close()

	return nil

}

func blockStorage(ctx context.Context, workflow_id string) error {

	time.Sleep(time.Second * 5)
	endpoint := "http://localhost:9090/blockstorage"
	resp, err := http.Post(endpoint, "application/x-www-form-urlencoded", bytes.NewBufferString(""))
	if err != nil {
		fmt.Printf("Error posting to %s: %v\n", endpoint, err)
		return err

	}

	defer resp.Body.Close()

	return nil

}

func subscriptionDetails(ctx context.Context, workflow_id string) error {

	time.Sleep(time.Second * 5)
	endpoint := "http://localhost:9090/subscription"
	resp, err := http.Post(endpoint, "application/x-www-form-urlencoded", bytes.NewBufferString(""))
	if err != nil {
		fmt.Printf("Error posting to %s: %v\n", endpoint, err)
		return err

	}

	defer resp.Body.Close()

	return nil

}
func Activity3(ctx context.Context, workflow_id string, rid string, id int32) (string, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect %v", err)

	}

	defer conn.Close()
	c := pb.NewRequestManagementClient(conn)
	req := &pb.NewRequest{
		Wid: workflow_id, // workflow ID
		Rid: rid,         // run ID
		Id:  int32(id),   // request ID
	}

	ctx2, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel() // Ensure the context is canceled when done

	resp, err2 := c.CreateRequest(ctx2, req)
	if err2 != nil {
		log.Fatalf("failed to create request: %v", err2)
	}

	ans := fmt.Sprintf("Request created with Wid: %v", resp.GetWid())

	return ans, nil
}

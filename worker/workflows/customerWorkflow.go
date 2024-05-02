package workflows

import (
	"context"
	"errors"
	"fmt"
	"log"
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

	workflow.Register(CustomerWorkflow)
	activity.Register(Activity1)

}

const (
	address = "localhost:50051"
)

func CustomerWorkflow(ctx workflow.Context, id int) error {
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

	err := workflow.ExecuteActivity(ctx, Activity1, wid, rid, id).Get(ctx, &Result)
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

	logger.Info("Workflow completed.", zap.String("Result", Result))

	return nil
}

func Activity1(ctx context.Context, workflow_id string, rid string, id int32) (string, error) {
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

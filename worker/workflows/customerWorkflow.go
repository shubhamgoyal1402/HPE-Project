package workflows

import (
	"context"
	"errors"
	"fmt"

	"github.com/shubhamgoyal1402/hpe-golang-workflow/project/Queue"
	"go.uber.org/cadence/activity"

	"time"

	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

var Response_Queue1 = Queue.Queue2{}
var Response_Queue2 = Queue.Queue2{}
var Response_Queue3 = Queue.Queue2{}

var s = 0
var Q1 = Queue.Queue{

	Size: 10,
}

var Q2 = Queue.Queue{

	Size: 10,
}
var Q3 = Queue.Queue{

	Size: 10,
}

func init() {

	workflow.Register(customerWorkflow)
	//activity.Register(taskEnqueueActivity)
	activity.Register(Activity1)
	activity.Register(Activity3)
	activity.Register(Activity2)
	//activity.Register(activty2)
}

/**
 * This is the hello world workflow sample.
 */

// ApplicationName is the task list for this sample
const TaskListName = "Service_buy_process"

func customerWorkflow(ctx workflow.Context, id int) error {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute * 60,
		StartToCloseTimeout:    time.Minute * 60,
		HeartbeatTimeout:       time.Minute * 60,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	wid := workflow.GetInfo(ctx).WorkflowExecution.ID // to get the Workflow id
	//fmt.Println(wid)
	logger := workflow.GetLogger(ctx)
	logger.Info("helloworld workflow started")
	var Result string

	err := workflow.ExecuteActivity(ctx, Activity1, wid, id).Get(ctx, &Result)
	if err != nil {
		logger.Error("Activity failed.", zap.Error(err))
		return err
	}
	err2 := workflow.ExecuteActivity(ctx, Activity2, id).Get(ctx, &Result)
	if err2 != nil {
		logger.Error("Activity failed.", zap.Error(err2))
		return err2
	}

	err3 := workflow.ExecuteActivity(ctx, Activity3, wid, id).Get(ctx, &Result)
	if err3 != nil {
		logger.Error("Activity failed.", zap.Error(err3))
		return err3
	}

	logger.Info("Workflow completed.", zap.String("Result", Result))

	return nil
}

func Activity1(ctx context.Context, workflow_id string, id int) (string, error) {

	logger := activity.GetLogger(ctx)
	logger.Info("Activty 1 started")

	switch id {
	case 1, 4:
		ans, err := activtiy1_fn(workflow_id, id, &Q1)
		return ans, err
	case 2, 5:
		ans, err := activtiy1_fn(workflow_id, id, &Q2)
		return ans, err
	case 3, 6:
		ans, err := activtiy1_fn(workflow_id, id, &Q3)
		return ans, err
	}
	//k, err := q1.Peek()

	//if err != nil {
	//	logger.Panic(err.Error())
	//}
	return "Completed", nil
}

func activtiy1_fn(workflow_id string, id int, q *Queue.Queue) (string, error) {

	for q.GetLength() >= q.Size/2 {
		time.Sleep(time.Millisecond)
	}

	customer1 := Queue.New(workflow_id, id)
	ans := fmt.Sprintf("Enqueud at %s", time.Now())
	_, err := q.Enqueue(customer1)
	if err != nil {
		panic(err)
	}
	//fmt.Println(q)
	//q.SortStudents()

	//fmt.Println(q)

	//	ans := fmt.Sprintf("The worklfow with request %d is enqueued", id)

	return ans, err
}

func Activity2(ctx context.Context, id int) (string, error) {

	logger := activity.GetLogger(ctx)
	logger.Info("Activty 2 started")

	switch id {
	case 1, 4:
		Q1.SortStudents()
		//Q1.Display()
		return "Queue 1 sorted", nil
	case 2, 5:
		Q2.SortStudents()
		//Q2.Display()
		return "Queue 2 sorted", nil
	case 3, 6:
		Q3.SortStudents()
		//Q3.Display()
		return "Queue 3 sorted", nil
	}

	return "Activity 2 Completed", nil
}
func Activity3(ctx context.Context, wid string, id int) (string, error) {

	logger := activity.GetLogger(ctx)
	logger.Info("Activty 2 started")

	switch id {
	case 1, 4:
		ans, err := Activity3_fn(wid, &Response_Queue1)
		return ans, err
	case 2, 5:
		ans, err := Activity3_fn(wid, &Response_Queue2)
		return ans, err
	case 3, 6:
		ans, err := Activity3_fn(wid, &Response_Queue3)
		return ans, err
	}

	return "Activity 3 Completed", nil
}

func Activity3_fn(wid string, q2 *Queue.Queue2) (string, error) {
	for s > -1 {

		response := q2.SearchAndRemove(wid)

		if response == true {
			return "Task Completed", nil
		}

		time.Sleep(time.Millisecond)
	}
	return "Cant complete", errors.New("error")

}

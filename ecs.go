package ecs

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aaronland/go-aws-auth"
	"github.com/aws/aws-sdk-go-v2/aws"
	aws_ecs "github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
)

type TaskResponse struct {
	Tasks      []string
	TaskOutput *aws_ecs.RunTaskOutput
}

type TaskOptions struct {
	Task            string
	Container       string
	Cluster         string
	LaunchType      string
	PlatformVersion string
	PublicIP        string
	Subnets         []string
	SecurityGroups  []string
}

type WaitTasksOptions struct {
	Cluster  string
	TaskArns []string
	Timeout  time.Duration
	Interval time.Duration
	Logger   *log.Logger
}

func NewService(uri string) (*aws_ecs.Client, error) {
	ctx := context.Background()
	return NewClient(ctx, uri)
}

func NewClient(ctx context.Context, uri string) (*aws_ecs.Client, error) {

	cfg, err := auth.NewConfig(ctx, uri)

	if err != nil {
		return nil, err
	}

	return aws_ecs.NewFromConfig(cfg), nil
}

func LaunchTask(ctx context.Context, ecs_client *aws_ecs.Client, task_opts *TaskOptions, cmd ...string) (*TaskResponse, error) {

	cluster := aws.String(task_opts.Cluster)
	task := aws.String(task_opts.Task)

	launch_type := aws.String(task_opts.LaunchType)
	platform_version := aws.String(task_opts.PlatformVersion)

	var public_ip types.AssignPublicIp

	switch task_opts.PublicIP {
	case "ENABLED":
		public_ip = types.AssignPublicIpEnabled
	default:
		public_ip = types.AssignPublicIpDisabled
	}

	/*
		subnets := make([]*string, len(task_opts.Subnets))

		for i, sn := range task_opts.Subnets {
			subnets[i] = aws.String(sn)
		}

		security_groups := make([]*string, len(task_opts.SecurityGroups))
		for i, sg := range task_opts.SecurityGroups {
			security_groups[i] = aws.String(sg)
		}

		aws_cmd := make([]*string, len(cmd))

		for i, str := range cmd {
			aws_cmd[i] = aws.String(str)
		}
	*/

	network := &types.NetworkConfiguration{
		AwsvpcConfiguration: &types.AwsVpcConfiguration{
			AssignPublicIp: public_ip,
			SecurityGroups: task_opts.SecurityGroups,
			Subnets:        task_opts.Subnets,
		},
	}

	process_override := &types.ContainerOverride{
		Name:    aws.String(task_opts.Container),
		Command: cmd,
	}

	overrides := &types.TaskOverride{
		ContainerOverrides: []*types.ContainerOverride{
			process_override,
		},
	}

	input := &aws_ecs.RunTaskInput{
		Cluster:              cluster,
		TaskDefinition:       task,
		LaunchType:           launch_type,
		PlatformVersion:      platform_version,
		NetworkConfiguration: network,
		Overrides:            overrides,
	}

	task_output, err := ecs_client.RunTask(input)

	if err != nil {
		return nil, err
	}

	if len(task_output.Tasks) == 0 {
		return nil, fmt.Errorf("run task returned no errors... but no tasks")
	}

	task_arns := make([]string, len(task_output.Tasks))

	for i, t := range task_output.Tasks {
		task_arns[i] = *t.TaskArn
	}

	task_rsp := &TaskResponse{
		Tasks:      task_arns,
		TaskOutput: task_output,
	}

	return task_rsp, nil
}

func WaitForTasksToComplete(ctx context.Context, ecs_client *aws_ecs.Client, opts *WaitTasksOptions) error {

	ctx, cancel := context.WithTimeout(ctx, opts.Timeout)
	defer cancel()

	ticker := time.NewTicker(opts.Interval)
	defer ticker.Stop()

	remaining := len(opts.TaskArns)

	for remaining > 0 {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case now := <-ticker.C:

			list_input := &aws_ecs.ListTasksInput{
				Cluster:       aws.String(opts.Cluster),
				DesiredStatus: aws.String("STOPPED"),
			}

			list_rsp, err := ecs_client.ListTasks(list_input)

			if err != nil {
				return fmt.Errorf("Failed to list tasks, %w", err)
			}

			for _, stopped_t := range list_rsp.TaskArns {

				for _, t := range opts.TaskArns {

					if *stopped_t == t {
						remaining -= 1
						break
					}
				}
			}

			if opts.Logger != nil {
				opts.Logger.Printf("%v %d tasks remaining", now, remaining)
			}
		}
	}

	return nil
}

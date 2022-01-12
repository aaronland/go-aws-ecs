package main

import (
	"flag"
	"fmt"
	"github.com/aaronland/go-aws-ecs"
	"github.com/sfomuseum/go-flags/multi"
	"log"
)

func main() {

	task := flag.String("task", "", "The name (and version) of your ECS task.")
	container := flag.String("container", "", "The name of your ECS container.")
	cluster := flag.String("cluster", "", "The name of your ECS cluster.")
	launch_type := flag.String("launch-type", "", "A valid ECS launch type.")
	platform_version := flag.String("platform-version", "", "A valid ECS platform version.")
	public_ip := flag.String("public-ip", "", "A valid ECS public IP string.")

	var subnets multi.MultiString
	var security_groups multi.MultiString

	flag.Var(&subnets, "subnet", "One or more subnets to run your ECS task in.")
	flag.Var(&security_groups, "security-group", "A valid AWS security group to run your task under.")

	dsn := flag.String("dsn", "", "A valid aaronland/go-aws-session DSN string.")

	flag.Parse()

	opts := &ecs.TaskOptions{
		Task:            *task,
		Container:       *container,
		Cluster:         *cluster,
		LaunchType:      *launch_type,
		PlatformVersion: *platform_version,
		PublicIP:        *public_ip,
		Subnets:         subnets,
		SecurityGroups:  security_groups,
	}

	cmd := flag.Args()

	rsp, err := ecs.LaunchTaskWithDSN(*dsn, opts, cmd...)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rsp)
}

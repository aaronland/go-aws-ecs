package main

import (
	"flag"
	"fmt"
	"github.com/aaronland/go-aws-ecs"
	"log"
)

func main() {

	task := flag.String("task", "", "")
	container := flag.String("container", "", "")
	cluster := flag.String("cluster", "", "")
	launch_type := flag.String("launch-type", "", "")
	public_ip := flag.String("public-ip", "", "")

	dsn := flag.String("dsn", "", "")

	flag.Parse()

	opts := &ecs.TaskOptions{
		Task:       *task,
		Container:  *container,
		Cluster:    *cluster,
		LaunchType: *launch_type,
		PublicIP:   *public_ip,
	}

	args := flag.Args()

	rsp, err := ecs.LaunchTaskWithDSN(*dsn, opts, args...)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rsp)
}

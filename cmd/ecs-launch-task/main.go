package main

import (
	"flag"
	"fmt"
	"github.com/aaronland/go-aws-ecs"
	"github.com/sfomuseum/go-flags/multi"	
	"log"
)

func main() {

	task := flag.String("task", "", "")
	container := flag.String("container", "", "")
	cluster := flag.String("cluster", "", "")
	launch_type := flag.String("launch-type", "", "")
	public_ip := flag.String("public-ip", "", "")

	var subnets multi.MultiString
	var security_groups multi.MultiString
	
	flag.Var(&subnets, "subnet", "...")
	flag.Var(&security_groups, "security-group", "...")	
	
	dsn := flag.String("dsn", "", "")

	flag.Parse()

	opts := &ecs.TaskOptions{
		Task:       *task,
		Container:  *container,
		Cluster:    *cluster,
		LaunchType: *launch_type,
		PublicIP:   *public_ip,
		Subnets:	subnets,
		SecurityGroups: security_groups,
	}

	args := flag.Args()

	rsp, err := ecs.LaunchTaskWithDSN(*dsn, opts, args...)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rsp)
}

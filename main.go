package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/nikitha/kubeops-cli/cmd"
)

func main(){
	fmt.Println("====================================")
	fmt.Println("Welcome to KubeOps CLI ")
	fmt.Println("====================================")

	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("go run main.go <command>")
		return
	}

	command := os.Args[1]

	switch command {
		 
	case "version":
		cmd.Version()

	case "deploy":

		if len(os.Args) < 3 {
			fmt.Println("Usage:")
			fmt.Println("go run main.go deploy <yaml-file>")
			return
		}
		cmd.Deploy(os.Args[2])

	case "plan":

		if len(os.Args) < 3 {
			fmt.Println("Usage:")
			fmt.Println("go run main.go plan <yaml-file>")
			return
		}
		cmd.Plan(os.Args[2])
	
	case "validate":

		if len(os.Args) < 3 {
			fmt.Println("Usage:")
			fmt.Println("go run main.go validate <yaml-file>")
			return
		}
		cmd.Validate(os.Args[2])

	case "connect":
		cmd.Connect()

	case "deployments":
		cmd.Deployments()

	case "delete":

		if len(os.Args) < 3 {
			fmt.Println("Usage:")
			fmt.Println("go run main.go delete <deployment-name>")
			return
		}
        cmd.Delete(os.Args[2])

	case "scale":

		if len(os.Args) < 4 {
			fmt.Println("Usage:")
			fmt.Println("go run main.go scale <deployment-name> <replicas>")
			return
		}

	replicas, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Invalid replica count")
		return
	}
	cmd.Scale(os.Args[2], int32(replicas))


	case "status":

	if len(os.Args) < 3 {
		fmt.Println("Usage:")
		fmt.Println("go run main.go status <deployment-name>")
		return
	}

	cmd.Status(os.Args[2])

	default:
		fmt.Println("Unknown Command")
	}
}
package main

import (
	"flag"
	"fmt"

	"github.com/joonghyunlee/go-os-billing-batch/metric/instance"
	"github.com/joonghyunlee/go-os-billing-batch/version"
)

func main() {
	resourceType := flag.String("type", "event", "Resource Type to use. event or metric")

	versionFlag := flag.Bool("version", false, "Version")
	flag.Parse()

	if *versionFlag {
		fmt.Println("Build Date:", version.BuildDate)
		fmt.Println("Git Commit:", version.GitCommit)
		fmt.Println("Version:", version.Version)
		fmt.Println("Go Version:", version.GoVersion)
		fmt.Println("OS / Arch:", version.OsArch)
		return
	}
	fmt.Println("Hello.")
	fmt.Println(*resourceType)

	instance.Aggregate()
}

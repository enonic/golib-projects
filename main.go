package main

import (
	"flag"
	"github.com/enonic/projects/config"
	"fmt"
	"github.com/golang/glog"
)

func main() {

	args := parseargs()

	clientConfig, err := config.Create(args.config)

	if err != nil {
		glog.Fatal(err)
	}

	orgs := NewClient(clientConfig).GetOrganizations()

	for _, o := range orgs {
		fmt.Printf(" [ %s ] \n", o.DisplayName)
		for _, p := range o.Projects {
			fmt.Printf("    - %s, %s, Backup: %t\n", p.Name, p.Id, p.Backup.Enabled)
		}
	}
}

type ArgsParsed struct {
	config string
}

func parseargs() *ArgsParsed {
	args := &ArgsParsed{}
	flag.StringVar(&args.config, "config", "config.yaml", "Config file for the service (json)")
	flag.Parse()
	flag.Set("logtostderr", "true")
	return args
}

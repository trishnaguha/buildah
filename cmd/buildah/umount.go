package main

import (
	"fmt"

	"github.com/urfave/cli"
)

var (
	umountFlags = []cli.Flag{
		cli.StringFlag{
			Name:  "name",
			Usage: "name or ID of the working container",
		},
		cli.StringFlag{
			Name:  "root",
			Usage: "root directory of the working container",
		},
		cli.StringFlag{
			Name:  "link",
			Usage: "a symlink to the root directory of the working container",
		},
	}
)

func umountCmd(c *cli.Context) error {
	name := ""
	if c.IsSet("name") {
		name = c.String("name")
	}
	root := ""
	if c.IsSet("root") {
		root = c.String("root")
	}
	link := ""
	if c.IsSet("link") {
		link = c.String("link")
		if link == "" {
			return fmt.Errorf("link location can not be empty")
		}
	}
	if name == "" && root == "" && link == "" {
		return fmt.Errorf("either --name or --root or --link, or some combination, must be specified")
	}

	store, err := getStore(c)
	if err != nil {
		return err
	}

	builder, err := openBuilder(store, name, root, link)
	if err != nil {
		return fmt.Errorf("error reading build container %q: %v", name, err)
	}

	err = builder.Unmount()
	if err != nil {
		return fmt.Errorf("error unmounting container: %v", err)
	}

	return nil
}

package cmd

import (
	"fmt"

	"github.com/urfave/cli"

	"github.com/jbrekelmans/kube-compose/pkg/up"
)

func NewUpCommand() cli.Command {
	return cli.Command{
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name: "no-build",
			},
		},
		Name:  "up",
		Usage: "creates pods and services in an order that respects depends_on in the docker compose file",
		Action: func(c *cli.Context) error {
			cfg, err := newConfigFromEnv()
			if err != nil {
				return err
			}
			err = updateConfigFromCli(cfg, c)
			if err != nil {
				return err
			}
			n := c.NArg()
			serviceArgs := make(map[string]bool, n)
			for i := 0; i < n; i++ {
				service := c.Args().Get(i)
				if _, ok := cfg.CanonicalComposeFile.Services[service]; !ok {
					return fmt.Errorf("No such service: %s", service)
				}
				serviceArgs[service] = true
			}
			return up.Run(cfg, serviceArgs)
		},
	}
}

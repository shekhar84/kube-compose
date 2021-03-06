package cmd

import (
	"fmt"

	"github.com/jbrekelmans/kube-compose/pkg/config"
	"github.com/urfave/cli"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	"k8s.io/client-go/tools/clientcmd"
)

const (
	environmentIDFlagName = "env-id"
	namespaceFlagName     = "namespace"
)

func GlobalFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   environmentIDFlagName + ", e",
			EnvVar: "KUBECOMPOSE_ENVID",
			Usage:  "used to isolate environments deployed to a shared namespace, by (1) using this value as a suffix of pod and service names and (2) using this value to isolate selectors",
		},
		cli.StringFlag{
			Name:   namespaceFlagName + ", n",
			EnvVar: "KUBECOMPOSE_NAMESPACE",
			Usage:  "the target Kubernetes namespace",
		},
	}
}

func newConfigFromEnv() (*config.Config, error) {
	cfg, err := config.New()
	if err != nil {
		return nil, err
	}
	loader := clientcmd.NewDefaultClientConfigLoadingRules()
	overrides := clientcmd.ConfigOverrides{}
	clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loader, &overrides)
	kubeConfig, err := clientConfig.ClientConfig()
	if err != nil {
		return nil, err
	}
	namespace, _, err := clientConfig.Namespace()
	if err != nil {
		return nil, err
	}
	cfg.KubeConfig = kubeConfig
	cfg.Namespace = namespace
	return cfg, nil
}

func updateConfigFromCli(cfg *config.Config, c *cli.Context) error {
	environmentID := c.GlobalString(environmentIDFlagName)
	cfg.Services = c.Args()
	if len(environmentID) == 0 && !c.GlobalIsSet(environmentIDFlagName) {
		return fmt.Errorf("the environment id is required")
	} else if len(environmentID) == 0 {
		return fmt.Errorf("environment id must not be empty")
	}
	cfg.EnvironmentID = environmentID

	namespace := c.GlobalString(namespaceFlagName)
	if len(namespace) > 0 || c.GlobalIsSet(namespaceFlagName) {
		if len(namespace) == 0 {
			return fmt.Errorf("namespace must not be empty")
		}
		cfg.Namespace = namespace
	}
	return nil
}

package commands

import (
	"fmt"

	"github.com/cloudfoundry-incubator/ducati-daemon/lib/namespace"
)

//go:generate counterfeiter --fake-name Repository . repository
type repository interface {
	Get(path string) (namespace.Namespace, error)
	Create(name string) (namespace.Namespace, error)
}

type CreateNamespace struct {
	Name       string
	Repository repository
}

func (cn CreateNamespace) Execute(context Context) error {
	_, err := cn.Repository.Create(cn.Name)
	if err != nil {
		return fmt.Errorf("create namespace: %s", err)
	}

	return nil
}

func (cn CreateNamespace) String() string {
	return fmt.Sprintf("ip netns add %s", cn.Name)
}
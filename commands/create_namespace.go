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
	Result     Namespace
}

func (cn *CreateNamespace) Execute(context Context) error {
	ns, err := cn.Repository.Create(cn.Name)
	if err != nil {
		return fmt.Errorf("failed to create namespace %q: %s", cn.Name, err)
	}

	cn.Result = ns
	return nil
}

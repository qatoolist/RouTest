// interfaces package

package interfaces

import "github.com/qatoolist/RouTest/internal/loaders"

type Config interface {
	Get(keys ...string) (interface{}, error)
	Set(keys []string, value interface{})
	GetHost() (Host, error)
	CopyFromTemp(cnf *loaders.Config) Config
}

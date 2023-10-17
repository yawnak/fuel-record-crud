package erradapt

import (
	"github.com/samber/lo"
	"github.com/yawnak/fuel-record-crud/internal/common/httperr"
)

type ErrAdapterFunc func(err error) *httperr.SError

type Adapter struct {
	errs []ErrAdapterFunc
}

func (adapter *Adapter) registerAdapterFunc(fn ErrAdapterFunc) {
	adapter.errs = append(adapter.errs, fn)
}

func (adapter *Adapter) Adapt(err error) *httperr.SError {
	if err == nil {
		return nil
	}
	if serr, ok := lo.ErrorsAs[*httperr.SError](err); ok {
		return serr
	}
	for _, fn := range adapter.errs {
		if adaptedErr := fn(err); adaptedErr != nil {
			return adaptedErr
		}
	}

	return httperr.NewInternal(err)
}

func NewAdapter() *Adapter {
	adapter := &Adapter{}

	//register adapters here
	registerCarAdapters(adapter)

	return adapter
}

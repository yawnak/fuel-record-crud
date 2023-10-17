package erradapt

import (
	"fmt"
	"testing"

	"github.com/samber/lo"
	"github.com/yawnak/fuel-record-crud/internal/common/httperr"
)

func TestLoAs(t *testing.T) {
	err := &httperr.SError{
		Code: 400,
		//Err:            fmt.Errorf("test error"),
		ErrType:        "test",
		DisplayMessage: "test message",
	}
	serr, ok := lo.ErrorsAs[*httperr.SError](err)
	if ok {
		fmt.Println("OK")
	}
	fmt.Println(serr)
}

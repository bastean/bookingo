package errs

import (
	"fmt"

	"github.com/bastean/bookingo/pkg/cmd/server/service/errors"
)

func MissingKey(what, where string) error {
	return errors.NewInternal(
		where,
		fmt.Sprintf("failure to obtain the value of the key [%v]", what),
		nil,
	)
}

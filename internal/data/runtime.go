package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	valueJSON := fmt.Sprintf("%d mins", r)
	quotedValueJSON := strconv.Quote(valueJSON)
	return []byte(quotedValueJSON), nil
}

package data

import (
	"fmt"
	"strconv"
)

type Pages int32

func (p Pages) MarshalJSON() ([]byte, error){
	jsonValue := fmt.Sprintf("%d стр.", p)
	quotedJSONValue := strconv.Quote(jsonValue)
	
	return []byte(quotedJSONValue), nil
}
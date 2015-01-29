package novels

import (
	"encoding/json"
)

type Novel struct {
	Id int64
	Name string
}

func New() *Novel {
	return &Novel{1, "foobar"}
}

func NovelToJson() string {
	novel := Novel{1, "foobar"}
	result, _ := json.Marshal(novel)
	return string(result)
}

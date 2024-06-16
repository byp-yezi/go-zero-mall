package codex

import (
	"go-zero-mall/common/codex/types"
	"strconv"
)

var _ CodeX = (*Status)(nil)

type Status struct {
	sts *types.Status
}

func (s *Status) Error() string {
	return s.Message()
}

func (s *Status) Code() int {
	return int(s.sts.Code)
}

func (s *Status) Message() string {
	if s.sts.Message == "" {
		return strconv.Itoa(int(s.sts.Code))
	}
	return s.sts.Message
}

func (s *Status) Details() []interface{} {
	return nil
}

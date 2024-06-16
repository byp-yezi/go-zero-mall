package codex

import "strconv"

type CodeX interface {
	Error() string
	Code() int
	Message() string
	Details() []interface{}
}

type Code struct {
	code int
	msg  string
}

func (c Code) Error() string {
	if len(c.msg) > 0 {
		return c.msg
	}
	return strconv.Itoa(c.code)
}

func (c Code) Code() int {
	return c.code
}

func (c Code) Message() string {
	return c.msg
}

func (c Code) Details() []interface{} {
	return nil
}

func New(code int, msg string) Code {
	return Code{code: code, msg: msg}
}

func add(code int, msg string) Code {
	return Code{code: code, msg: msg}
}

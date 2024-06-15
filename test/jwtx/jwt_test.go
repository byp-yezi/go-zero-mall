package jwtx

import (
	"go-zero-mall/common/jwtx"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetToken(t *testing.T) {
	output, _ := jwtx.GetToken("123", 1, 2, 3)
	assert.Equal(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjMsImlhdCI6MSwidWlkIjozfQ._xh6-tOhVYCr8ePMeYYeS7ak5WjsE0tfaI-qN0TLeV4", output)
}

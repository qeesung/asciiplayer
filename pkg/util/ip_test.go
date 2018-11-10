package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetIPList(t *testing.T) {
	assertions := assert.New(t)
	ipList, err := GetIPList()
	assertions.Nil(err)
	assertions.Contains(ipList, "127.0.0.1")
}

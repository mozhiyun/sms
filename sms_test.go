package sms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClientTencent_SendSmsCode(t *testing.T) {
	client := NewClientTencent("", "", "", "", "")
	_, err := client.SendSmsCode("")
	assert.Nil(t, err)
}

func TestVerifyMobileFormat(t *testing.T) {
	assert.Equal(t, false, VerifyMobileFormat(""))
	assert.Equal(t, true, VerifyMobileFormat(""))
}

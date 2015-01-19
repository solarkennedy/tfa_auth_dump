package main

import (
	"testing"
)

func TestMain(t *testing.T) {
}

// Test the spec:
// https://code.google.com/p/google-authenticator/wiki/KeyUriFormat
func Test_generate_otp_uri(t *testing.T) {
	actual := generate_otp_uri("test_email", "test_string")
	expected := "otpauth://totp/test_email?secret=test_string"
	if expected != actual {
		t.Error("Expected, actual: ", expected, actual)
	}
}

func Benchmark_generate_otp_uri(b *testing.B) {
	generate_otp_uri("test_email", "test_string")
}

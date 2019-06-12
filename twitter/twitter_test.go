package twitter

import (
	"testing"
)

func TestSendTweet(t *testing.T) {
	err := sendTweet("go test")
	if err != nil {
		t.Fatalf("Test failed. Error: %v", err)
	}
}

package main

import (
	"testing"
	"time"
)

func TestHelloWorldEndpoint(t *testing.T) {

}

func TestLateEndpoint(t *testing.T) {

}

func TestPasswordEndpoint(t *testing.T) {

}

func TestMain(t *testing.T) {
	// Run your main function in a separate goroutine
	go main()

	// Wait for some time to ensure the server is started
	time.Sleep(1 * time.Second)

	// Now run your tests
	t.Run("TestHelloWorldEndpoint", TestHelloWorldEndpoint)
	t.Run("TestLateEndpoint", TestLateEndpoint)
	t.Run("TestPasswordEndpoint", TestPasswordEndpoint)
}

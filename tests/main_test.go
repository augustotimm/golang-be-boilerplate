package tests

import (
	mockDB "backend-boilerplate/tests/mock-db"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	mockDB.Setup()
	code := m.Run()
	mockDB.Teardown()
	os.Exit(code)
}

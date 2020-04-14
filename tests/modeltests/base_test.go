package modeltests

import (
	"log"
	"os"
	"testing"

	"github.com/Victoria-engine/api-v2/pkg/utl/testutils"
	"github.com/joho/godotenv"
)

// TestMain : Entry point for the package test
func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}

	testutils.Database()

	os.Exit(m.Run())
}

package controllertests

import (
	"log"
	"os"
	"testing"

	"github.com/Victoria-engine/api-v2/app/utils/testutils"
	"github.com/joho/godotenv"
)

// TestMain : TestMain
func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}

	testutils.Database()

	os.Exit(m.Run())
}

//  +build integration

package integrations

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/wbrush/mmhmm/setup"
)

func TestMain(m *testing.M) {
	fmt.Printf("\n\nRunning TestMain(): Setting up test environment\n")

	//  set up integration test specific environment vars here
	os.Setenv("DB_DATABASE", "template_service")
	os.Setenv("DB_USER", "app")
	os.Setenv("DB_PASSWORD", "app")
	os.Setenv("DB_PORT", "5521")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_CONNS", "4")
	os.Setenv("DB_MIGRATION_PATH", "../dao/postgres")

	//  set up regular environment vars here
	filename := "../files/.env"
	if _, err := os.Stat(filename); err == nil {
		_ = godotenv.Load(filename)
	}

	//  initialize service
	setup.SetupAndRun(false, "commit", "builtAt", "../docs/")

	//  perform various initializations
	InitializeDBs()

	fmt.Printf("\n\nPerforming Integration Tests\n")
	os.Exit(m.Run())
}

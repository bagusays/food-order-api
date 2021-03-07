package cmd

import (
	"context"
	"errors"
	"fmt"
	"food-order-api/internal/shared/config"
	"food-order-api/internal/shared/infrastructure"
	"strconv"

	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "run migration files",
	RunE:  runMigration,
}

func init() {
	migrateCmd.PersistentFlags().Int("step", 0, "maximum migration steps")
	migrateCmd.PersistentFlags().String("direction", "up", "migration direction")
}

func runMigration(cmd *cobra.Command, args []string) error {
	cfg := config.New("./config")

	direction := cmd.Flag("direction").Value.String()
	step, err := strconv.Atoi(cmd.Flag("step").Value.String())
	if err != nil {
		return errors.New("run migrate missing step: " + err.Error())
	}

	err = migrationHandler(cfg.MySQL, "./sql", direction, step)
	if err != nil {
		return err
	}

	return nil
}

func migrationHandler(cfg config.MySQL, dir string, direction string, step int) error {
	migration := &migrate.FileMigrationSource{
		Dir: dir,
	}
	migrate.SetTable("schema_migration")

	sqlxDB, err := infrastructure.NewMySQL(context.Background(), cfg)
	if err != nil {
		return err
	}

	var n int
	if direction == "down" {
		n, err = migrate.ExecMax(sqlxDB.DB, "mysql", migration, migrate.Down, step)
	} else {
		n, err = migrate.ExecMax(sqlxDB.DB, "mysql", migration, migrate.Up, step)
	}
	if err != nil {
		return err
	}

	defer sqlxDB.Close()

	fmt.Printf("Applied %d migrations\n", n)
	return nil
}

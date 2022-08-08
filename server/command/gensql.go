package command

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/name/pjName/server/pkg/configutil"
)

func RunGensql() error {
	ctx := context.Background()
	db := configutil.GetDB()

	now := time.Now()
	version := now.Format("20060102150405")
	name := fmt.Sprintf("migrations/%s", version)

	var buf bytes.Buffer

	if err := db.Schema.WriteTo(ctx, &buf); err != nil {
		return fmt.Errorf("Error printing schema changes: %w\n", err)
	}

	s := buf.String()

	if s == "BEGIN;\nCOMMIT;\n" {
		fmt.Println("Nothing to migrate.")
		return nil
	}

	f, err := os.Create(name + ".sql")
	if err != nil {
		return fmt.Errorf("Error creating migration file: %w", err)
	}
	defer f.Close()

	_, _ = f.WriteString("-- +migrate Up\n")
	_, _ = f.WriteString(s)
	_, _ = f.WriteString("\n-- +migrate Down\n")
	fmt.Printf("create migration file: %s\n", name)
	return nil
}

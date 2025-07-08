package db

import (
	"context"
	"os"

	"api-ai/ent"
	"api-ai/ent/migrate"
	"api-ai/internal/logger"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func NewClient() *ent.Client {
	dsn := os.Getenv("DATABASE_URL")

	client, err := ent.Open(dialect.Postgres, dsn)
	if err != nil {
		logger.Default.Fatalf("failed opening connection to postgres: %v", err)
	}

	// TODO: check docs about migrations.
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),  // drop unused indexes
		migrate.WithDropColumn(true), // drop unused columns
	); err != nil {
		logger.Default.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}

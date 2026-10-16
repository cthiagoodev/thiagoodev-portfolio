package database_test

import (
	"context"
	"testing"

	"github.com/cthiagoodev/portfolio/services/app-api/internal/infrastructure/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewPool_ValidationError_EmptyDatabaseURL(t *testing.T) {
	ctx := context.Background()

	pool, err := database.NewPool(ctx, "   ")

	require.Error(t, err)
	assert.Nil(t, pool)
	assert.Contains(t, err.Error(), "database URL is required")
}

func TestNewPool_Failure_InvalidDatabaseURL(t *testing.T) {
	ctx := context.Background()

	pool, err := database.NewPool(ctx, "invalid://bad url with spaces\n")

	require.Error(t, err)
	assert.Nil(t, pool)
	assert.Contains(t, err.Error(), "failed to create database connection pool")
}

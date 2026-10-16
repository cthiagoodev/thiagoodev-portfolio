//go:build integration

package database_test

import (
	"testing"

	"github.com/cthiagoodev/portfolio/services/app-api/internal/infrastructure/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewPool_Integration_Success(t *testing.T) {
	ctx, pool := database.SetupTestDatabase(t)

	var result int
	queryError := pool.QueryRow(ctx, "SELECT 1").Scan(&result)

	require.NoError(t, queryError)
	assert.Equal(t, 1, result)
}

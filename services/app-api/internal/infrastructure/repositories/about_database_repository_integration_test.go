//go:build integration

package repositories_test

import (
	"context"
	"errors"
	"testing"

	domainerrors "github.com/cthiagoodev/portfolio/services/app-api/internal/domain/errors"
	"github.com/cthiagoodev/portfolio/services/app-api/internal/infrastructure/database"
	infrastructurerepositories "github.com/cthiagoodev/portfolio/services/app-api/internal/infrastructure/repositories"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestDatabase(t *testing.T) (context.Context, *pgxpool.Pool, *infrastructurerepositories.AboutDatabaseRepository) {
	ctx, pool := database.SetupTestDatabase(t)

	t.Cleanup(func() {
		_, _ = pool.Exec(context.Background(), "TRUNCATE TABLE about CASCADE")
	})

	_, _ = pool.Exec(ctx, "TRUNCATE TABLE about CASCADE")

	aboutDatabaseRepository := infrastructurerepositories.NewAboutDatabaseRepository(pool)
	return ctx, pool, aboutDatabaseRepository
}

func TestAboutDatabaseRepository_Get_Success(t *testing.T) {
	ctx, pool, aboutDatabaseRepository := setupTestDatabase(t)
	expectedName := "Thiago Sousa"
	expectedPhone := "+5511999999999"
	expectedDescription := "Software Engineer"
	expectedText := "Bio completa sobre engenharia de software"

	insertQuery := `
		INSERT INTO about (name, phone, description, text)
		VALUES ($1, $2, $3, $4)
	`
	_, insertError := pool.Exec(ctx, insertQuery, expectedName, expectedPhone, expectedDescription, expectedText)
	require.NoError(t, insertError)

	result, err := aboutDatabaseRepository.Get(ctx)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.NotEmpty(t, result.UUID)
	assert.Equal(t, expectedName, result.Name)
	assert.Equal(t, expectedPhone, result.Phone)
	assert.Equal(t, expectedDescription, result.Description)
	assert.Equal(t, expectedText, result.Text)
	assert.False(t, result.CreatedAt.IsZero())
	assert.False(t, result.UpdatedAt.IsZero())
}

func TestAboutDatabaseRepository_Get_Success_WithNullFields(t *testing.T) {
	ctx, pool, aboutDatabaseRepository := setupTestDatabase(t)
	expectedName := "Thiago Minimal"

	insertQuery := `
		INSERT INTO about (name, phone, description, text)
		VALUES ($1, NULL, NULL, NULL)
	`
	_, insertError := pool.Exec(ctx, insertQuery, expectedName)
	require.NoError(t, insertError)

	result, err := aboutDatabaseRepository.Get(ctx)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, expectedName, result.Name)
	assert.Empty(t, result.Phone)
	assert.Empty(t, result.Description)
	assert.Empty(t, result.Text)
}

func TestAboutDatabaseRepository_Get_NotFound(t *testing.T) {
	ctx, _, aboutDatabaseRepository := setupTestDatabase(t)

	result, err := aboutDatabaseRepository.Get(ctx)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.True(t, errors.Is(err, domainerrors.ErrNotFound))
}

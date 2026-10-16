//go:build integration

package repositories_test

import (
	"context"
	"testing"

	"github.com/cthiagoodev/portfolio/services/app-api/internal/infrastructure/database"
	infrastructurerepositories "github.com/cthiagoodev/portfolio/services/app-api/internal/infrastructure/repositories"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupSocialContactTestDatabase(t *testing.T) (context.Context, *pgxpool.Pool, *infrastructurerepositories.SocialContactDatabaseRepository) {
	ctx, pool := database.SetupTestDatabase(t)

	t.Cleanup(func() {
		_, _ = pool.Exec(context.Background(), "TRUNCATE TABLE social_contacts CASCADE")
	})

	_, _ = pool.Exec(ctx, "TRUNCATE TABLE social_contacts CASCADE")

	socialContactDatabaseRepository := infrastructurerepositories.NewSocialContactDatabaseRepository(pool)
	return ctx, pool, socialContactDatabaseRepository
}

func TestSocialContactDatabaseRepository_GetAll_Success(t *testing.T) {
	ctx, pool, socialContactDatabaseRepository := setupSocialContactTestDatabase(t)
	firstLabel := "github"
	firstURL := "https://github.com/cthiagoodev"
	firstImagePath := "/images/icons/github.svg"
	secondLabel := "linkedin"
	secondURL := "https://linkedin.com/in/cthiagoodev"
	secondImagePath := "/images/icons/linkedin.svg"

	insertQuery := `
		INSERT INTO social_contacts (label, url, image_path)
		VALUES ($1, $2, $3), ($4, $5, $6)
	`
	_, insertError := pool.Exec(
		ctx,
		insertQuery,
		firstLabel,
		firstURL,
		firstImagePath,
		secondLabel,
		secondURL,
		secondImagePath,
	)
	require.NoError(t, insertError)

	result, err := socialContactDatabaseRepository.GetAll(ctx)

	require.NoError(t, err)
	require.Len(t, result, 2)
	assert.NotEmpty(t, result[0].UUID)
	assert.Equal(t, firstLabel, result[0].Label)
	assert.Equal(t, firstURL, result[0].URL)
	assert.Equal(t, firstImagePath, result[0].ImagePath)
	assert.False(t, result[0].CreatedAt.IsZero())
	assert.False(t, result[0].UpdatedAt.IsZero())
	assert.NotEmpty(t, result[1].UUID)
	assert.Equal(t, secondLabel, result[1].Label)
	assert.Equal(t, secondURL, result[1].URL)
	assert.Equal(t, secondImagePath, result[1].ImagePath)
	assert.False(t, result[1].CreatedAt.IsZero())
	assert.False(t, result[1].UpdatedAt.IsZero())
}

func TestSocialContactDatabaseRepository_GetAll_Success_WithNullImagePath(t *testing.T) {
	ctx, pool, socialContactDatabaseRepository := setupSocialContactTestDatabase(t)
	expectedLabel := "github"
	expectedURL := "https://github.com/cthiagoodev"

	insertQuery := `
		INSERT INTO social_contacts (label, url, image_path)
		VALUES ($1, $2, NULL)
	`
	_, insertError := pool.Exec(ctx, insertQuery, expectedLabel, expectedURL)
	require.NoError(t, insertError)

	result, err := socialContactDatabaseRepository.GetAll(ctx)

	require.NoError(t, err)
	require.Len(t, result, 1)
	assert.Equal(t, expectedLabel, result[0].Label)
	assert.Equal(t, expectedURL, result[0].URL)
	assert.Empty(t, result[0].ImagePath)
}

func TestSocialContactDatabaseRepository_GetAll_EmptyList(t *testing.T) {
	ctx, _, socialContactDatabaseRepository := setupSocialContactTestDatabase(t)

	result, err := socialContactDatabaseRepository.GetAll(ctx)

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Empty(t, result)
}


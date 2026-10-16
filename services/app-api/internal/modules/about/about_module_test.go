package about_test

import (
	"testing"

	"github.com/cthiagoodev/portfolio/services/app-api/internal/modules/about"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewAboutModule_Success(t *testing.T) {
	var pool *pgxpool.Pool

	aboutModule := about.NewAboutModule(pool)

	require.NotNil(t, aboutModule)
	assert.NotNil(t, aboutModule.Handler)
}

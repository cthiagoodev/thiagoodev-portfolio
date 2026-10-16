package social_contact_test

import (
	"testing"

	"github.com/cthiagoodev/portfolio/services/app-api/internal/modules/social_contact"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewSocialContactModule_Success(t *testing.T) {
	var pool *pgxpool.Pool

	socialContactModule := social_contact.NewSocialContactModule(pool)

	require.NotNil(t, socialContactModule)
	assert.NotNil(t, socialContactModule.Handler)
}


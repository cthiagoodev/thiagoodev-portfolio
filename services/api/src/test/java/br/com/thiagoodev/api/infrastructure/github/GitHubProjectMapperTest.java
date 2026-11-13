package br.com.thiagoodev.api.infrastructure.github;

import br.com.thiagoodev.api.domain.entities.Project;
import org.junit.jupiter.api.Test;

import java.time.OffsetDateTime;
import java.time.ZoneOffset;
import java.util.List;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertNull;
import static org.junit.jupiter.api.Assertions.assertTrue;

class GitHubProjectMapperTest {

    @Test
    void shouldMapGitHubRepositoryToProject_whenAllFieldsArePresent() {
        OffsetDateTime createdAt = OffsetDateTime.of(2024, 1, 10, 12, 0, 0, 0, ZoneOffset.UTC);
        OffsetDateTime updatedAt = OffsetDateTime.of(2024, 2, 10, 14, 0, 0, 0, ZoneOffset.UTC);
        OffsetDateTime pushedAt = OffsetDateTime.of(2024, 2, 15, 16, 0, 0, 0, ZoneOffset.UTC);
        GitHubRepositoryResponse response = new GitHubRepositoryResponse(
                42L,
                "portfolio",
                "Personal portfolio",
                "Java",
                List.of("portfolio", "spring-boot"),
                false,
                "https://thiagoodev.com",
                "https://github.com/cthiagoodev/portfolio",
                createdAt,
                updatedAt,
                pushedAt
        );

        Project project = GitHubProjectMapper.toDomain(response);

        assertNull(project.getUuid());
        assertEquals(42L, project.getExternalId());
        assertEquals("portfolio", project.getName());
        assertEquals("Personal portfolio", project.getDescription());
        assertEquals("https://thiagoodev.com", project.getUrl());
        assertEquals("https://github.com/cthiagoodev/portfolio", project.getRepoUrl());
        assertEquals(List.of("Java"), project.getLanguages());
        assertEquals(createdAt.toLocalDateTime(), project.getCreatedAt());
        assertEquals(updatedAt.toLocalDateTime(), project.getUpdatedAt());
        assertEquals(pushedAt.toLocalDateTime(), project.getPushedAt());
    }

    @Test
    void shouldMapOptionalFieldsSafely_whenNullableFieldsAreNull() {
        GitHubRepositoryResponse response = new GitHubRepositoryResponse(
                99L,
                "minimal-repo",
                null,
                null,
                List.of("portfolio"),
                false,
                null,
                "https://github.com/cthiagoodev/minimal-repo",
                null,
                null,
                null
        );

        Project project = GitHubProjectMapper.toDomain(response);

        assertEquals("", project.getDescription());
        assertEquals("", project.getUrl());
        assertTrue(project.getLanguages().isEmpty());
        assertNotNull(project.getCreatedAt());
        assertNotNull(project.getUpdatedAt());
        assertNull(project.getPushedAt());
    }
}

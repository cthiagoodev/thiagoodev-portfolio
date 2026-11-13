package br.com.thiagoodev.api.infrastructure.github;

import br.com.thiagoodev.api.domain.entities.Project;

import java.time.LocalDateTime;
import java.time.OffsetDateTime;
import java.util.List;

public final class GitHubProjectMapper {

    private GitHubProjectMapper() {
    }

    public static Project toDomain(GitHubRepositoryResponse response) {
        List<String> languages = response.language() != null && !response.language().isBlank()
                ? List.of(response.language())
                : List.of();

        return new Project(
                null,
                response.id(),
                response.name(),
                response.description(),
                response.homepage(),
                response.htmlUrl(),
                languages,
                List.of(),
                toLocalDateTime(response.pushedAt()),
                toLocalDateTime(response.createdAt()),
                toLocalDateTime(response.updatedAt())
        );
    }

    public static List<Project> toDomainList(List<GitHubRepositoryResponse> responses) {
        if (responses == null) {
            return List.of();
        }

        return responses.stream()
                .map(GitHubProjectMapper::toDomain)
                .toList();
    }

    private static LocalDateTime toLocalDateTime(OffsetDateTime dateTime) {
        return dateTime != null ? dateTime.toLocalDateTime() : null;
    }
}

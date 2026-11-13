package br.com.thiagoodev.api.infrastructure.github;

import com.fasterxml.jackson.annotation.JsonProperty;

import java.time.OffsetDateTime;
import java.util.List;

public record GitHubRepositoryResponse(
        Long id,
        String name,
        String description,
        String language,
        List<String> topics,
        boolean fork,
        String homepage,
        @JsonProperty("html_url") String htmlUrl,
        @JsonProperty("created_at") OffsetDateTime createdAt,
        @JsonProperty("updated_at") OffsetDateTime updatedAt,
        @JsonProperty("pushed_at") OffsetDateTime pushedAt
) {
}

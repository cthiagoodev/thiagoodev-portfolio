package br.com.thiagoodev.api.infrastructure.github;

import br.com.thiagoodev.api.domain.entities.Project;
import org.junit.jupiter.api.Test;
import org.springframework.http.MediaType;
import org.springframework.test.web.client.MockRestServiceServer;
import org.springframework.web.client.RestClient;

import java.util.List;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.springframework.test.web.client.match.MockRestRequestMatchers.requestTo;
import static org.springframework.test.web.client.response.MockRestResponseCreators.withSuccess;

class GitHubProjectClientTest {

    @Test
    void shouldReturnOnlyPortfolioRepositories_whenGitHubReturnsMixedTopics() {
        RestClient.Builder builder = RestClient.builder();
        MockRestServiceServer server = MockRestServiceServer.bindTo(builder).build();
        GitHubProjectClient client = new GitHubProjectClient(builder.build(), "https://api.github.com/users/cthiagoodev/repos");
        server.expect(requestTo("https://api.github.com/users/cthiagoodev/repos?per_page=100&page=1"))
                .andRespond(withSuccess("""
                        [
                          {
                            "id": 1,
                            "name": "portfolio-web",
                            "description": "Portfolio frontend",
                            "language": "TypeScript",
                            "topics": ["portfolio", "astro"],
                            "fork": false,
                            "homepage": "https://thiagoodev.com",
                            "html_url": "https://github.com/cthiagoodev/portfolio-web",
                            "created_at": "2024-01-01T10:00:00Z",
                            "updated_at": "2024-02-01T10:00:00Z",
                            "pushed_at": "2024-02-10T10:00:00Z"
                          },
                          {
                            "id": 2,
                            "name": "sandbox-repo",
                            "description": "Not for portfolio",
                            "language": "Go",
                            "topics": ["study"],
                            "fork": false,
                            "homepage": null,
                            "html_url": "https://github.com/cthiagoodev/sandbox-repo",
                            "created_at": "2024-01-01T10:00:00Z",
                            "updated_at": "2024-02-01T10:00:00Z",
                            "pushed_at": "2024-02-10T10:00:00Z"
                          }
                        ]
                        """, MediaType.APPLICATION_JSON));

        List<Project> result = client.fetchPortfolioRepositories();

        assertEquals(1, result.size());
        assertEquals(1L, result.getFirst().getExternalId());
        assertEquals("portfolio-web", result.getFirst().getName());
        assertEquals(List.of("TypeScript"), result.getFirst().getLanguages());
        server.verify();
    }
}

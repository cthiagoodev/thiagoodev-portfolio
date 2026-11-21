package br.com.thiagoodev.api.infrastructure.github;

import br.com.thiagoodev.api.domain.entities.Project;
import br.com.thiagoodev.api.domain.gateways.ProjectsGateway;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.core.ParameterizedTypeReference;
import org.springframework.stereotype.Component;
import org.springframework.web.client.RestClient;
import org.springframework.web.util.UriComponentsBuilder;

import java.net.URI;
import java.util.ArrayList;
import java.util.List;

@Component
public class GitHubProjectClient implements ProjectsGateway {

    private static final int PER_PAGE = 100;
    private static final String PORTFOLIO_TOPIC = "portfolio";

    private final RestClient restClient;
    private final String repositoriesUrl;

    public GitHubProjectClient(
            RestClient restClient,
            @Value("${github.api.repos-url}") String repositoriesUrl
    ) {
        this.restClient = restClient;
        this.repositoriesUrl = repositoriesUrl;
    }

    @Override
    public List<Project> fetchProjects() {
        final List<GitHubRepositoryResponse> portfolioRepositories = new ArrayList<>();
        int page = 1;

        while (true) {
            URI pageUri = UriComponentsBuilder.fromUriString(repositoriesUrl)
                    .queryParam("per_page", PER_PAGE)
                    .queryParam("page", page)
                    .build()
                    .toUri();

            List<GitHubRepositoryResponse> pageRepositories = restClient.get()
                    .uri(pageUri)
                    .retrieve()
                    .body(new ParameterizedTypeReference<>() {
                    });

            final boolean noData = pageRepositories == null || pageRepositories.isEmpty();

            if (noData) {
                break;
            }

            pageRepositories.stream()
                    .filter(this::hasPortfolioTopic)
                    .forEach(portfolioRepositories::add);

            if (pageRepositories.size() < PER_PAGE) {
                break;
            }

            page++;
        }


        return GitHubProjectMapper.toDomainList(portfolioRepositories);
    }

    private boolean hasPortfolioTopic(GitHubRepositoryResponse repository) {
        return repository.topics() != null && repository.topics().contains(PORTFOLIO_TOPIC);
    }
}

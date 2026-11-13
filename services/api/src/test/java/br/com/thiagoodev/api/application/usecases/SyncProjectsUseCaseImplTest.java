package br.com.thiagoodev.api.application.usecases;

import br.com.thiagoodev.api.domain.entities.Project;
import br.com.thiagoodev.api.domain.entities.Skill;
import br.com.thiagoodev.api.domain.exceptions.ResourceNotFoundException;
import br.com.thiagoodev.api.domain.gateways.GitHubProjectGateway;
import br.com.thiagoodev.api.domain.repositories.ProjectRepository;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import java.util.List;
import java.util.UUID;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.mockito.ArgumentMatchers.anyList;
import static org.mockito.Mockito.never;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
class SyncProjectsUseCaseImplTest {

    @Mock
    private ProjectRepository projectRepository;

    @Mock
    private GitHubProjectGateway gitHubProjectGateway;

    @InjectMocks
    private SyncProjectsUseCaseImpl syncProjectsUseCase;

    @Test
    void shouldUpsertFetchedProjectsAndDeleteRemovedProjects_whenRepositoriesAreFetched() {
        Skill skill = new Skill(UUID.randomUUID(), "Java", "https://openjdk.org", "", null, null);
        UUID existingUuid = UUID.randomUUID();
        Project existingProject = new Project(existingUuid, 10L, "Old Name", "Old Desc", "", "", List.of("Go"), List.of(skill), null, null, null);
        Project removedProject = new Project(UUID.randomUUID(), 20L, "Removed", "Desc", "", "", List.of("Rust"), List.of(), null, null, null);
        Project fetchedUpdated = new Project(null, 10L, "New Name", "New Desc", "https://app.dev", "https://github.com/repo", List.of("Java"), List.of(), null, null, null);
        Project fetchedNew = new Project(null, 30L, "Brand New", "Fresh Desc", "", "https://github.com/new", List.of("TypeScript"), List.of(), null, null, null);
        when(gitHubProjectGateway.fetchPortfolioRepositories()).thenReturn(List.of(fetchedUpdated, fetchedNew));
        when(projectRepository.findAll()).thenReturn(List.of(existingProject, removedProject));
        when(projectRepository.saveAll(anyList())).thenAnswer(invocation -> invocation.getArgument(0));

        List<Project> result = syncProjectsUseCase.execute();

        assertEquals(2, result.size());
        assertEquals(existingUuid, existingProject.getUuid());
        assertEquals("New Name", existingProject.getName());
        assertEquals(1, existingProject.getSkills().size());
        verify(projectRepository, times(1)).deleteAll(List.of(removedProject));
        verify(projectRepository, times(1)).saveAll(anyList());
    }

    @Test
    void shouldThrowResourceNotFoundException_whenGitHubReturnsNoPortfolioRepositories() {
        when(gitHubProjectGateway.fetchPortfolioRepositories()).thenReturn(List.of());

        assertThrows(ResourceNotFoundException.class, () -> syncProjectsUseCase.execute());
        verify(projectRepository, never()).deleteAll(anyList());
        verify(projectRepository, never()).saveAll(anyList());
    }
}

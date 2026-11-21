package br.com.thiagoodev.api.application.usecases;

import br.com.thiagoodev.api.domain.entities.Project;
import br.com.thiagoodev.api.domain.entities.Skill;
import br.com.thiagoodev.api.domain.gateways.ProjectsGateway;
import br.com.thiagoodev.api.domain.repositories.ProjectRepository;
import br.com.thiagoodev.api.domain.repositories.SkillRepository;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import java.util.List;
import java.util.UUID;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;
import static org.mockito.ArgumentMatchers.any;
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
    private ProjectsGateway projectsGateway;

    @Mock
    private SkillRepository skillRepository;

    @InjectMocks
    private SyncProjectsUseCaseImpl syncProjectsUseCase;

    @Test
    void shouldUpsertFetchedProjectsAndDeleteRemovedProjects_whenRepositoriesAreFetched() {
        Skill javaSkill = new Skill(UUID.randomUUID(), "Java", "https://openjdk.org", "", null, null);
        UUID existingUuid = UUID.randomUUID();
        Project existingProject = new Project(existingUuid, 10L, "Old Name", "Old Desc", "", "", List.of("Go"), List.of(), null, null, null);
        Project fetchedUpdated = new Project(null, 10L, "New Name", "New Desc", "https://app.dev", "https://github.com/repo", List.of("Java"), List.of(), null, null, null);
        Project fetchedNew = new Project(null, 30L, "Brand New", "Fresh Desc", "", "https://github.com/new", List.of("TypeScript"), List.of(), null, null, null);
        when(projectsGateway.fetchProjects()).thenReturn(List.of(fetchedUpdated, fetchedNew));
        when(projectRepository.findAllByExternalIdIn(List.of(10L, 30L))).thenReturn(List.of(existingProject));
        when(skillRepository.findAllByLabelIn(List.of("Java"))).thenReturn(List.of(javaSkill));
        when(skillRepository.findAllByLabelIn(List.of("TypeScript"))).thenReturn(List.of());
        when(projectRepository.saveAll(any())).thenAnswer(invocation -> List.copyOf(invocation.getArgument(0)));

        List<Project> result = syncProjectsUseCase.execute();

        assertEquals(2, result.size());
        assertEquals(existingUuid, existingProject.getUuid());
        assertEquals("New Name", existingProject.getName());
        assertEquals(1, existingProject.getSkills().size());
        verify(projectRepository, times(1)).deleteAllByExternalIdNotNullAndExternalIdNotIn(List.of(10L, 30L));
        verify(projectRepository, times(1)).saveAll(any());
    }

    @Test
    void shouldReturnEmptyList_whenGitHubReturnsNoPortfolioRepositories() {
        when(projectsGateway.fetchProjects()).thenReturn(List.of());

        List<Project> result = syncProjectsUseCase.execute();

        assertTrue(result.isEmpty());
        verify(projectRepository, never()).deleteAllByExternalIdNotNullAndExternalIdNotIn(anyList());
        verify(projectRepository, never()).saveAll(any());
    }
}

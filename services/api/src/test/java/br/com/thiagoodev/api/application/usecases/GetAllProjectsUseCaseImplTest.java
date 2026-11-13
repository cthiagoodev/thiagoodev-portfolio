package br.com.thiagoodev.api.application.usecases;

import br.com.thiagoodev.api.domain.entities.Project;
import br.com.thiagoodev.api.domain.exceptions.ResourceNotFoundException;
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
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
class GetAllProjectsUseCaseImplTest {

    @Mock
    private ProjectRepository projectRepository;

    @InjectMocks
    private GetAllProjectsUseCaseImpl getAllProjectsUseCase;

    @Test
    void shouldReturnAllProjects_whenProjectsExist() {
        Project project = new Project(UUID.randomUUID(), 1L, "Portfolio", "Desc", "", "", null, null, null, null, null);
        List<Project> expectedProjects = List.of(project);
        when(projectRepository.findAll()).thenReturn(expectedProjects);

        List<Project> result = getAllProjectsUseCase.execute();

        assertEquals(expectedProjects, result);
        verify(projectRepository, times(1)).findAll();
    }

    @Test
    void shouldThrowResourceNotFoundException_whenNoProjectsExist() {
        when(projectRepository.findAll()).thenReturn(List.of());

        assertThrows(ResourceNotFoundException.class, () -> getAllProjectsUseCase.execute());
        verify(projectRepository, times(1)).findAll();
    }
}

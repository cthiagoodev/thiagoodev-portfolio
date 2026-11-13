package br.com.thiagoodev.api.application.usecases;

import br.com.thiagoodev.api.domain.entities.Project;
import br.com.thiagoodev.api.domain.exceptions.ResourceNotFoundException;
import br.com.thiagoodev.api.domain.repositories.ProjectRepository;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import java.util.Optional;
import java.util.UUID;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
class GetProjectByUuidUseCaseImplTest {

    @Mock
    private ProjectRepository projectRepository;

    @InjectMocks
    private GetProjectByUuidUseCaseImpl getProjectByUuidUseCase;

    @Test
    void shouldReturnProject_whenProjectExists() {
        UUID uuid = UUID.randomUUID();
        Project expectedProject = new Project(uuid, 1L, "Portfolio", "Desc", "", "", null, null, null, null, null);
        when(projectRepository.findById(uuid)).thenReturn(Optional.of(expectedProject));

        Project result = getProjectByUuidUseCase.execute(uuid);

        assertEquals(expectedProject, result);
        verify(projectRepository, times(1)).findById(uuid);
    }

    @Test
    void shouldThrowResourceNotFoundException_whenProjectDoesNotExist() {
        UUID uuid = UUID.randomUUID();
        when(projectRepository.findById(uuid)).thenReturn(Optional.empty());

        assertThrows(ResourceNotFoundException.class, () -> getProjectByUuidUseCase.execute(uuid));
        verify(projectRepository, times(1)).findById(uuid);
    }
}

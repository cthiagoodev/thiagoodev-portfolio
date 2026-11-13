package br.com.thiagoodev.api.application.usecases;

import br.com.thiagoodev.api.domain.entities.WorkExperience;
import br.com.thiagoodev.api.domain.exceptions.ResourceNotFoundException;
import br.com.thiagoodev.api.domain.repositories.WorkExperienceRepository;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import java.time.LocalDate;
import java.util.Optional;
import java.util.UUID;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
class GetWorkExperienceByUuidUseCaseImplTest {

    @Mock
    private WorkExperienceRepository workExperienceRepository;

    @InjectMocks
    private GetWorkExperienceByUuidUseCaseImpl getWorkExperienceByUuidUseCase;

    @Test
    void shouldReturnWorkExperience_whenExperienceExists() {
        UUID uuid = UUID.randomUUID();
        WorkExperience expectedExperience = new WorkExperience(
                uuid,
                "Software Engineer",
                "Backend",
                "Company",
                LocalDate.of(2022, 1, 1),
                null,
                null,
                null
        );
        when(workExperienceRepository.findById(uuid)).thenReturn(Optional.of(expectedExperience));

        WorkExperience result = getWorkExperienceByUuidUseCase.execute(uuid);

        assertEquals(expectedExperience, result);
        verify(workExperienceRepository, times(1)).findById(uuid);
    }

    @Test
    void shouldThrowResourceNotFoundException_whenExperienceDoesNotExist() {
        UUID uuid = UUID.randomUUID();
        when(workExperienceRepository.findById(uuid)).thenReturn(Optional.empty());

        assertThrows(ResourceNotFoundException.class, () -> getWorkExperienceByUuidUseCase.execute(uuid));
        verify(workExperienceRepository, times(1)).findById(uuid);
    }
}

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
import java.util.List;
import java.util.UUID;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
class GetAllWorkExperiencesUseCaseImplTest {

    @Mock
    private WorkExperienceRepository workExperienceRepository;

    @InjectMocks
    private GetAllWorkExperiencesUseCaseImpl getAllWorkExperiencesUseCase;

    @Test
    void shouldReturnAllWorkExperiences_whenExperiencesExist() {
        WorkExperience experience = new WorkExperience(
                UUID.randomUUID(),
                "Software Engineer",
                "Backend",
                "Company",
                LocalDate.of(2022, 1, 1),
                null,
                null,
                null
        );
        List<WorkExperience> expectedExperiences = List.of(experience);
        when(workExperienceRepository.findAll()).thenReturn(expectedExperiences);

        List<WorkExperience> result = getAllWorkExperiencesUseCase.execute();

        assertEquals(expectedExperiences, result);
        verify(workExperienceRepository, times(1)).findAll();
    }

    @Test
    void shouldThrowResourceNotFoundException_whenNoExperiencesExist() {
        when(workExperienceRepository.findAll()).thenReturn(List.of());

        assertThrows(ResourceNotFoundException.class, () -> getAllWorkExperiencesUseCase.execute());
        verify(workExperienceRepository, times(1)).findAll();
    }
}

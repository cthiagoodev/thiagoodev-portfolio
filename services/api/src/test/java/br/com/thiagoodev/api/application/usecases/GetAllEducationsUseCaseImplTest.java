package br.com.thiagoodev.api.application.usecases;

import br.com.thiagoodev.api.domain.entities.Education;
import br.com.thiagoodev.api.domain.exceptions.ResourceNotFoundException;
import br.com.thiagoodev.api.domain.repositories.EducationRepository;
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
class GetAllEducationsUseCaseImplTest {

    @Mock
    private EducationRepository educationRepository;

    @InjectMocks
    private GetAllEducationsUseCaseImpl getAllEducationsUseCase;

    @Test
    void shouldReturnAllEducations_whenEducationsExist() {
        Education education = new Education(
                UUID.randomUUID(),
                "Ciência da Computação",
                "Bacharelado",
                "Universidade",
                LocalDate.of(2018, 2, 1),
                null,
                null,
                null
        );
        List<Education> expectedEducations = List.of(education);
        when(educationRepository.findAll()).thenReturn(expectedEducations);

        List<Education> result = getAllEducationsUseCase.execute();

        assertEquals(expectedEducations, result);
        verify(educationRepository, times(1)).findAll();
    }

    @Test
    void shouldThrowResourceNotFoundException_whenNoEducationsExist() {
        when(educationRepository.findAll()).thenReturn(List.of());

        assertThrows(ResourceNotFoundException.class, () -> getAllEducationsUseCase.execute());
        verify(educationRepository, times(1)).findAll();
    }
}

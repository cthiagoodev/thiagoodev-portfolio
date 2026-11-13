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
import java.util.Optional;
import java.util.UUID;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
class GetEducationByUuidUseCaseImplTest {

    @Mock
    private EducationRepository educationRepository;

    @InjectMocks
    private GetEducationByUuidUseCaseImpl getEducationByUuidUseCase;

    @Test
    void shouldReturnEducation_whenEducationExists() {
        UUID uuid = UUID.randomUUID();
        Education expectedEducation = new Education(
                uuid,
                "Ciência da Computação",
                "Bacharelado",
                "Universidade",
                LocalDate.of(2018, 2, 1),
                null,
                null,
                null
        );
        when(educationRepository.findById(uuid)).thenReturn(Optional.of(expectedEducation));

        Education result = getEducationByUuidUseCase.execute(uuid);

        assertEquals(expectedEducation, result);
        verify(educationRepository, times(1)).findById(uuid);
    }

    @Test
    void shouldThrowResourceNotFoundException_whenEducationDoesNotExist() {
        UUID uuid = UUID.randomUUID();
        when(educationRepository.findById(uuid)).thenReturn(Optional.empty());

        assertThrows(ResourceNotFoundException.class, () -> getEducationByUuidUseCase.execute(uuid));
        verify(educationRepository, times(1)).findById(uuid);
    }
}

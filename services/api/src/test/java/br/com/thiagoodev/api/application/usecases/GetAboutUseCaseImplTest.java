package br.com.thiagoodev.api.application.usecases;

import br.com.thiagoodev.api.domain.entities.About;
import br.com.thiagoodev.api.domain.exceptions.ResourceNotFoundException;
import br.com.thiagoodev.api.domain.repositories.AboutRepository;
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
class GetAboutUseCaseImplTest {

    @Mock
    private AboutRepository aboutRepository;

    @InjectMocks
    private GetAboutUseCaseImpl getAboutUseCase;

    @Test
    void shouldReturnAbout_whenAboutExists() {
        About expectedAbout = new About(
                UUID.randomUUID(),
                "Thiago Sousa",
                "+5511999999999",
                "Software Engineer",
                "Bio completa",
                null,
                null
        );
        when(aboutRepository.findFirstByOrderByCreatedAtAsc()).thenReturn(Optional.of(expectedAbout));

        About result = getAboutUseCase.execute();

        assertEquals(expectedAbout, result);
        verify(aboutRepository, times(1)).findFirstByOrderByCreatedAtAsc();
    }

    @Test
    void shouldThrowResourceNotFoundException_whenAboutDoesNotExist() {
        when(aboutRepository.findFirstByOrderByCreatedAtAsc()).thenReturn(Optional.empty());

        assertThrows(ResourceNotFoundException.class, () -> getAboutUseCase.execute());
        verify(aboutRepository, times(1)).findFirstByOrderByCreatedAtAsc();
    }

    @Test
    void shouldPropagateException_whenRepositoryFails() {
        RuntimeException databaseException = new RuntimeException("database connection failed");
        when(aboutRepository.findFirstByOrderByCreatedAtAsc()).thenThrow(databaseException);

        RuntimeException thrown = assertThrows(RuntimeException.class, () -> getAboutUseCase.execute());

        assertEquals(databaseException, thrown);
        verify(aboutRepository, times(1)).findFirstByOrderByCreatedAtAsc();
    }
}

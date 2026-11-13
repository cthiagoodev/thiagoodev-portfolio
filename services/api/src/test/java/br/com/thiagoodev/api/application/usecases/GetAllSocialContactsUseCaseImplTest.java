package br.com.thiagoodev.api.application.usecases;

import br.com.thiagoodev.api.domain.entities.SocialContact;
import br.com.thiagoodev.api.domain.exceptions.ResourceNotFoundException;
import br.com.thiagoodev.api.domain.repositories.SocialContactRepository;
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
class GetAllSocialContactsUseCaseImplTest {

    @Mock
    private SocialContactRepository socialContactRepository;

    @InjectMocks
    private GetAllSocialContactsUseCaseImpl getAllSocialContactsUseCase;

    @Test
    void shouldReturnAllSocialContacts_whenContactsExist() {
        SocialContact firstContact = new SocialContact(
                UUID.randomUUID(),
                "github",
                "https://github.com/cthiagoodev",
                "/images/icons/github.svg",
                null,
                null
        );
        SocialContact secondContact = new SocialContact(
                UUID.randomUUID(),
                "linkedin",
                "https://linkedin.com/in/cthiagoodev",
                "/images/icons/linkedin.svg",
                null,
                null
        );
        List<SocialContact> expectedContacts = List.of(firstContact, secondContact);
        when(socialContactRepository.findAll()).thenReturn(expectedContacts);

        List<SocialContact> result = getAllSocialContactsUseCase.execute();

        assertEquals(expectedContacts, result);
        verify(socialContactRepository, times(1)).findAll();
    }

    @Test
    void shouldThrowResourceNotFoundException_whenNoContactsExist() {
        when(socialContactRepository.findAll()).thenReturn(List.of());

        assertThrows(ResourceNotFoundException.class, () -> getAllSocialContactsUseCase.execute());
        verify(socialContactRepository, times(1)).findAll();
    }

    @Test
    void shouldPropagateException_whenRepositoryFails() {
        RuntimeException databaseException = new RuntimeException("database connection failed");
        when(socialContactRepository.findAll()).thenThrow(databaseException);

        RuntimeException thrown = assertThrows(RuntimeException.class, () -> getAllSocialContactsUseCase.execute());

        assertEquals(databaseException, thrown);
        verify(socialContactRepository, times(1)).findAll();
    }
}

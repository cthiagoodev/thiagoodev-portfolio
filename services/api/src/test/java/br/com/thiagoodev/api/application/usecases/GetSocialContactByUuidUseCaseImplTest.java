package br.com.thiagoodev.api.application.usecases;

import br.com.thiagoodev.api.domain.entities.SocialContact;
import br.com.thiagoodev.api.domain.exceptions.ResourceNotFoundException;
import br.com.thiagoodev.api.domain.repositories.SocialContactRepository;
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
class GetSocialContactByUuidUseCaseImplTest {

    @Mock
    private SocialContactRepository socialContactRepository;

    @InjectMocks
    private GetSocialContactByUuidUseCaseImpl getSocialContactByUuidUseCase;

    @Test
    void shouldReturnSocialContact_whenContactExists() {
        UUID uuid = UUID.randomUUID();
        SocialContact expectedContact = new SocialContact(uuid, "github", "https://github.com/cthiagoodev", "", null, null);
        when(socialContactRepository.findById(uuid)).thenReturn(Optional.of(expectedContact));

        SocialContact result = getSocialContactByUuidUseCase.execute(uuid);

        assertEquals(expectedContact, result);
        verify(socialContactRepository, times(1)).findById(uuid);
    }

    @Test
    void shouldThrowResourceNotFoundException_whenContactDoesNotExist() {
        UUID uuid = UUID.randomUUID();
        when(socialContactRepository.findById(uuid)).thenReturn(Optional.empty());

        assertThrows(ResourceNotFoundException.class, () -> getSocialContactByUuidUseCase.execute(uuid));
        verify(socialContactRepository, times(1)).findById(uuid);
    }
}

package br.com.thiagoodev.api.application.usecases;

import br.com.thiagoodev.api.domain.entities.Skill;
import br.com.thiagoodev.api.domain.exceptions.ResourceNotFoundException;
import br.com.thiagoodev.api.domain.repositories.SkillRepository;
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
class GetSkillByUuidUseCaseImplTest {

    @Mock
    private SkillRepository skillRepository;

    @InjectMocks
    private GetSkillByUuidUseCaseImpl getSkillByUuidUseCase;

    @Test
    void shouldReturnSkill_whenSkillExists() {
        UUID uuid = UUID.randomUUID();
        Skill expectedSkill = new Skill(uuid, "Java", "https://openjdk.org", "", null, null);
        when(skillRepository.findById(uuid)).thenReturn(Optional.of(expectedSkill));

        Skill result = getSkillByUuidUseCase.execute(uuid);

        assertEquals(expectedSkill, result);
        verify(skillRepository, times(1)).findById(uuid);
    }

    @Test
    void shouldThrowResourceNotFoundException_whenSkillDoesNotExist() {
        UUID uuid = UUID.randomUUID();
        when(skillRepository.findById(uuid)).thenReturn(Optional.empty());

        assertThrows(ResourceNotFoundException.class, () -> getSkillByUuidUseCase.execute(uuid));
        verify(skillRepository, times(1)).findById(uuid);
    }
}

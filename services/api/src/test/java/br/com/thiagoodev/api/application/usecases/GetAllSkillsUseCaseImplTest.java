package br.com.thiagoodev.api.application.usecases;

import br.com.thiagoodev.api.domain.entities.Skill;
import br.com.thiagoodev.api.domain.exceptions.ResourceNotFoundException;
import br.com.thiagoodev.api.domain.repositories.SkillRepository;
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
class GetAllSkillsUseCaseImplTest {

    @Mock
    private SkillRepository skillRepository;

    @InjectMocks
    private GetAllSkillsUseCaseImpl getAllSkillsUseCase;

    @Test
    void shouldReturnAllSkills_whenSkillsExist() {
        Skill skill = new Skill(UUID.randomUUID(), "Java", "https://openjdk.org", "/icons/java.svg", null, null);
        List<Skill> expectedSkills = List.of(skill);
        when(skillRepository.findAll()).thenReturn(expectedSkills);

        List<Skill> result = getAllSkillsUseCase.execute();

        assertEquals(expectedSkills, result);
        verify(skillRepository, times(1)).findAll();
    }

    @Test
    void shouldThrowResourceNotFoundException_whenNoSkillsExist() {
        when(skillRepository.findAll()).thenReturn(List.of());

        assertThrows(ResourceNotFoundException.class, () -> getAllSkillsUseCase.execute());
        verify(skillRepository, times(1)).findAll();
    }
}

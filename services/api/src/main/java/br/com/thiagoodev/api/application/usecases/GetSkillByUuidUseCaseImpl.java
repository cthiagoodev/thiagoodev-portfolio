package br.com.thiagoodev.api.application.usecases;

import br.com.thiagoodev.api.domain.entities.Skill;
import br.com.thiagoodev.api.domain.exceptions.ResourceNotFoundException;
import br.com.thiagoodev.api.domain.repositories.SkillRepository;
import br.com.thiagoodev.api.domain.usecases.GetSkillByUuidUseCase;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.UUID;

@Service
@RequiredArgsConstructor
public class GetSkillByUuidUseCaseImpl implements GetSkillByUuidUseCase {

    private final SkillRepository skillRepository;

    @Override
    @Transactional(readOnly = true)
    public Skill execute(UUID uuid) {
        return skillRepository.findById(uuid)
                .orElseThrow(() -> new ResourceNotFoundException("resource not found"));
    }
}

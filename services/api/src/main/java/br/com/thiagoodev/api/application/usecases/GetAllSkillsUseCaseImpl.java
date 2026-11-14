package br.com.thiagoodev.api.application.usecases;

import br.com.thiagoodev.api.domain.entities.Skill;
import br.com.thiagoodev.api.domain.exceptions.ResourceNotFoundException;
import br.com.thiagoodev.api.domain.repositories.SkillRepository;
import br.com.thiagoodev.api.domain.usecases.GetAllSkillsUseCase;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Service
public class GetAllSkillsUseCaseImpl implements GetAllSkillsUseCase {

    private final SkillRepository skillRepository;

    public GetAllSkillsUseCaseImpl(SkillRepository skillRepository) {
        this.skillRepository = skillRepository;
    }

    @Override
    @Transactional(readOnly = true)
    public List<Skill> execute() {
        List<Skill> skills = skillRepository.findAll();
        if (skills.isEmpty()) {
            throw new ResourceNotFoundException("resource not found");
        }
        return skills;
    }
}

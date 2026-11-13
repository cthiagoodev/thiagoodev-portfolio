package br.com.thiagoodev.api.application.usecases;

import br.com.thiagoodev.api.domain.entities.WorkExperience;
import br.com.thiagoodev.api.domain.exceptions.ResourceNotFoundException;
import br.com.thiagoodev.api.domain.repositories.WorkExperienceRepository;
import br.com.thiagoodev.api.domain.usecases.GetAllWorkExperiencesUseCase;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Service
@RequiredArgsConstructor
public class GetAllWorkExperiencesUseCaseImpl implements GetAllWorkExperiencesUseCase {

    private final WorkExperienceRepository workExperienceRepository;

    @Override
    @Transactional(readOnly = true)
    public List<WorkExperience> execute() {
        List<WorkExperience> workExperiences = workExperienceRepository.findAll();
        if (workExperiences.isEmpty()) {
            throw new ResourceNotFoundException("resource not found");
        }
        return workExperiences;
    }
}

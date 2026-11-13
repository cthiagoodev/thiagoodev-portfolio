package br.com.thiagoodev.api.application.usecases;

import br.com.thiagoodev.api.domain.entities.WorkExperience;
import br.com.thiagoodev.api.domain.exceptions.ResourceNotFoundException;
import br.com.thiagoodev.api.domain.repositories.WorkExperienceRepository;
import br.com.thiagoodev.api.domain.usecases.GetWorkExperienceByUuidUseCase;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.UUID;

@Service
@RequiredArgsConstructor
public class GetWorkExperienceByUuidUseCaseImpl implements GetWorkExperienceByUuidUseCase {

    private final WorkExperienceRepository workExperienceRepository;

    @Override
    @Transactional(readOnly = true)
    public WorkExperience execute(UUID uuid) {
        return workExperienceRepository.findById(uuid)
                .orElseThrow(() -> new ResourceNotFoundException("resource not found"));
    }
}

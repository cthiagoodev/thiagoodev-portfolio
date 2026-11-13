package br.com.thiagoodev.api.application.usecases;

import br.com.thiagoodev.api.domain.entities.Education;
import br.com.thiagoodev.api.domain.exceptions.ResourceNotFoundException;
import br.com.thiagoodev.api.domain.repositories.EducationRepository;
import br.com.thiagoodev.api.domain.usecases.GetEducationByUuidUseCase;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.UUID;

@Service
@RequiredArgsConstructor
public class GetEducationByUuidUseCaseImpl implements GetEducationByUuidUseCase {

    private final EducationRepository educationRepository;

    @Override
    @Transactional(readOnly = true)
    public Education execute(UUID uuid) {
        return educationRepository.findById(uuid)
                .orElseThrow(() -> new ResourceNotFoundException("resource not found"));
    }
}

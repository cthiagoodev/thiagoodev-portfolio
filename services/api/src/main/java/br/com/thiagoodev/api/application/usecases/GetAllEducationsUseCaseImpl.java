package br.com.thiagoodev.api.application.usecases;

import br.com.thiagoodev.api.domain.entities.Education;
import br.com.thiagoodev.api.domain.exceptions.ResourceNotFoundException;
import br.com.thiagoodev.api.domain.repositories.EducationRepository;
import br.com.thiagoodev.api.domain.usecases.GetAllEducationsUseCase;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Service
@RequiredArgsConstructor
public class GetAllEducationsUseCaseImpl implements GetAllEducationsUseCase {

    private final EducationRepository educationRepository;

    @Override
    @Transactional(readOnly = true)
    public List<Education> execute() {
        List<Education> educations = educationRepository.findAll();
        if (educations.isEmpty()) {
            throw new ResourceNotFoundException("resource not found");
        }
        return educations;
    }
}

package br.com.thiagoodev.api.application.usecases;

import br.com.thiagoodev.api.domain.entities.Project;
import br.com.thiagoodev.api.domain.exceptions.ResourceNotFoundException;
import br.com.thiagoodev.api.domain.repositories.ProjectRepository;
import br.com.thiagoodev.api.domain.usecases.GetProjectByUuidUseCase;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.UUID;

@Service
@RequiredArgsConstructor
public class GetProjectByUuidUseCaseImpl implements GetProjectByUuidUseCase {

    private final ProjectRepository projectRepository;

    @Override
    @Transactional(readOnly = true)
    public Project execute(UUID uuid) {
        return projectRepository.findById(uuid)
                .orElseThrow(() -> new ResourceNotFoundException("resource not found"));
    }
}

package br.com.thiagoodev.api.application.usecases;

import br.com.thiagoodev.api.domain.entities.Project;
import br.com.thiagoodev.api.domain.exceptions.ResourceNotFoundException;
import br.com.thiagoodev.api.domain.repositories.ProjectRepository;
import br.com.thiagoodev.api.domain.usecases.GetAllProjectsUseCase;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Service
@RequiredArgsConstructor
public class GetAllProjectsUseCaseImpl implements GetAllProjectsUseCase {

    private final ProjectRepository projectRepository;

    @Override
    @Transactional(readOnly = true)
    public List<Project> execute() {
        List<Project> projects = projectRepository.findAll();
        if (projects.isEmpty()) {
            throw new ResourceNotFoundException("resource not found");
        }
        return projects;
    }
}

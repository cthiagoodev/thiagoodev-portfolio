package br.com.thiagoodev.api.application.usecases;

import br.com.thiagoodev.api.domain.entities.Project;
import br.com.thiagoodev.api.domain.exceptions.ResourceNotFoundException;
import br.com.thiagoodev.api.domain.gateways.ProjectsGateway;
import br.com.thiagoodev.api.domain.repositories.ProjectRepository;
import br.com.thiagoodev.api.domain.usecases.SyncProjectsUseCase;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Set;
import java.util.function.Function;
import java.util.stream.Collectors;

@Service
public class SyncProjectsUseCaseImpl implements SyncProjectsUseCase {

    private final ProjectRepository projectRepository;
    private final ProjectsGateway projectsGateway;

    public SyncProjectsUseCaseImpl(ProjectRepository projectRepository, ProjectsGateway projectsGateway) {
        this.projectRepository = projectRepository;
        this.projectsGateway = projectsGateway;
    }

    @Override
    @Transactional
    public List<Project> execute() {
        List<Project> fetchedProjects = projectsGateway.fetchProjects();
        if (fetchedProjects.isEmpty()) {
            throw new ResourceNotFoundException("no portfolio repositories found to sync");
        }

        List<Project> existingProjects = projectRepository.findAll();
        Map<Long, Project> existingByExternalId = existingProjects.stream()
                .filter(project -> project.getExternalId() != null)
                .collect(Collectors.toMap(Project::getExternalId, Function.identity(), (first, second) -> first));

        Set<Long> fetchedExternalIds = fetchedProjects.stream()
                .map(Project::getExternalId)
                .filter(Objects::nonNull)
                .collect(Collectors.toSet());

        List<Project> projectsToDelete = existingProjects.stream()
                .filter(project -> project.getExternalId() != null && !fetchedExternalIds.contains(project.getExternalId()))
                .toList();

        if (!projectsToDelete.isEmpty()) {
            projectRepository.deleteAll(projectsToDelete);
        }

        List<Project> projectsToSave = new ArrayList<>();
        for (Project fetchedProject : fetchedProjects) {
            Project existingProject = existingByExternalId.get(fetchedProject.getExternalId());
            if (existingProject != null) {
                existingProject.syncFromGitHub(
                        fetchedProject.getName(),
                        fetchedProject.getDescription(),
                        fetchedProject.getUrl(),
                        fetchedProject.getRepoUrl(),
                        fetchedProject.getLanguages(),
                        fetchedProject.getPushedAt(),
                        fetchedProject.getUpdatedAt()
                );
                projectsToSave.add(existingProject);
            } else {
                projectsToSave.add(fetchedProject);
            }
        }

        return projectRepository.saveAll(projectsToSave);
    }
}

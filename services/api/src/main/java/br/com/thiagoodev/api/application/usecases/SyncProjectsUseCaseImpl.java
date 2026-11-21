package br.com.thiagoodev.api.application.usecases;

import br.com.thiagoodev.api.domain.entities.Project;
import br.com.thiagoodev.api.domain.entities.Skill;
import br.com.thiagoodev.api.domain.gateways.ProjectsGateway;
import br.com.thiagoodev.api.domain.repositories.ProjectRepository;
import br.com.thiagoodev.api.domain.repositories.SkillRepository;
import br.com.thiagoodev.api.domain.usecases.SyncProjectsUseCase;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;
import java.util.Map;
import java.util.function.Function;
import java.util.stream.Collectors;

@Service
public class SyncProjectsUseCaseImpl implements SyncProjectsUseCase {

    private final ProjectRepository projectRepository;
    private final ProjectsGateway projectsGateway;
    private final SkillRepository skillRepository;

    public SyncProjectsUseCaseImpl(
            ProjectRepository projectRepository,
            ProjectsGateway projectsGateway,
            SkillRepository skillRepository
    ) {
        this.projectRepository = projectRepository;
        this.projectsGateway = projectsGateway;
        this.skillRepository = skillRepository;
    }

    @Override
    @Transactional
    public List<Project> execute() {
        List<Project> remoteProjects = projectsGateway.fetchProjects();

        if (remoteProjects.isEmpty()) {
            return List.of();
        }

        List<Long> remoteExternalIds = this.getExternalIdList(remoteProjects);

        projectRepository.deleteAllByExternalIdNotNullAndExternalIdNotIn(remoteExternalIds);

        Map<Long, Project> projects = this.findAllByExternalIdIn(remoteExternalIds);

        remoteProjects.forEach(item -> {
            Project project = projects.get(item.getExternalId());

            if (project == null) {
                this.addProject(item, projects);
            } else {
                this.setProject(item, project);
            }
        });

        return projectRepository.saveAll(projects.values());
    }

    private List<Long> getExternalIdList(List<Project> projects) {
        return projects.stream().map(Project::getExternalId).toList();
    }

    private Map<Long, Project> findAllByExternalIdIn(List<Long> ids) {
        return projectRepository
                .findAllByExternalIdIn(ids)
                .stream()
                .collect(Collectors.toMap(
                        Project::getExternalId,
                        Function.identity(),
                        (project, project2) -> project));
    }

    private void addProject(Project item, Map<Long, Project> projects) {
        item.setSkills(this.getSkills(item.getLanguages()));
        projects.put(item.getExternalId(), item);
    }

    private void setProject(Project item, Project project) {
        project.setName(item.getName());
        project.setDescription(item.getDescription());
        project.setUrl(item.getUrl());
        project.setRepoUrl(item.getRepoUrl());
        project.setLanguages(item.getLanguages());
        project.setSkills(this.getSkills(item.getLanguages()));
    }

    private List<Skill> getSkills(List<String> languages) {
        return skillRepository.findAllByLabelIn(languages);
    }
}

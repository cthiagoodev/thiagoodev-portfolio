package br.com.thiagoodev.api.domain.gateways;

import br.com.thiagoodev.api.domain.entities.Project;

import java.util.List;

public interface ProjectsGateway {

    List<Project> fetchProjects();
}

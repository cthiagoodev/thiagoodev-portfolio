package br.com.thiagoodev.api.domain.usecases;

import br.com.thiagoodev.api.domain.entities.Project;

import java.util.List;

public interface SyncProjectsUseCase {

    List<Project> execute();
}

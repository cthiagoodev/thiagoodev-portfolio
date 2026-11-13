package br.com.thiagoodev.api.domain.usecases;

import br.com.thiagoodev.api.domain.entities.Project;

import java.util.UUID;

public interface GetProjectByUuidUseCase {

    Project execute(UUID uuid);
}

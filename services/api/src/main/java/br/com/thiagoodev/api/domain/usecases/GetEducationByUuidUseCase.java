package br.com.thiagoodev.api.domain.usecases;

import br.com.thiagoodev.api.domain.entities.Education;

import java.util.UUID;

public interface GetEducationByUuidUseCase {

    Education execute(UUID uuid);
}

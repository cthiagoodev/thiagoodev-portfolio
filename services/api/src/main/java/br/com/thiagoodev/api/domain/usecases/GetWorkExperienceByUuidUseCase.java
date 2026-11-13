package br.com.thiagoodev.api.domain.usecases;

import br.com.thiagoodev.api.domain.entities.WorkExperience;

import java.util.UUID;

public interface GetWorkExperienceByUuidUseCase {

    WorkExperience execute(UUID uuid);
}

package br.com.thiagoodev.api.domain.usecases;

import br.com.thiagoodev.api.domain.entities.WorkExperience;

import java.util.List;

public interface GetAllWorkExperiencesUseCase {

    List<WorkExperience> execute();
}

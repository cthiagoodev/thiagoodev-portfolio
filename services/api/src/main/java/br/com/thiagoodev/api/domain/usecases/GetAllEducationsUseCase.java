package br.com.thiagoodev.api.domain.usecases;

import br.com.thiagoodev.api.domain.entities.Education;

import java.util.List;

public interface GetAllEducationsUseCase {

    List<Education> execute();
}

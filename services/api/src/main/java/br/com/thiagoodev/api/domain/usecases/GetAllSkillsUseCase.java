package br.com.thiagoodev.api.domain.usecases;

import br.com.thiagoodev.api.domain.entities.Skill;

import java.util.List;

public interface GetAllSkillsUseCase {

    List<Skill> execute();
}

package br.com.thiagoodev.api.domain.usecases;

import br.com.thiagoodev.api.domain.entities.Skill;

import java.util.UUID;

public interface GetSkillByUuidUseCase {

    Skill execute(UUID uuid);
}

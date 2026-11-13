package br.com.thiagoodev.api.domain.usecases;

import br.com.thiagoodev.api.domain.entities.SocialContact;

import java.util.UUID;

public interface GetSocialContactByUuidUseCase {

    SocialContact execute(UUID uuid);
}

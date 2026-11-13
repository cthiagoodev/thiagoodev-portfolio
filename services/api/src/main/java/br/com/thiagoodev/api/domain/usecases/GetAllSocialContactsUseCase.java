package br.com.thiagoodev.api.domain.usecases;

import br.com.thiagoodev.api.domain.entities.SocialContact;

import java.util.List;

public interface GetAllSocialContactsUseCase {

    List<SocialContact> execute();
}

package br.com.thiagoodev.api.application.usecases;

import br.com.thiagoodev.api.domain.entities.SocialContact;
import br.com.thiagoodev.api.domain.exceptions.ResourceNotFoundException;
import br.com.thiagoodev.api.domain.repositories.SocialContactRepository;
import br.com.thiagoodev.api.domain.usecases.GetAllSocialContactsUseCase;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Service
public class GetAllSocialContactsUseCaseImpl implements GetAllSocialContactsUseCase {

    private final SocialContactRepository socialContactRepository;

    public GetAllSocialContactsUseCaseImpl(SocialContactRepository socialContactRepository) {
        this.socialContactRepository = socialContactRepository;
    }

    @Override
    @Transactional(readOnly = true)
    public List<SocialContact> execute() {
        List<SocialContact> socialContacts = socialContactRepository.findAll();
        if (socialContacts.isEmpty()) {
            throw new ResourceNotFoundException("resource not found");
        }
        return socialContacts;
    }
}

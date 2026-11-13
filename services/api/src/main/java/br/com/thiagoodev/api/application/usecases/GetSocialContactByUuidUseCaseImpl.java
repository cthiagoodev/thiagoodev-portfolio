package br.com.thiagoodev.api.application.usecases;

import br.com.thiagoodev.api.domain.entities.SocialContact;
import br.com.thiagoodev.api.domain.exceptions.ResourceNotFoundException;
import br.com.thiagoodev.api.domain.repositories.SocialContactRepository;
import br.com.thiagoodev.api.domain.usecases.GetSocialContactByUuidUseCase;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.UUID;

@Service
@RequiredArgsConstructor
public class GetSocialContactByUuidUseCaseImpl implements GetSocialContactByUuidUseCase {

    private final SocialContactRepository socialContactRepository;

    @Override
    @Transactional(readOnly = true)
    public SocialContact execute(UUID uuid) {
        return socialContactRepository.findById(uuid)
                .orElseThrow(() -> new ResourceNotFoundException("resource not found"));
    }
}

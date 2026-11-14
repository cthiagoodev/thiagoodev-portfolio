package br.com.thiagoodev.api.application.usecases;

import br.com.thiagoodev.api.domain.entities.About;
import br.com.thiagoodev.api.domain.exceptions.ResourceNotFoundException;
import br.com.thiagoodev.api.domain.repositories.AboutRepository;
import br.com.thiagoodev.api.domain.usecases.GetAboutUseCase;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
public class GetAboutUseCaseImpl implements GetAboutUseCase {

    private final AboutRepository aboutRepository;

    public GetAboutUseCaseImpl(AboutRepository aboutRepository) {
        this.aboutRepository = aboutRepository;
    }

    @Override
    @Transactional(readOnly = true)
    public About execute() {
        return aboutRepository.findFirstByOrderByCreatedAtAsc()
                .orElseThrow(() -> new ResourceNotFoundException("resource not found"));
    }
}

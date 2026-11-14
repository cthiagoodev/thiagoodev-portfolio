package br.com.thiagoodev.api.presentation.controllers;

import br.com.thiagoodev.api.domain.entities.SocialContact;
import br.com.thiagoodev.api.domain.usecases.GetAllSocialContactsUseCase;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("/social-contacts")
public class SocialContactController {

    private final GetAllSocialContactsUseCase getAllSocialContactsUseCase;

    public SocialContactController(GetAllSocialContactsUseCase getAllSocialContactsUseCase) {
        this.getAllSocialContactsUseCase = getAllSocialContactsUseCase;
    }

    @GetMapping
    public ResponseEntity<List<SocialContact>> getAll() {
        List<SocialContact> socialContacts = getAllSocialContactsUseCase.execute();
        return ResponseEntity.ok(socialContacts);
    }
}

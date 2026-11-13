package br.com.thiagoodev.api.presentation.controllers;

import br.com.thiagoodev.api.domain.entities.SocialContact;
import br.com.thiagoodev.api.domain.usecases.GetAllSocialContactsUseCase;
import br.com.thiagoodev.api.domain.usecases.GetSocialContactByUuidUseCase;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;
import java.util.UUID;

@RestController
@RequestMapping("/social-contacts")
@RequiredArgsConstructor
public class SocialContactController {

    private final GetAllSocialContactsUseCase getAllSocialContactsUseCase;
    private final GetSocialContactByUuidUseCase getSocialContactByUuidUseCase;

    @GetMapping
    public ResponseEntity<List<SocialContact>> getAll() {
        List<SocialContact> socialContacts = getAllSocialContactsUseCase.execute();
        return ResponseEntity.ok(socialContacts);
    }

    @GetMapping("/{uuid}")
    public ResponseEntity<SocialContact> getByUuid(@PathVariable UUID uuid) {
        SocialContact socialContact = getSocialContactByUuidUseCase.execute(uuid);
        return ResponseEntity.ok(socialContact);
    }
}

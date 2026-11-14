package br.com.thiagoodev.api.presentation.controllers;

import br.com.thiagoodev.api.domain.entities.About;
import br.com.thiagoodev.api.domain.usecases.GetAboutUseCase;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/about")
public class AboutController {

    private final GetAboutUseCase getAboutUseCase;

    public AboutController(GetAboutUseCase getAboutUseCase) {
        this.getAboutUseCase = getAboutUseCase;
    }

    @GetMapping
    public ResponseEntity<About> get() {
        About about = getAboutUseCase.execute();
        return ResponseEntity.ok(about);
    }
}

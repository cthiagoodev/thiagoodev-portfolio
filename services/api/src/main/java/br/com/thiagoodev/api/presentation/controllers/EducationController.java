package br.com.thiagoodev.api.presentation.controllers;

import br.com.thiagoodev.api.domain.entities.Education;
import br.com.thiagoodev.api.domain.usecases.GetAllEducationsUseCase;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("/educations")
public class EducationController {

    private final GetAllEducationsUseCase getAllEducationsUseCase;

    public EducationController(GetAllEducationsUseCase getAllEducationsUseCase) {
        this.getAllEducationsUseCase = getAllEducationsUseCase;
    }

    @GetMapping
    public ResponseEntity<List<Education>> getAll() {
        List<Education> educations = getAllEducationsUseCase.execute();
        return ResponseEntity.ok(educations);
    }
}

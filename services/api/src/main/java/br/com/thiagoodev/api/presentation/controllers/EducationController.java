package br.com.thiagoodev.api.presentation.controllers;

import br.com.thiagoodev.api.domain.entities.Education;
import br.com.thiagoodev.api.domain.usecases.GetAllEducationsUseCase;
import br.com.thiagoodev.api.domain.usecases.GetEducationByUuidUseCase;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;
import java.util.UUID;

@RestController
@RequestMapping("/educations")
@RequiredArgsConstructor
public class EducationController {

    private final GetAllEducationsUseCase getAllEducationsUseCase;
    private final GetEducationByUuidUseCase getEducationByUuidUseCase;

    @GetMapping
    public ResponseEntity<List<Education>> getAll() {
        List<Education> educations = getAllEducationsUseCase.execute();
        return ResponseEntity.ok(educations);
    }

    @GetMapping("/{uuid}")
    public ResponseEntity<Education> getByUuid(@PathVariable UUID uuid) {
        Education education = getEducationByUuidUseCase.execute(uuid);
        return ResponseEntity.ok(education);
    }
}

package br.com.thiagoodev.api.presentation.controllers;

import br.com.thiagoodev.api.domain.entities.WorkExperience;
import br.com.thiagoodev.api.domain.usecases.GetAllWorkExperiencesUseCase;
import br.com.thiagoodev.api.domain.usecases.GetWorkExperienceByUuidUseCase;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;
import java.util.UUID;

@RestController
@RequestMapping("/work-experiences")
@RequiredArgsConstructor
public class WorkExperienceController {

    private final GetAllWorkExperiencesUseCase getAllWorkExperiencesUseCase;
    private final GetWorkExperienceByUuidUseCase getWorkExperienceByUuidUseCase;

    @GetMapping
    public ResponseEntity<List<WorkExperience>> getAll() {
        List<WorkExperience> workExperiences = getAllWorkExperiencesUseCase.execute();
        return ResponseEntity.ok(workExperiences);
    }

    @GetMapping("/{uuid}")
    public ResponseEntity<WorkExperience> getByUuid(@PathVariable UUID uuid) {
        WorkExperience workExperience = getWorkExperienceByUuidUseCase.execute(uuid);
        return ResponseEntity.ok(workExperience);
    }
}

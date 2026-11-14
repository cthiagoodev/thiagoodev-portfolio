package br.com.thiagoodev.api.presentation.controllers;

import br.com.thiagoodev.api.domain.entities.WorkExperience;
import br.com.thiagoodev.api.domain.usecases.GetAllWorkExperiencesUseCase;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("/work-experiences")
public class WorkExperienceController {

    private final GetAllWorkExperiencesUseCase getAllWorkExperiencesUseCase;

    public WorkExperienceController(GetAllWorkExperiencesUseCase getAllWorkExperiencesUseCase) {
        this.getAllWorkExperiencesUseCase = getAllWorkExperiencesUseCase;
    }

    @GetMapping
    public ResponseEntity<List<WorkExperience>> getAll() {
        List<WorkExperience> workExperiences = getAllWorkExperiencesUseCase.execute();
        return ResponseEntity.ok(workExperiences);
    }
}

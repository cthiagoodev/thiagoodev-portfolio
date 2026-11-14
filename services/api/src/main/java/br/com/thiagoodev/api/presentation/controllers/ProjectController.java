package br.com.thiagoodev.api.presentation.controllers;

import br.com.thiagoodev.api.domain.entities.Project;
import br.com.thiagoodev.api.domain.usecases.GetAllProjectsUseCase;
import br.com.thiagoodev.api.domain.usecases.SyncProjectsUseCase;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("/projects")
public class ProjectController {

    private final GetAllProjectsUseCase getAllProjectsUseCase;
    private final SyncProjectsUseCase syncProjectsUseCase;

    public ProjectController(GetAllProjectsUseCase getAllProjectsUseCase, SyncProjectsUseCase syncProjectsUseCase) {
        this.getAllProjectsUseCase = getAllProjectsUseCase;
        this.syncProjectsUseCase = syncProjectsUseCase;
    }

    @GetMapping
    public ResponseEntity<List<Project>> getAll() {
        List<Project> projects = getAllProjectsUseCase.execute();
        return ResponseEntity.ok(projects);
    }

    @PostMapping("/sync")
    public ResponseEntity<List<Project>> sync() {
        List<Project> syncedProjects = syncProjectsUseCase.execute();
        return ResponseEntity.ok(syncedProjects);
    }
}

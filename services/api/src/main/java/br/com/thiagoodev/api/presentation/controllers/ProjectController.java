package br.com.thiagoodev.api.presentation.controllers;

import br.com.thiagoodev.api.domain.entities.Project;
import br.com.thiagoodev.api.domain.usecases.GetAllProjectsUseCase;
import br.com.thiagoodev.api.domain.usecases.GetProjectByUuidUseCase;
import br.com.thiagoodev.api.domain.usecases.SyncProjectsUseCase;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;
import java.util.UUID;

@RestController
@RequestMapping("/projects")
@RequiredArgsConstructor
public class ProjectController {

    private final GetAllProjectsUseCase getAllProjectsUseCase;
    private final GetProjectByUuidUseCase getProjectByUuidUseCase;
    private final SyncProjectsUseCase syncProjectsUseCase;

    @GetMapping
    public ResponseEntity<List<Project>> getAll() {
        List<Project> projects = getAllProjectsUseCase.execute();
        return ResponseEntity.ok(projects);
    }

    @GetMapping("/{uuid}")
    public ResponseEntity<Project> getByUuid(@PathVariable UUID uuid) {
        Project project = getProjectByUuidUseCase.execute(uuid);
        return ResponseEntity.ok(project);
    }

    @PostMapping("/sync")
    public ResponseEntity<List<Project>> sync() {
        List<Project> syncedProjects = syncProjectsUseCase.execute();
        return ResponseEntity.ok(syncedProjects);
    }
}

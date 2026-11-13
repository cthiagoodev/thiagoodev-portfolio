package br.com.thiagoodev.api.presentation.controllers;

import br.com.thiagoodev.api.domain.entities.Skill;
import br.com.thiagoodev.api.domain.usecases.GetAllSkillsUseCase;
import br.com.thiagoodev.api.domain.usecases.GetSkillByUuidUseCase;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;
import java.util.UUID;

@RestController
@RequestMapping("/skills")
@RequiredArgsConstructor
public class SkillController {

    private final GetAllSkillsUseCase getAllSkillsUseCase;
    private final GetSkillByUuidUseCase getSkillByUuidUseCase;

    @GetMapping
    public ResponseEntity<List<Skill>> getAll() {
        List<Skill> skills = getAllSkillsUseCase.execute();
        return ResponseEntity.ok(skills);
    }

    @GetMapping("/{uuid}")
    public ResponseEntity<Skill> getByUuid(@PathVariable UUID uuid) {
        Skill skill = getSkillByUuidUseCase.execute(uuid);
        return ResponseEntity.ok(skill);
    }
}

package br.com.thiagoodev.api.presentation.controllers;

import br.com.thiagoodev.api.domain.entities.Skill;
import br.com.thiagoodev.api.domain.usecases.GetAllSkillsUseCase;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("/skills")
public class SkillController {

    private final GetAllSkillsUseCase getAllSkillsUseCase;

    public SkillController(GetAllSkillsUseCase getAllSkillsUseCase) {
        this.getAllSkillsUseCase = getAllSkillsUseCase;
    }

    @GetMapping
    public ResponseEntity<List<Skill>> getAll() {
        List<Skill> skills = getAllSkillsUseCase.execute();
        return ResponseEntity.ok(skills);
    }
}

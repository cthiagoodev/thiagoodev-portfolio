package br.com.thiagoodev.api.domain.entities;

import br.com.thiagoodev.api.domain.exceptions.DomainValidationException;
import org.junit.jupiter.api.Test;

import java.time.LocalDateTime;
import java.util.List;
import java.util.UUID;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.junit.jupiter.api.Assertions.assertTrue;

class ProjectTest {

    @Test
    void shouldCreateProject_whenAllFieldsAreValid() {
        String name = "Portfolio Web";
        String repoUrl = "https://github.com/cthiagoodev/portfolio";

        Project project = new Project(
                null,
                null,
                name,
                "Meu portfólio pessoal",
                "https://thiagoodev.com",
                repoUrl,
                null,
                null,
                null,
                null,
                null
        );

        assertEquals(name, project.getName());
        assertEquals(repoUrl, project.getRepoUrl());
        assertNotNull(project.getLanguages());
        assertTrue(project.getLanguages().isEmpty());
        assertNotNull(project.getSkills());
        assertTrue(project.getSkills().isEmpty());
        assertNotNull(project.getCreatedAt());
        assertNotNull(project.getUpdatedAt());
    }

    @Test
    void shouldThrowDomainValidationException_whenNameIsBlank() {
        DomainValidationException exception = assertThrows(
                DomainValidationException.class,
                () -> new Project(
                        null,
                        null,
                        "   ",
                        "Descrição",
                        "",
                        "",
                        List.of("Java"),
                        null,
                        null,
                        LocalDateTime.now(),
                        LocalDateTime.now()
                )
        );

        assertTrue(exception.getMessage().contains("project name is required"));
    }

    @Test
    void shouldAddSkill_whenSkillIsNotPresent() {
        UUID skillUuid = UUID.randomUUID();
        Project project = new Project(null, null, "Projeto", "", "", "", null, null, null, null, null);
        Skill skill = new Skill(skillUuid, "Java", "https://openjdk.org", "", null, null);

        project.addSkill(skill);

        assertEquals(1, project.getSkills().size());
        assertEquals(skillUuid, project.getSkills().getFirst().getUuid());
        assertEquals("Java", project.getSkills().getFirst().getLabel());
    }

    @Test
    void shouldIgnoreDuplicateSkill_whenAddingExistingSkill() {
        UUID skillUuid = UUID.randomUUID();
        Project project = new Project(null, null, "Projeto", "", "", "", null, null, null, null, null);
        Skill skill = new Skill(skillUuid, "Java", "https://openjdk.org", "", null, null);

        project.addSkill(skill);
        project.addSkill(skill);

        assertEquals(1, project.getSkills().size());
    }

    @Test
    void shouldRemoveSkill_whenSkillUuidMatches() {
        UUID skillUuid1 = UUID.randomUUID();
        UUID skillUuid2 = UUID.randomUUID();
        Skill skill1 = new Skill(skillUuid1, "Java", "https://openjdk.org", "", null, null);
        Skill skill2 = new Skill(skillUuid2, "Docker", "https://docker.com", "", null, null);
        Project project = new Project(null, null, "Projeto", "", "", "", null, List.of(skill1, skill2), null, null, null);

        project.removeSkill(skillUuid1);

        assertEquals(1, project.getSkills().size());
        assertEquals(skillUuid2, project.getSkills().getFirst().getUuid());
    }
}

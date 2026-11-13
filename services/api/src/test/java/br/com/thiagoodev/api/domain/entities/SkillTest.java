package br.com.thiagoodev.api.domain.entities;

import br.com.thiagoodev.api.domain.exceptions.DomainValidationException;
import org.junit.jupiter.api.Test;

import java.time.LocalDateTime;
import java.util.UUID;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.junit.jupiter.api.Assertions.assertTrue;

class SkillTest {

    @Test
    void shouldCreateSkill_whenAllFieldsAreValid() {
        UUID uuid = UUID.randomUUID();
        String label = "Java";
        String url = "https://openjdk.org";
        String imagePath = "/images/skills/java.svg";

        Skill skill = new Skill(
                uuid,
                label,
                url,
                imagePath,
                null,
                null
        );

        assertEquals(uuid, skill.getUuid());
        assertEquals(label, skill.getLabel());
        assertEquals(url, skill.getUrl());
        assertEquals(imagePath, skill.getImagePath());
        assertNotNull(skill.getCreatedAt());
        assertNotNull(skill.getUpdatedAt());
    }

    @Test
    void shouldThrowDomainValidationException_whenLabelIsBlank() {
        DomainValidationException exception = assertThrows(
                DomainValidationException.class,
                () -> new Skill(
                        null,
                        "   ",
                        "https://openjdk.org",
                        "/images/skills/java.svg",
                        LocalDateTime.now(),
                        LocalDateTime.now()
                )
        );

        assertTrue(exception.getMessage().contains("label is required"));
    }
}

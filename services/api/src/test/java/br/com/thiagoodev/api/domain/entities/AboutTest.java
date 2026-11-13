package br.com.thiagoodev.api.domain.entities;

import br.com.thiagoodev.api.domain.exceptions.DomainValidationException;
import org.junit.jupiter.api.Test;

import java.time.LocalDateTime;
import java.util.UUID;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.junit.jupiter.api.Assertions.assertTrue;

class AboutTest {

    @Test
    void shouldCreateAbout_whenAllFieldsAreValid() {
        UUID uuid = UUID.randomUUID();
        String name = "Thiago Sousa";
        String phone = "+5511999999999";
        String description = "Software Engineer";
        String text = "Bio completa sobre mim";

        About about = new About(
                uuid,
                name,
                phone,
                description,
                text,
                null,
                null
        );

        assertEquals(uuid, about.getUuid());
        assertEquals(name, about.getName());
        assertEquals(phone, about.getPhone());
        assertEquals(description, about.getDescription());
        assertEquals(text, about.getText());
        assertNotNull(about.getCreatedAt());
        assertNotNull(about.getUpdatedAt());
    }

    @Test
    void shouldThrowDomainValidationException_whenNameIsBlank() {
        DomainValidationException exception = assertThrows(
                DomainValidationException.class,
                () -> new About(
                        null,
                        "   ",
                        "+5511999999999",
                        "Descrição",
                        "Texto",
                        LocalDateTime.now(),
                        LocalDateTime.now()
                )
        );

        assertTrue(exception.getMessage().contains("name is required"));
    }

    @Test
    void shouldThrowDomainValidationException_whenPhoneExceedsMaxLength() {
        DomainValidationException exception = assertThrows(
                DomainValidationException.class,
                () -> new About(
                        null,
                        "Thiago",
                        "+5511999999999999999999999",
                        "Descrição",
                        "Texto",
                        LocalDateTime.now(),
                        LocalDateTime.now()
                )
        );

        assertTrue(exception.getMessage().contains("phone must have at most 20 characters"));
    }
}

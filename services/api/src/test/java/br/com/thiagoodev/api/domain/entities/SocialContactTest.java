package br.com.thiagoodev.api.domain.entities;

import br.com.thiagoodev.api.domain.exceptions.DomainValidationException;
import org.junit.jupiter.api.Test;

import java.time.LocalDateTime;
import java.util.UUID;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.junit.jupiter.api.Assertions.assertTrue;

class SocialContactTest {

    @Test
    void shouldCreateSocialContact_whenAllFieldsAreValid() {
        UUID uuid = UUID.randomUUID();
        String label = "github";
        String url = "https://github.com/cthiagoodev";
        String imagePath = "/images/icons/github.svg";

        SocialContact contact = new SocialContact(
                uuid,
                label,
                url,
                imagePath,
                null,
                null
        );

        assertEquals(uuid, contact.getUuid());
        assertEquals(label, contact.getLabel());
        assertEquals(url, contact.getUrl());
        assertEquals(imagePath, contact.getImagePath());
        assertNotNull(contact.getCreatedAt());
        assertNotNull(contact.getUpdatedAt());
    }

    @Test
    void shouldThrowDomainValidationException_whenLabelIsBlank() {
        DomainValidationException exception = assertThrows(
                DomainValidationException.class,
                () -> new SocialContact(
                        null,
                        "   ",
                        "https://github.com/cthiagoodev",
                        "",
                        LocalDateTime.now(),
                        LocalDateTime.now()
                )
        );

        assertTrue(exception.getMessage().contains("label is required"));
    }

    @Test
    void shouldThrowDomainValidationException_whenUrlIsBlank() {
        DomainValidationException exception = assertThrows(
                DomainValidationException.class,
                () -> new SocialContact(
                        null,
                        "github",
                        "   ",
                        "",
                        LocalDateTime.now(),
                        LocalDateTime.now()
                )
        );

        assertTrue(exception.getMessage().contains("url is required"));
    }
}

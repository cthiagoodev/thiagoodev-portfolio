package br.com.thiagoodev.api.domain.entities;

import br.com.thiagoodev.api.domain.exceptions.DomainValidationException;
import org.junit.jupiter.api.Test;

import java.time.LocalDate;
import java.time.LocalDateTime;
import java.util.UUID;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertNull;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.junit.jupiter.api.Assertions.assertTrue;

class WorkExperienceTest {

    @Test
    void shouldCreateWorkExperience_whenAllFieldsAreValid() {
        UUID uuid = UUID.randomUUID();
        LocalDate startDate = LocalDate.of(2022, 1, 1);
        LocalDate endDate = LocalDate.of(2023, 1, 1);

        WorkExperience workExperience = new WorkExperience(
                uuid,
                "Software Engineer",
                "Desenvolvimento backend em Java e Spring Boot",
                "Tech Company",
                startDate,
                endDate,
                null,
                null
        );

        assertEquals(uuid, workExperience.getUuid());
        assertEquals("Software Engineer", workExperience.getRole());
        assertEquals("Desenvolvimento backend em Java e Spring Boot", workExperience.getDescription());
        assertEquals("Tech Company", workExperience.getCompany());
        assertEquals(startDate, workExperience.getStartDate());
        assertEquals(endDate, workExperience.getEndDate());
        assertNotNull(workExperience.getCreatedAt());
        assertNotNull(workExperience.getUpdatedAt());
    }

    @Test
    void shouldCreateWorkExperience_whenEndDateIsNull() {
        LocalDate startDate = LocalDate.of(2024, 1, 1);

        WorkExperience workExperience = new WorkExperience(
                null,
                "Senior Software Engineer",
                "Liderança técnica",
                "Global Tech",
                startDate,
                null,
                null,
                null
        );

        assertNull(workExperience.getEndDate());
    }

    @Test
    void shouldThrowDomainValidationException_whenRoleIsBlank() {
        LocalDate startDate = LocalDate.of(2022, 1, 1);

        DomainValidationException exception = assertThrows(
                DomainValidationException.class,
                () -> new WorkExperience(
                        null,
                        "   ",
                        "Descrição",
                        "Empresa",
                        startDate,
                        null,
                        LocalDateTime.now(),
                        LocalDateTime.now()
                )
        );

        assertTrue(exception.getMessage().contains("role is required"));
    }

    @Test
    void shouldThrowDomainValidationException_whenDescriptionIsBlank() {
        LocalDate startDate = LocalDate.of(2022, 1, 1);

        DomainValidationException exception = assertThrows(
                DomainValidationException.class,
                () -> new WorkExperience(
                        null,
                        "Cargo",
                        "   ",
                        "Empresa",
                        startDate,
                        null,
                        LocalDateTime.now(),
                        LocalDateTime.now()
                )
        );

        assertTrue(exception.getMessage().contains("description is required"));
    }

    @Test
    void shouldThrowDomainValidationException_whenCompanyIsBlank() {
        LocalDate startDate = LocalDate.of(2022, 1, 1);

        DomainValidationException exception = assertThrows(
                DomainValidationException.class,
                () -> new WorkExperience(
                        null,
                        "Cargo",
                        "Descrição",
                        "   ",
                        startDate,
                        null,
                        LocalDateTime.now(),
                        LocalDateTime.now()
                )
        );

        assertTrue(exception.getMessage().contains("company is required"));
    }

    @Test
    void shouldThrowDomainValidationException_whenStartDateIsNull() {
        DomainValidationException exception = assertThrows(
                DomainValidationException.class,
                () -> new WorkExperience(
                        null,
                        "Cargo",
                        "Descrição",
                        "Empresa",
                        null,
                        null,
                        LocalDateTime.now(),
                        LocalDateTime.now()
                )
        );

        assertTrue(exception.getMessage().contains("start date is required"));
    }

    @Test
    void shouldThrowDomainValidationException_whenEndDateIsBeforeStartDate() {
        LocalDate startDate = LocalDate.of(2023, 1, 1);
        LocalDate endDate = LocalDate.of(2022, 1, 1);

        DomainValidationException exception = assertThrows(
                DomainValidationException.class,
                () -> new WorkExperience(
                        null,
                        "Cargo",
                        "Descrição",
                        "Empresa",
                        startDate,
                        endDate,
                        LocalDateTime.now(),
                        LocalDateTime.now()
                )
        );

        assertTrue(exception.getMessage().contains("end date cannot be before start date"));
    }
}

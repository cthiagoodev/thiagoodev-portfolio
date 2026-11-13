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

class EducationTest {

    @Test
    void shouldCreateEducation_whenAllFieldsAreValid() {
        UUID uuid = UUID.randomUUID();
        LocalDate startDate = LocalDate.of(2018, 2, 1);
        LocalDate endDate = LocalDate.of(2022, 12, 1);

        Education education = new Education(
                uuid,
                "Ciência da Computação",
                "Bacharelado em Ciência da Computação",
                "Universidade Federal",
                startDate,
                endDate,
                null,
                null
        );

        assertEquals(uuid, education.getUuid());
        assertEquals("Ciência da Computação", education.getCourse());
        assertEquals("Bacharelado em Ciência da Computação", education.getDescription());
        assertEquals("Universidade Federal", education.getEducationalInstitution());
        assertEquals(startDate, education.getStartDate());
        assertEquals(endDate, education.getEndDate());
        assertNotNull(education.getCreatedAt());
        assertNotNull(education.getUpdatedAt());
    }

    @Test
    void shouldCreateEducation_whenEndDateIsNull() {
        LocalDate startDate = LocalDate.of(2023, 2, 1);

        Education education = new Education(
                null,
                "Pós-Graduação em Arquitetura de Software",
                "Especialização em sistemas distribuídos",
                "Instituto de Tecnologia",
                startDate,
                null,
                null,
                null
        );

        assertNull(education.getEndDate());
    }

    @Test
    void shouldThrowDomainValidationException_whenCourseIsBlank() {
        LocalDate startDate = LocalDate.of(2020, 1, 1);

        DomainValidationException exception = assertThrows(
                DomainValidationException.class,
                () -> new Education(
                        null,
                        "   ",
                        "Descrição",
                        "Instituição",
                        startDate,
                        null,
                        LocalDateTime.now(),
                        LocalDateTime.now()
                )
        );

        assertTrue(exception.getMessage().contains("course is required"));
    }

    @Test
    void shouldThrowDomainValidationException_whenDescriptionIsBlank() {
        LocalDate startDate = LocalDate.of(2020, 1, 1);

        DomainValidationException exception = assertThrows(
                DomainValidationException.class,
                () -> new Education(
                        null,
                        "Curso",
                        "   ",
                        "Instituição",
                        startDate,
                        null,
                        LocalDateTime.now(),
                        LocalDateTime.now()
                )
        );

        assertTrue(exception.getMessage().contains("description is required"));
    }

    @Test
    void shouldThrowDomainValidationException_whenEducationalInstitutionIsBlank() {
        LocalDate startDate = LocalDate.of(2020, 1, 1);

        DomainValidationException exception = assertThrows(
                DomainValidationException.class,
                () -> new Education(
                        null,
                        "Curso",
                        "Descrição",
                        "   ",
                        startDate,
                        null,
                        LocalDateTime.now(),
                        LocalDateTime.now()
                )
        );

        assertTrue(exception.getMessage().contains("educational institution is required"));
    }

    @Test
    void shouldThrowDomainValidationException_whenStartDateIsNull() {
        DomainValidationException exception = assertThrows(
                DomainValidationException.class,
                () -> new Education(
                        null,
                        "Curso",
                        "Descrição",
                        "Instituição",
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
                () -> new Education(
                        null,
                        "Curso",
                        "Descrição",
                        "Instituição",
                        startDate,
                        endDate,
                        LocalDateTime.now(),
                        LocalDateTime.now()
                )
        );

        assertTrue(exception.getMessage().contains("end date cannot be before start date"));
    }
}

package br.com.thiagoodev.api.domain.entities;

import br.com.thiagoodev.api.domain.exceptions.DomainValidationException;
import com.fasterxml.jackson.annotation.JsonProperty;
import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.Table;
import lombok.AccessLevel;
import lombok.Getter;
import lombok.NoArgsConstructor;

import java.time.LocalDate;
import java.time.LocalDateTime;
import java.util.UUID;

import static br.com.thiagoodev.api.domain.utils.Sanitizer.sanitize;

@Entity
@Table(name = "work_experience")
@Getter
@NoArgsConstructor(access = AccessLevel.PROTECTED)
public class WorkExperience {

    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    @Column(name = "uuid", updatable = false, nullable = false)
    private UUID uuid;

    @Column(name = "role", nullable = false, length = 255)
    private String role;

    @Column(name = "description", nullable = false, columnDefinition = "TEXT")
    private String description;

    @Column(name = "company", nullable = false, length = 255)
    private String company;

    @JsonProperty("start_date")
    @Column(name = "start_date", nullable = false)
    private LocalDate startDate;

    @JsonProperty("end_date")
    @Column(name = "end_date")
    private LocalDate endDate;

    @JsonProperty("created_at")
    @Column(name = "created_at", nullable = false, updatable = false)
    private LocalDateTime createdAt;

    @JsonProperty("updated_at")
    @Column(name = "updated_at", nullable = false)
    private LocalDateTime updatedAt;

    public WorkExperience(
            UUID uuid,
            String role,
            String description,
            String company,
            LocalDate startDate,
            LocalDate endDate,
            LocalDateTime createdAt,
            LocalDateTime updatedAt
    ) {
        this.uuid = uuid;
        this.role = sanitize(role);
        this.description = sanitize(description);
        this.company = sanitize(company);
        this.startDate = startDate;
        this.endDate = endDate;
        this.createdAt = createdAt != null ? createdAt : LocalDateTime.now();
        this.updatedAt = updatedAt != null ? updatedAt : LocalDateTime.now();

        this.validate();
    }

    public void validate() {
        if (this.role.isEmpty()) {
            throw new DomainValidationException("domain entity validation error: role is required");
        }

        if (this.description.isEmpty()) {
            throw new DomainValidationException("domain entity validation error: description is required");
        }

        if (this.company.isEmpty()) {
            throw new DomainValidationException("domain entity validation error: company is required");
        }

        if (this.startDate == null) {
            throw new DomainValidationException("domain entity validation error: start date is required");
        }

        if (this.endDate != null && this.endDate.isBefore(this.startDate)) {
            throw new DomainValidationException("domain entity validation error: end date cannot be before start date");
        }
    }
}

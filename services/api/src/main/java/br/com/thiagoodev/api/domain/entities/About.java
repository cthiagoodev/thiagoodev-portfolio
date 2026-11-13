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

import java.time.LocalDateTime;
import java.util.UUID;

import static br.com.thiagoodev.api.domain.utils.Sanitizer.sanitize;

@Entity
@Table(name = "about")
@Getter
@NoArgsConstructor(access = AccessLevel.PROTECTED)
public class About {

    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    @Column(name = "uuid", updatable = false, nullable = false)
    private UUID uuid;

    @Column(name = "name", nullable = false, length = 255)
    private String name;

    @Column(name = "phone", length = 20)
    private String phone;

    @Column(name = "description", length = 255)
    private String description;

    @Column(name = "text", columnDefinition = "TEXT")
    private String text;

    @JsonProperty("created_at")
    @Column(name = "created_at", nullable = false, updatable = false)
    private LocalDateTime createdAt;

    @JsonProperty("updated_at")
    @Column(name = "updated_at", nullable = false)
    private LocalDateTime updatedAt;

    public About(
            UUID uuid,
            String name,
            String phone,
            String description,
            String text,
            LocalDateTime createdAt,
            LocalDateTime updatedAt
    ) {
        this.uuid = uuid;
        this.name = sanitize(name);
        this.phone = sanitize(phone);
        this.description = sanitize(description);
        this.text = sanitize(text);
        this.createdAt = createdAt != null ? createdAt : LocalDateTime.now();
        this.updatedAt = updatedAt != null ? updatedAt : LocalDateTime.now();

        this.validate();
    }

    public void validate() {
        if (this.name.isEmpty()) {
            throw new DomainValidationException("domain entity validation error: name is required");
        }

        if (this.phone.length() > 20) {
            throw new DomainValidationException("domain entity validation error: phone must have at most 20 characters");
        }
    }
}

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
@Table(name = "skills")
@Getter
@NoArgsConstructor(access = AccessLevel.PROTECTED)
public class Skill {

    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    @Column(name = "uuid", updatable = false, nullable = false)
    private UUID uuid;

    @Column(name = "label", nullable = false, unique = true, length = 255)
    private String label;

    @Column(name = "url", columnDefinition = "TEXT")
    private String url;

    @JsonProperty("image_path")
    @Column(name = "image_path", columnDefinition = "TEXT")
    private String imagePath;

    @JsonProperty("created_at")
    @Column(name = "created_at", nullable = false, updatable = false)
    private LocalDateTime createdAt;

    @JsonProperty("updated_at")
    @Column(name = "updated_at", nullable = false)
    private LocalDateTime updatedAt;

    public Skill(
            UUID uuid,
            String label,
            String url,
            String imagePath,
            LocalDateTime createdAt,
            LocalDateTime updatedAt
    ) {
        this.uuid = uuid;
        this.label = sanitize(label);
        this.url = sanitize(url);
        this.imagePath = sanitize(imagePath);
        this.createdAt = createdAt != null ? createdAt : LocalDateTime.now();
        this.updatedAt = updatedAt != null ? updatedAt : LocalDateTime.now();

        this.validate();
    }

    public void validate() {
        if (this.label.isEmpty()) {
            throw new DomainValidationException("domain entity validation error: label is required");
        }
    }
}

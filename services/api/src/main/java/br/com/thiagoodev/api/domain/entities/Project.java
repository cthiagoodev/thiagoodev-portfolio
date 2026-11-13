package br.com.thiagoodev.api.domain.entities;

import br.com.thiagoodev.api.domain.exceptions.DomainValidationException;
import com.fasterxml.jackson.annotation.JsonProperty;
import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.FetchType;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.JoinColumn;
import jakarta.persistence.JoinTable;
import jakarta.persistence.ManyToMany;
import jakarta.persistence.Table;
import lombok.AccessLevel;
import lombok.Getter;
import lombok.NoArgsConstructor;
import org.hibernate.annotations.JdbcTypeCode;
import org.hibernate.type.SqlTypes;

import java.time.LocalDateTime;
import java.util.ArrayList;
import java.util.List;
import java.util.UUID;

import static br.com.thiagoodev.api.domain.utils.Sanitizer.sanitize;

@Entity
@Table(name = "projects")
@Getter
@NoArgsConstructor(access = AccessLevel.PROTECTED)
public class Project {

    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    @Column(name = "uuid", updatable = false, nullable = false)
    private UUID uuid;

    @JsonProperty("external_id")
    @Column(name = "external_id")
    private Long externalId;

    @Column(name = "name", nullable = false, length = 255)
    private String name;

    @Column(name = "description", columnDefinition = "TEXT")
    private String description;

    @Column(name = "url", columnDefinition = "TEXT")
    private String url;

    @JsonProperty("repo_url")
    @Column(name = "repo_url", columnDefinition = "TEXT")
    private String repoUrl;

    @JdbcTypeCode(SqlTypes.ARRAY)
    @Column(name = "languages", columnDefinition = "varchar(255)[]")
    private List<String> languages;

    @ManyToMany(fetch = FetchType.EAGER)
    @JoinTable(
            name = "projects_skills",
            joinColumns = @JoinColumn(name = "project_id"),
            inverseJoinColumns = @JoinColumn(name = "skill_id")
    )
    private List<Skill> skills;

    @JsonProperty("pushed_at")
    @Column(name = "pushed_at")
    private LocalDateTime pushedAt;

    @JsonProperty("created_at")
    @Column(name = "created_at", nullable = false, updatable = false)
    private LocalDateTime createdAt;

    @JsonProperty("updated_at")
    @Column(name = "updated_at", nullable = false)
    private LocalDateTime updatedAt;

    public Project(
            UUID uuid,
            Long externalId,
            String name,
            String description,
            String url,
            String repoUrl,
            List<String> languages,
            List<Skill> skills,
            LocalDateTime pushedAt,
            LocalDateTime createdAt,
            LocalDateTime updatedAt
    ) {
        this.uuid = uuid;
        this.externalId = externalId;
        this.name = sanitize(name);
        this.description = sanitize(description);
        this.url = sanitize(url);
        this.repoUrl = sanitize(repoUrl);
        this.languages = languages != null ? new ArrayList<>(languages) : new ArrayList<>();
        this.skills = skills != null ? new ArrayList<>(skills) : new ArrayList<>();
        this.pushedAt = pushedAt;
        this.createdAt = createdAt != null ? createdAt : LocalDateTime.now();
        this.updatedAt = updatedAt != null ? updatedAt : LocalDateTime.now();

        this.validate();
    }

    public void syncFromGitHub(
            String name,
            String description,
            String url,
            String repoUrl,
            List<String> languages,
            LocalDateTime pushedAt,
            LocalDateTime updatedAt
    ) {
        this.name = sanitize(name);
        this.description = sanitize(description);
        this.url = sanitize(url);
        this.repoUrl = sanitize(repoUrl);
        this.languages = languages != null ? new ArrayList<>(languages) : new ArrayList<>();
        this.pushedAt = pushedAt;
        this.updatedAt = updatedAt != null ? updatedAt : LocalDateTime.now();

        this.validate();
    }

    public void validate() {
        if (this.name.isEmpty()) {
            throw new DomainValidationException("domain entity validation error: project name is required");
        }
    }

    public void addSkill(Skill skill) {
        if (skill == null) {
            return;
        }

        boolean exists = this.skills.stream().anyMatch(existingSkill ->
                (skill.getUuid() != null && skill.getUuid().equals(existingSkill.getUuid())) ||
                (!skill.getLabel().isEmpty() && skill.getLabel().equalsIgnoreCase(existingSkill.getLabel()))
        );

        if (!exists) {
            this.skills.add(skill);
        }
    }

    public void removeSkill(UUID skillUuid) {
        if (skillUuid == null) {
            return;
        }

        this.skills.removeIf(skill -> skillUuid.equals(skill.getUuid()));
    }
}

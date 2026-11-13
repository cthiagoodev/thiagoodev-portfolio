package br.com.thiagoodev.api.domain.repositories;

import br.com.thiagoodev.api.domain.entities.Skill;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.UUID;

@Repository
public interface SkillRepository extends JpaRepository<Skill, UUID> {
}

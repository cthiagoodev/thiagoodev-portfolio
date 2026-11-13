package br.com.thiagoodev.api.domain.repositories;

import br.com.thiagoodev.api.domain.entities.Education;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.UUID;

@Repository
public interface EducationRepository extends JpaRepository<Education, UUID> {
}

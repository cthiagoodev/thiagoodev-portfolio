package br.com.thiagoodev.api.domain.repositories;

import br.com.thiagoodev.api.domain.entities.Project;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.UUID;

@Repository
public interface ProjectRepository extends JpaRepository<Project, UUID> {
}

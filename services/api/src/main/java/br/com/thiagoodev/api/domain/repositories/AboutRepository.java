package br.com.thiagoodev.api.domain.repositories;

import br.com.thiagoodev.api.domain.entities.About;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.Optional;
import java.util.UUID;

@Repository
public interface AboutRepository extends JpaRepository<About, UUID> {

    Optional<About> findFirstByOrderByCreatedAtAsc();
}

package echochat.userservice.repository;

import echochat.userservice.entity.User;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;

import java.util.UUID;

public interface UserRepository extends JpaRepository<User, UUID>, JpaSpecificationExecutor<User> {
    @Query(value = """
        SELECT * FROM users
        WHERE is_public = true
        AND similarity(name, :query) > 0.3
        ORDER BY similarity(name, :query) DESC
        """,
            countQuery = """
        SELECT count(*) FROM users
        WHERE is_public = true
        AND similarity(name, :query) > 0.3
        """,
            nativeQuery = true)
    Page<User> searchByNameSimilar(@Param("query") String query, Pageable pageable);
}

package echochat.userservice.entity;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.Id;
import lombok.Builder;
import lombok.Getter;
import lombok.RequiredArgsConstructor;
import lombok.Setter;

import java.time.LocalDateTime;
import java.util.UUID;

@Entity
@Builder
@RequiredArgsConstructor
@Getter
@Setter
public class User {
    @Id
    private UUID id;

    @Column(nullable = false, name = "display_name")
    private String displayName;

    private String bio;

    @Column(name="is_public")
    private boolean isPublic;

    @Column(nullable = false, name = "created_at")
    private LocalDateTime createdAt;

    @Column(nullable = false, name = "updated_at")
    private LocalDateTime updatedAt;
}

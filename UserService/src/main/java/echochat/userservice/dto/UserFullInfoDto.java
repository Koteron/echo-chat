package echochat.userservice.dto;

import lombok.NonNull;

import java.time.LocalDateTime;

public record UserFullInfoDto(
        @NonNull
        String displayName,

        String bio,

        @NonNull
        LocalDateTime createdAt,

        @NonNull
        LocalDateTime updatedAt
) {
}

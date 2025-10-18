package echochat.userservice.dto;

import lombok.NonNull;

import java.util.UUID;

public record UserUpdateDto(
        @NonNull
        UUID id,

        String displayName,

        String bio
) {
}

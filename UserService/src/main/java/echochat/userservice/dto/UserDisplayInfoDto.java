package echochat.userservice.dto;

import lombok.NonNull;

public record UserDisplayInfoDto(
        @NonNull
        String displayName
) {
}

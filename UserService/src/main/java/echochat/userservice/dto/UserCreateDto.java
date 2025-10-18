package echochat.userservice.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.NonNull;

import java.util.UUID;

public record UserCreateDto(
        @NonNull
        @JsonProperty(value = "keycloak_id")
        UUID keycloakId,

        @NonNull
        @JsonProperty(value = "display_name")
        String displayName
) {
}

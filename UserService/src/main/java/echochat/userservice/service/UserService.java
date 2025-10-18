package echochat.userservice.service;

import echochat.userservice.dto.UserCreateDto;
import echochat.userservice.dto.UserDisplayInfoDto;
import echochat.userservice.dto.UserFullInfoDto;
import echochat.userservice.dto.UserUpdateDto;

import java.util.UUID;

public interface UserService {
    void delete(UUID id);
    void create(UserCreateDto userCreateDto);
    UserFullInfoDto update(UserUpdateDto userUpdateDto);
    UserFullInfoDto getFullInfo(UUID id);
    UserDisplayInfoDto getDisplayInfo(UUID id);
    String getDisplayName(UUID id);
}

package echochat.userservice.service;

import echochat.userservice.dto.UserCreateDto;
import echochat.userservice.dto.UserDisplayInfoDto;
import echochat.userservice.dto.UserFullInfoDto;
import echochat.userservice.dto.UserUpdateDto;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;

import java.util.List;
import java.util.UUID;

public interface UserService {
    void delete(UUID id);
    void create(UserCreateDto userCreateDto);
    UserFullInfoDto update(UserUpdateDto userUpdateDto);
    UserFullInfoDto getFullInfo(UUID id);
    UserDisplayInfoDto getDisplayInfo(UUID id);
    List<String> getDisplayNames(List<UUID> ids);
    Page<UserDisplayInfoDto> searchDisplayInfos(String nameSearchString, Pageable pageable);
}

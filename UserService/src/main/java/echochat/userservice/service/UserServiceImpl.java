package echochat.userservice.service;

import echochat.userservice.dto.UserCreateDto;
import echochat.userservice.dto.UserDisplayInfoDto;
import echochat.userservice.dto.UserFullInfoDto;
import echochat.userservice.dto.UserUpdateDto;
import echochat.userservice.entity.User;
import echochat.userservice.exception.NotFoundException;
import echochat.userservice.mapper.UserMapper;
import echochat.userservice.repository.UserRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.domain.Specification;
import org.springframework.stereotype.Service;

import java.time.LocalDateTime;
import java.util.List;
import java.util.UUID;

@Service
@RequiredArgsConstructor
public class UserServiceImpl implements UserService {
    private final UserRepository userRepository;
    private final UserMapper userMapper;

    @Override
    public void delete(UUID id) {
        userRepository.deleteById(id);
    }

    @Override
    public void create(UserCreateDto userCreateDto) {
        userRepository.save(
                User.builder()
                .id(userCreateDto.keycloakId())
                .displayName(userCreateDto.displayName())
                .createdAt(LocalDateTime.now())
                .updatedAt(LocalDateTime.now())
                .build()
        );
    }

    @Override
    public UserFullInfoDto update(UserUpdateDto userUpdateDto) {
        User newUser = User.builder()
                .bio(userUpdateDto.bio())
                .displayName(userUpdateDto.displayName())
                .updatedAt(LocalDateTime.now())
                .build();
        return userMapper.toUserFullInfoDto(userRepository.save(newUser));
    }

    @Override
    public UserFullInfoDto getFullInfo(UUID id) {
        return userMapper.toUserFullInfoDto(userRepository.findById(id)
                .orElseThrow(() -> new NotFoundException("User not found!")));
    }

    @Override
    public UserDisplayInfoDto getDisplayInfo(UUID id) {
        return userMapper.toUserDisplayInfoDto(userRepository.findById(id)
                .orElseThrow(() -> new NotFoundException("User not found!")));
    }

    @Override
    public List<String> getDisplayNames(List<UUID> ids) {
        return userRepository.findAllById(ids).stream().map(User::getDisplayName).toList();
    }

    @Override
    public Page<UserDisplayInfoDto> searchDisplayInfos(String nameSearchString, Pageable pageable) {
        return userRepository.searchByNameSimilar(nameSearchString, pageable)
                .map(userMapper::toUserDisplayInfoDto);
    }
}

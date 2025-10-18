package echochat.userservice.controller;

import echochat.userservice.dto.UserCreateDto;
import echochat.userservice.dto.UserDisplayInfoDto;
import echochat.userservice.dto.UserFullInfoDto;
import echochat.userservice.dto.UserUpdateDto;
import echochat.userservice.service.UserService;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.*;

import java.util.UUID;

@RestController
@RequiredArgsConstructor
@RequestMapping("/user")
public class UserController {
    private final UserService userService;

    @GetMapping("/{id}")
    public UserDisplayInfoDto getDisplayInfo(@PathVariable UUID id) {
        return userService.getDisplayInfo(id);
    }

    @GetMapping("/full/{id}")
    public UserFullInfoDto getFullInfo(@PathVariable UUID id) {
        return userService.getFullInfo(id);
    }

    @DeleteMapping("/{id}")
    public void delete(@PathVariable UUID id) {
        userService.delete(id);
    }

    @PostMapping("/create")
    public void create(@RequestBody UserCreateDto userCreateDto) {
        userService.create(userCreateDto);
    }

    @PatchMapping("/update")
    public UserFullInfoDto update(@RequestBody UserUpdateDto userUpdateDto) {
        return userService.update(userUpdateDto);
    }
}

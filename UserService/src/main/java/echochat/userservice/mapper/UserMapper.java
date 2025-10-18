package echochat.userservice.mapper;

import echochat.userservice.dto.UserDisplayInfoDto;
import echochat.userservice.dto.UserFullInfoDto;
import echochat.userservice.entity.User;
import org.mapstruct.Mapper;

@Mapper(componentModel = "spring")
public interface UserMapper {
    UserFullInfoDto toUserFullInfoDto(User user);
    UserDisplayInfoDto toUserDisplayInfoDto(User user);
}

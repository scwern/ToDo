package user

import "ToDo/internal/domain/user"

func ToUser(dto CreateUserDTO) user.User {
	return user.NewUser(dto.Name, dto.Email, dto.Password)
}

func ToUserDTO(u user.User) DTO {
	return DTO{
		ID:    u.ID().String(),
		Name:  u.Name(),
		Email: u.Email(),
	}
}

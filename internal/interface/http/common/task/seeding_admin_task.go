package task

import (
	"context"
	"userman/internal/domain/user"
)

func SeedingAdminTask(ctx context.Context, userRepo user.UserRepository) {
	user := user.NewUser("admin", "admin@admin.com", "Adm1nUserman")
	user.SetRole("admin")

	if exits, _ := userRepo.GetByEmail(ctx, user.Email); exits != nil {
		return
	}

	userRepo.Create(ctx, user)
}

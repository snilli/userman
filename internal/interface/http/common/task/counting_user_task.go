package task

import (
	"context"
	"log"
	"time"
	"userman/internal/domain/user"
)

func CountingUserTask(ctx context.Context, userRepo user.UserRepository) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("User count logger stopped")
			return
		case <-ticker.C:
			count, err := userRepo.Count(ctx)
			if err != nil {
				log.Printf("Error counting users : %v", err)
				continue
			}
			log.Printf("Current user in the system : %d", count)
		}
	}
}

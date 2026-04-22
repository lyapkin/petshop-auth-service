package account

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
)

func (uc *Usecase) Delete(ctx context.Context, id uuid.UUID) error {
	uc.log.InfoContext(ctx, "account deletion started")

	if err := uc.accountRepo.Delete(ctx, id); err != nil {
		uc.log.ErrorContext(ctx, "account deletion failed", slog.String("err", err.Error()))
		return err
	}

	uc.log.InfoContext(ctx, "account deletion succeeded")
	return nil
}

package permission

import (
	"context"
	"log/slog"
)

func (u *Usecase) Delete(ctx context.Context, id int) error {
	u.log.InfoContext(ctx, "permission deletion started")

	if err := u.permissionRepo.Delete(ctx, id); err != nil {
		u.log.ErrorContext(ctx, "permission db delete failed", slog.String("err", err.Error()))
		return err
	}

	u.log.InfoContext(ctx, "permission deletion succeeded")
	return nil
}

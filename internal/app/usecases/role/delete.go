package role

import (
	"context"
	"log/slog"
)

func (u *Usecase) Delete(ctx context.Context, id int) error {
	u.log.InfoContext(ctx, "role deletion started")

	// TODO: cancel delete if role is base

	if err := u.roleRepo.Delete(ctx, id); err != nil {
		u.log.ErrorContext(ctx, "role deletion failed", slog.String("err", err.Error()))
		return err
	}

	u.log.InfoContext(ctx, "role deletion succeeded")
	return nil
}

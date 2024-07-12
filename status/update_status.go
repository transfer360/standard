package status

import (
	"context"
	"github.com/transfer360/sys360/sys_updates"
)

func SetStatus(ctx context.Context, SRef string, statusCode int, source string) error {

	update := sys_updates.NoticeStatusChange{
		Sref:       SRef,
		StatusCode: statusCode,
		Source:     source,
	}

	return update.Update(ctx)

}

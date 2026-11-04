package whatsapp

import (
	"context"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store"
)

type clientVersion struct {
	Major int
	Minor int
	Patch int
}

var version clientVersion

// RefreshWebVersion must run before clients start: whatsmeow stores the version globally.
func RefreshWebVersion(ctx context.Context) error {
	latest, err := whatsmeow.GetLatestVersion(ctx, nil)
	if err != nil {
		return err
	}
	store.SetWAVersion(*latest)
	return nil
}

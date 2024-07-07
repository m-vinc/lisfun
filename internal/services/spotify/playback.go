package spotify

import (
	"context"
	"lisfun/internal/services/models"
	"log"
)

func (s *Service) PlaybackState(ctx context.Context, params *models.SpotifyPlaybackStateParams) (*models.SpotifyPlaybackStateResult, error) {
	client, err := s.Client(ctx, params.User)
	if err != nil {
		return nil, err
	}

	playback, err := client.PlayerCurrentlyPlaying(ctx)

	log.Printf("playing: %+v; err: %+v", playback, err)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

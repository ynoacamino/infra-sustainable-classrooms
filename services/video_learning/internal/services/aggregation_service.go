package services

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	videolearningdb "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/gen/database"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/internal/repositories"
	goaLog "goa.design/clue/log"
)

// AggregationService handles periodic aggregation of cached data to database
type AggregationService struct {
	repoManager *repositories.RepositoryManager
	ctx         context.Context
	ticker      *time.Ticker
	stopChan    chan bool

	// Configuration
	batchSize  int
	maxRetries int
}

// AggregationServiceConfig holds configuration for the aggregation service
type AggregationServiceConfig struct {
	BatchSize  int // Maximum number of keys to process in one batch
	MaxRetries int // Maximum number of retries for failed operations
}

// NewAggregationService creates a new aggregation service with default configuration
func NewAggregationService(ctx context.Context, repoManager *repositories.RepositoryManager) *AggregationService {
	return NewAggregationServiceWithConfig(ctx, repoManager, AggregationServiceConfig{
		BatchSize:  100,
		MaxRetries: 3,
	})
}

// NewAggregationServiceWithConfig creates a new aggregation service with custom configuration
func NewAggregationServiceWithConfig(ctx context.Context, repoManager *repositories.RepositoryManager, config AggregationServiceConfig) *AggregationService {
	return &AggregationService{
		repoManager: repoManager,
		ctx:         ctx,
		stopChan:    make(chan bool),
		batchSize:   config.BatchSize,
		maxRetries:  config.MaxRetries,
	}
}

// Start begins the periodic aggregation process
func (as *AggregationService) Start(interval time.Duration) {
	as.ticker = time.NewTicker(interval)

	goaLog.Printf(as.ctx, "Starting aggregation service with interval: %v", interval)

	go func() {
		defer as.ticker.Stop()

		for {
			select {
			case <-as.ticker.C:
				if err := as.aggregateAndFlush(); err != nil {
					goaLog.Printf(as.ctx, "Error during aggregation: %v", err)
				}
			case <-as.stopChan:
				goaLog.Printf(as.ctx, "Stopping aggregation service")
				return
			}
		}
	}()
}

// Stop stops the aggregation service
func (as *AggregationService) Stop() {
	if as.ticker != nil {
		as.stopChan <- true
	}
}

// aggregateAndFlush aggregates cached data and flushes it to the database
func (as *AggregationService) aggregateAndFlush() error {
	goaLog.Printf(as.ctx, "Starting aggregation cycle")

	// Aggregate video views
	if err := as.aggregateVideoViews(); err != nil {
		goaLog.Printf(as.ctx, "Error aggregating video views: %v", err)
	}

	// Aggregate video likes
	if err := as.aggregateVideoLikes(); err != nil {
		goaLog.Printf(as.ctx, "Error aggregating video likes: %v", err)
	}

	// Aggregate user category likes
	if err := as.aggregateUserCategoryLikes(); err != nil {
		goaLog.Printf(as.ctx, "Error aggregating user category likes: %v", err)
	}

	// Aggregate user video likes
	if err := as.aggregateUserVideoLikes(); err != nil {
		goaLog.Printf(as.ctx, "Error aggregating user video likes: %v", err)
	}

	goaLog.Printf(as.ctx, "Aggregation cycle completed")
	return nil
}

// processBatch processes a batch of keys and values within a transaction
func (as *AggregationService) processBatch(keys []string, values []string, processor func(pgx.Tx, []string, []string) error) error {
	return as.repoManager.TxManager.WithTransaction(as.ctx, func(tx pgx.Tx) error {
		return processor(tx, keys, values)
	})
}

// aggregateVideoViews processes cached view counts and updates the database
func (as *AggregationService) aggregateVideoViews() error {
	// Find all video view keys
	keys, err := as.repoManager.CacheRepo.Scan(as.ctx, "video:views:*")
	if err != nil {
		return fmt.Errorf("failed to scan video view keys: %w", err)
	}

	if len(keys) == 0 {
		return nil
	}

	// Get all values at once
	values, err := as.repoManager.CacheRepo.MGet(as.ctx, keys)
	if err != nil {
		return fmt.Errorf("failed to get video view values: %w", err)
	}

	// Process in batches for better performance
	totalProcessed := 0
	for i := 0; i < len(keys); i += as.batchSize {
		end := i + as.batchSize
		if end > len(keys) {
			end = len(keys)
		}

		batchKeys := keys[i:end]
		batchValues := values[i:end]

		err := as.processBatch(batchKeys, batchValues, func(tx pgx.Tx, bKeys []string, bValues []string) error {
			txRepos := as.repoManager.WithTransaction(tx)
			processed := 0

			// Process each video view update in this batch
			for j, key := range bKeys {
				if bValues[j] == "" {
					continue
				}

				// Extract video ID from key (video:views:123)
				parts := strings.Split(key, ":")
				if len(parts) != 3 {
					continue
				}

				videoID, err := strconv.ParseInt(parts[2], 10, 64)
				if err != nil {
					continue
				}

				cachedViews, err := strconv.ParseInt(bValues[j], 10, 32)
				if err != nil || cachedViews <= 0 {
					continue
				}

				// Update video views in database
				err = txRepos.VideoRepo.IncrementVideoViews(as.ctx, videolearningdb.IncrementVideoViewsParams{
					ID:    videoID,
					Views: int32(cachedViews),
				})
				if err != nil {
					goaLog.Printf(as.ctx, "Failed to update views for video %d: %v", videoID, err)
					continue
				}

				// Delete the cached entry after successful update
				as.repoManager.CacheRepo.Delete(as.ctx, key)
				processed++
			}

			goaLog.Printf(as.ctx, "Processed %d video view updates in batch", processed)
			return nil
		})

		if err != nil {
			goaLog.Printf(as.ctx, "Failed to process video views batch: %v", err)
			continue
		}

		totalProcessed += end - i
	}

	goaLog.Printf(as.ctx, "Updated views for %d videos total", totalProcessed)
	return nil
}

// aggregateVideoLikes processes cached like counts and updates the database
func (as *AggregationService) aggregateVideoLikes() error {
	// Find all video like keys
	keys, err := as.repoManager.CacheRepo.Scan(as.ctx, "video:likes:*")
	if err != nil {
		return fmt.Errorf("failed to scan video like keys: %w", err)
	}

	if len(keys) == 0 {
		return nil
	}

	// Get all values at once
	values, err := as.repoManager.CacheRepo.MGet(as.ctx, keys)
	if err != nil {
		return fmt.Errorf("failed to get video like values: %w", err)
	}

	// Use transaction manager
	return as.repoManager.TxManager.WithTransaction(as.ctx, func(tx pgx.Tx) error {
		txRepos := as.repoManager.WithTransaction(tx)

		// Process each video like update
		for i, key := range keys {
			if values[i] == "" {
				continue
			}

			// Extract video ID from key (video:likes:123)
			parts := strings.Split(key, ":")
			if len(parts) != 3 {
				continue
			}

			videoID, err := strconv.ParseInt(parts[2], 10, 64)
			if err != nil {
				continue
			}

			cachedLikes, err := strconv.ParseInt(values[i], 10, 32)
			if err != nil {
				continue
			}

			// Update video likes in database (can be negative for unlikes)
			err = txRepos.VideoRepo.IncrementVideoLikes(as.ctx, videolearningdb.IncrementVideoLikesParams{
				ID:    videoID,
				Likes: int32(cachedLikes),
			})
			if err != nil {
				goaLog.Printf(as.ctx, "Failed to update likes for video %d: %v", videoID, err)
				continue
			}

			// Delete the cached entry after successful update
			as.repoManager.CacheRepo.Delete(as.ctx, key)
		}

		goaLog.Printf(as.ctx, "Updated likes for %d videos", len(keys))
		return nil
	})
}

// aggregateUserCategoryLikes processes cached user category likes and updates the database
func (as *AggregationService) aggregateUserCategoryLikes() error {
	// Find all user category like keys
	keys, err := as.repoManager.CacheRepo.Scan(as.ctx, "user:category:likes:*")
	if err != nil {
		return fmt.Errorf("failed to scan user category like keys: %w", err)
	}

	if len(keys) == 0 {
		return nil
	}

	// Get all values at once
	values, err := as.repoManager.CacheRepo.MGet(as.ctx, keys)
	if err != nil {
		return fmt.Errorf("failed to get user category like values: %w", err)
	}

	// Use transaction manager
	return as.repoManager.TxManager.WithTransaction(as.ctx, func(tx pgx.Tx) error {
		txRepos := as.repoManager.WithTransaction(tx)

		// Process each user category like update
		for i, key := range keys {
			if values[i] == "" {
				continue
			}

			// Extract user ID and category ID from key (user:category:likes:123:456)
			parts := strings.Split(key, ":")
			if len(parts) != 5 {
				continue
			}

			userID, err := strconv.ParseInt(parts[3], 10, 64)
			if err != nil {
				continue
			}

			categoryID, err := strconv.ParseInt(parts[4], 10, 64)
			if err != nil {
				continue
			}

			cachedLikes, err := strconv.ParseInt(values[i], 10, 32)
			if err != nil {
				continue
			}

			// Update or create user category like in database using upsert method
			err = txRepos.UserCategoryLikeRepo.UpsertUserCategoryLike(as.ctx, videolearningdb.UpsertUserCategoryLikeParams{
				UserID:     userID,
				CategoryID: categoryID,
				Likes:      int32(cachedLikes),
			})
			if err != nil {
				goaLog.Printf(as.ctx, "Failed to update category likes for user %d, category %d: %v", userID, categoryID, err)
				continue
			}

			// Delete the cached entry after successful update
			as.repoManager.CacheRepo.Delete(as.ctx, key)
		}

		goaLog.Printf(as.ctx, "Updated category likes for %d user-category pairs", len(keys))
		return nil
	})
}

// aggregateUserVideoLikes processes cached user video likes and updates the database
func (as *AggregationService) aggregateUserVideoLikes() error {
	// Find all user video like keys
	keys, err := as.repoManager.CacheRepo.Scan(as.ctx, "user:like:*")
	if err != nil {
		return fmt.Errorf("failed to scan user video like keys: %w", err)
	}

	if len(keys) == 0 {
		return nil
	}

	// Get all values at once
	values, err := as.repoManager.CacheRepo.MGet(as.ctx, keys)
	if err != nil {
		return fmt.Errorf("failed to get user video like values: %w", err)
	}

	// Use transaction manager
	return as.repoManager.TxManager.WithTransaction(as.ctx, func(tx pgx.Tx) error {
		txRepos := as.repoManager.WithTransaction(tx)

		// Process each user video like update
		for i, key := range keys {
			if values[i] == "" {
				continue
			}

			// Extract user ID and video ID from key (user:like:123:456)
			parts := strings.Split(key, ":")
			if len(parts) != 4 {
				continue
			}

			userID, err := strconv.ParseInt(parts[2], 10, 64)
			if err != nil {
				continue
			}

			videoID, err := strconv.ParseInt(parts[3], 10, 64)
			if err != nil {
				continue
			}

			liked := values[i] == "1"

			// Update or create user video like in database
			err = txRepos.UserCategoryLikeRepo.UpsertUserVideoLike(as.ctx, videolearningdb.UpsertUserVideoLikeParams{
				UserID:  userID,
				VideoID: videoID,
				Liked:   liked,
			})
			if err != nil {
				goaLog.Printf(as.ctx, "Failed to update video like for user %d, video %d: %v", userID, videoID, err)
				continue
			}

			// Keep the cache entry for quick lookups but refresh expiration
			as.repoManager.CacheRepo.Set(as.ctx, key, values[i], 24*time.Hour)
		}

		goaLog.Printf(as.ctx, "Synced video likes to database for %d user-video pairs", len(keys))
		return nil
	})
}

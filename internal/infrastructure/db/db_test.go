package db

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/haru-256/blog-ddd-uber-fx/internal/domain/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserRepositoryImpl(t *testing.T) {
	// Initialize logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	repo := NewUserRepositoryImpl(logger)

	ctx := context.Background()

	t.Run("Create", func(t *testing.T) {
		tests := []struct {
			name    string
			input   *models.User
			wantErr bool
		}{
			{
				name:    "Success",
				input:   mustNewUser("test-user-1", []*models.Task{}),
				wantErr: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := repo.Create(ctx, tt.input)
				if tt.wantErr {
					assert.Error(t, err)
					return
				}
				require.NoError(t, err)
				assert.Equal(t, tt.input, got)

				// Verify it exists
				found, err := repo.GetById(ctx, tt.input.Id())
				require.NoError(t, err)
				assert.Equal(t, tt.input.Id(), found.Id())
			})
		}
	})

	t.Run("GetAll", func(t *testing.T) {
		tests := []struct {
			name    string
			wantErr bool
		}{
			{
				name:    "Success",
				wantErr: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := repo.GetAll(ctx)
				if tt.wantErr {
					assert.Error(t, err)
					return
				}
				require.NoError(t, err)
				assert.NotNil(t, got)
				// Length should be at least what we initialized/added.
				// Since we added one in previous test and init has 2, expecting >= 3.
				// But tests order is not guaranteed if we don't control it.
				// However, these sub-tests run sequentially.
			})
		}
	})

	t.Run("GetById", func(t *testing.T) {
		// Get an existing user ID from init data
		var existingUserId *models.UserId
		if len(usersDB) > 0 {
			var err error
			existingUserId, err = models.BuildUserId(usersDB[0].id)
			require.NoError(t, err)
		}

		tests := []struct {
			name    string
			id      *models.UserId
			wantErr bool
		}{
			{
				name:    "Success - Existing User (from init data)",
				id:      existingUserId,
				wantErr: false,
			},
			{
				name:    "NotFound",
				id:      models.NewUserId(),
				wantErr: true,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				require.NotNil(t, tt.id)
				got, err := repo.GetById(ctx, tt.id)
				if tt.wantErr {
					assert.Error(t, err)
					return
				}
				require.NoError(t, err)
				assert.Equal(t, tt.id, got.Id())
			})
		}
	})

	t.Run("Update", func(t *testing.T) {
		// Prepare a user to update
		userToUpdate := mustNewUser("original-name", []*models.Task{})
		_, _ = repo.Create(ctx, userToUpdate)

		tests := []struct {
			name    string
			input   *models.User
			wantErr bool
		}{
			{
				name: "Success",
				input: func() *models.User {
					// We need to construct a user with the SAME ID as userToUpdate, but new name
					u, err := models.BuildUser(userToUpdate.Id(), "updated-name", []*models.Task{})
					if err != nil {
						panic(err)
					}
					return u
				}(),
				wantErr: false,
			},
			{
				name:    "NotFound",
				input:   mustNewUser("name", []*models.Task{}),
				wantErr: true,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := repo.Update(ctx, tt.input)
				if tt.wantErr {
					assert.Error(t, err)
					return
				}
				require.NoError(t, err)
				assert.Equal(t, tt.input.Name(), got.Name())

				// Verify update
				found, _ := repo.GetById(ctx, tt.input.Id())
				assert.Equal(t, "updated-name", found.Name())
			})
		}
	})

	t.Run("Delete", func(t *testing.T) {
		// Prepare user to delete
		userToDelete := mustNewUser("name", []*models.Task{})
		_, _ = repo.Create(ctx, userToDelete)

		tests := []struct {
			name    string
			id      *models.UserId
			wantErr bool
		}{
			{
				name:    "Success",
				id:      userToDelete.Id(),
				wantErr: false,
			},
			{
				name:    "NotFound",
				id:      models.NewUserId(),
				wantErr: true,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				err := repo.Delete(ctx, tt.id)
				if tt.wantErr {
					assert.Error(t, err)
					return
				}
				require.NoError(t, err)

				// Verify deletion
				_, err = repo.GetById(ctx, tt.id)
				assert.Error(t, err)
			})
		}
	})
}

func TestTaskRepositoryImpl(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	repo := NewTaskRepositoryImpl(logger)
	ctx := context.Background()

	t.Run("Create", func(t *testing.T) {
		tests := []struct {
			name    string
			input   *models.Task
			wantErr bool
		}{
			{
				name:    "Success",
				input:   mustNewTask("task-name-1"),
				wantErr: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := repo.Create(ctx, tt.input)
				if tt.wantErr {
					assert.Error(t, err)
					return
				}
				require.NoError(t, err)
				assert.Equal(t, tt.input, got)
			})
		}
	})

	t.Run("GetAll", func(t *testing.T) {
		tests := []struct {
			name    string
			wantErr bool
		}{
			{
				name:    "Success",
				wantErr: false,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := repo.GetAll(ctx)
				if tt.wantErr {
					assert.Error(t, err)
					return
				}
				require.NoError(t, err)
				assert.NotNil(t, got)
			})
		}
	})
}

func mustNewUser(name string, tasks []*models.Task) *models.User {
	u, err := models.NewUser(name, tasks)
	if err != nil {
		panic(err)
	}
	return u
}

func mustNewTask(name string) *models.Task {
	t, err := models.NewTask(name)
	if err != nil {
		panic(err)
	}
	return t
}

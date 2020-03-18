package sqlstore

import (
	"context"
	"github.com/dev-tim/message-board-api/internal/app/common"
	"github.com/dev-tim/message-board-api/internal/app/model"
	"time"
)

type MessagesRepository struct {
	store *Store
}

func NewMessageRepository(store *Store) *MessagesRepository {
	return &MessagesRepository{
		store: store,
	}
}

func (r *MessagesRepository) Create(m *model.Message) (*model.Message, error) {
	logger := common.GetLogger()

	timeout := 3 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	logger.Debug("Starting message insert exec")

	if err := r.store.db.QueryRowContext(ctx, "INSERT into "+
		"messages (id, name, email, text, external_creation_time) "+
		"VALUES   ($1, $2,   $3,    $4,   $5) "+
		"RETURNING id",
		m.Id, m.Name, m.Email, m.Text, m.CreationTime).Scan(&m.Id); err != nil {
		logger.Error("Failed to insert new message ", err)

		return nil, err
	}

	logger.Debug("Finished inserting new message to db")
	return m, nil
}

func (r *MessagesRepository) FindById(id string) (*model.Message, error) {
	logger := common.GetLogger()

	timeout := 3 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	logger.Debug("Starting message select exec")
	m := &model.Message{}
	err := r.store.db.QueryRowContext(ctx, "SELECT * FROM messages WHERE id=$1", id).Scan(&m.Id,
		&m.Name,
		&m.Email,
		&m.Text,
		&m.CreationTime,
		&m.CreatedAt,
		&m.UpdatedAt)

	if err != nil {
		logger.Error("Failed to select message by id", id, err)
		return nil, err
	}

	logger.Debug("Finished selecting message from db")
	return m, nil
}

func (r *MessagesRepository) FindLatest(limit, offset int) ([]*model.Message, error) {
	logger := common.GetLogger()

	timeout := 3 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	logger.Debug("Starting message select exec")

	rows, err := r.store.db.QueryContext(ctx, "SELECT * FROM messages ORDER BY external_creation_time DESC LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		logger.Error("Failed to select messageList", err)
		return nil, err
	}
	defer rows.Close()

	messageList := make([]*model.Message, 0)

	for rows.Next() {
		var m model.Message
		if err := rows.Scan(&m.Id,
			&m.Name,
			&m.Email,
			&m.Text,
			&m.CreationTime,
			&m.CreatedAt,
			&m.UpdatedAt); err != nil {
			logger.Error("Failed to parse message", err)
		}
		messageList = append(messageList, &m)
	}

	rerr := rows.Close()
	if rerr != nil {
		logger.Error("Failed to close result set", err)
	}

	logger.Debug("Finished selecting new message to db")
	return messageList, nil
}

func (r *MessagesRepository) Update(id, text string) (*int64, error) {
	logger := common.GetLogger()

	timeout := 3 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	logger.Debug("Starting message select exec")

	res, err := r.store.db.ExecContext(ctx, "UPDATE messages SET text=$1 WHERE id=$2", text, id)
	if err != nil {
		logger.Error("Failed to update messages", err)
		return nil, err
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		logger.Error("Failed to get number of updated messages", err)
		return nil, err
	}

	logger.Debug("Finished selecting new message to db")

	return &affectedRows, nil
}

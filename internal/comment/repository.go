package comment

import (
	"context"
	"github.com/jxlwqq/blog-microservices/internal/pkg/dbcontext"
	"github.com/jxlwqq/blog-microservices/internal/pkg/log"
)

func NewRepository(logger *log.Logger, db *dbcontext.DB) Repository {
	return repository{
		logger: logger,
		db:     db,
	}
}

type Repository interface {
	Create(ctx context.Context, comment *Comment) error
	Update(ctx context.Context, comment *Comment) error
	Delete(ctx context.Context, id uint64) error
	DeleteByUUID(ctx context.Context, uuid string) error
	ListByPostID(ctx context.Context, postID uint64, offset, limit int) ([]*Comment, error)
	Get(ctx context.Context, id uint64) (*Comment, error)
	GetByUUID(ctx context.Context, uuid string) (*Comment, error)
	CountByPostID(ctx context.Context, postID uint64) (uint64, error)
}

type repository struct {
	logger *log.Logger
	db     *dbcontext.DB
}

func (r repository) CountByPostID(ctx context.Context, postID uint64) (uint64, error) {
	var count int64
	err := r.db.Model(&Comment{}).Where("post_id = ?", postID).Count(&count).Error
	return uint64(count), err
}

func (r repository) Get(ctx context.Context, id uint64) (*Comment, error) {
	comment := &Comment{}
	err := r.db.First(comment, id).Error
	return comment, err
}

func (r repository) GetByUUID(ctx context.Context, uuid string) (*Comment, error) {
	comment := &Comment{}
	err := r.db.First(comment, "uuid = ?", uuid).Error
	return comment, err
}

func (r repository) Create(ctx context.Context, comment *Comment) error {
	return r.db.Create(comment).Error
}

func (r repository) Update(ctx context.Context, comment *Comment) error {
	return r.db.Save(comment).Error
}

func (r repository) Delete(ctx context.Context, id uint64) error {
	return r.db.Delete(&Comment{ID: id}).Error
}

func (r repository) DeleteByUUID(ctx context.Context, uuid string) error {
	return r.db.Delete(&Comment{UUID: uuid}).Error
}

func (r repository) ListByPostID(ctx context.Context, postID uint64, offset, limit int) ([]*Comment, error) {
	var comments []*Comment
	err := r.db.Where("post_id = ?", postID).Offset(offset).Limit(limit).Find(&comments).Error
	return comments, err
}

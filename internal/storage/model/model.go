package model

import (
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"

	"gorm.io/gorm"

	"github.com/jdxj/cyber-wagon/config"
	"github.com/jdxj/cyber-wagon/internal/storage/dao"
	"github.com/jdxj/cyber-wagon/internal/util"
)

var (
	ErrEmptyPath    = errors.New("empty path")
	ErrFileNotFound = errors.New("file not found")
)

func NewStorage(cfg config.Storage) *Storage {
	if cfg.Path == "" {
		panic(ErrEmptyPath)
	}
	return &Storage{path: cfg.Path}
}

type Storage struct {
	path string
}

func (s *Storage) WriteFile(ctx context.Context, fileID, userID uint64, filename string, r io.Reader) error {
	tmpPath := filepath.Join(os.TempDir(), tempName())
	tmpFile, err := os.Create(tmpPath)
	if err != nil {
		return err
	}
	defer func() {
		_ = tmpFile.Sync()
		_ = tmpFile.Close()
	}()

	mf := newMD5Filter(tmpFile)
	if _, err = io.Copy(mf, r); err != nil {
		return err
	}

	var (
		sum     = mf.Sum()
		newDir  = filepath.Join(s.path, sum[:3])
		newPath = filepath.Join(newDir, sum)
	)
	err = os.MkdirAll(newDir, os.ModePerm)
	if err != nil {
		return err
	}

	if _, err = os.Stat(newPath); os.IsNotExist(err) {
		if err = os.Rename(tmpPath, newPath); err != nil {
			return err
		}
	}

	f := &dao.File{
		Model:    gorm.Model{ID: uint(fileID)},
		UserID:   userID,
		Filename: filename,
		MD5:      sum,
	}
	return util.DB.WithContext(ctx).
		Create(f).Error
}

func (s *Storage) ReadFile(ctx context.Context, fileID, userID uint64) (*FileInfo, error) {
	f := &dao.File{
		Model:  gorm.Model{ID: uint(fileID)},
		UserID: userID,
	}
	err := util.DB.WithContext(ctx).
		Take(f).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrFileNotFound
		}
		return nil, err
	}

	if len(f.MD5) < 3 {
		return nil, ErrFileNotFound
	}

	fi := &FileInfo{
		ID:        uint64(f.ID),
		CreatedAt: f.CreatedAt,
		Filename:  f.Filename,
		UserID:    f.UserID,
		MD5:       f.MD5,
		path:      filepath.Join(s.path, f.MD5[:3], f.MD5),
	}
	return fi, nil
}

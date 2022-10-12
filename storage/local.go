package storage

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"syscall"

	"github.com/huacnlee/gobackup/helper"
	"github.com/huacnlee/gobackup/logger"
)

// Local storage
//
// type: local
// path: /data/backups
type Local struct {
	Base
	destPath string
}

func (ctx *Local) open() (err error) {
	ctx.destPath = ctx.model.StoreWith.Viper.GetString("path")
	helper.MkdirP(ctx.destPath)
	return
}

func (ctx *Local) close() {}

func (ctx *Local) upload(fileKey string) (err error) {
	fullPath := filepath.Clean(path.Join(ctx.destPath, filepath.Base(ctx.archivePath)))
	err = moveFile(ctx.archivePath, fullPath)
	if err != nil {
		return err
	}
	logger.Info("Store successed", fullPath)
	return nil
}

func (ctx *Local) delete(fileKey string) error {
	return os.RemoveAll(path.Join(ctx.destPath, fileKey))
}

func moveFile(src string, dst string) error {
	isDir, _ := isDirectory(dst)
	if isDir {
		fileName := filepath.Base(src)
		return moveFile(src, filepath.Clean(path.Join(dst, fileName)))
	}
	if runtime.GOOS == "windows" {
		from, _ := syscall.UTF16PtrFromString(src)
		to, _ := syscall.UTF16PtrFromString(dst)
		return syscall.MoveFile(from, to)
	} else {
		return os.Rename(src, dst)
	}
}

// isDirectory determines if a file represented
// by `path` is a directory or not
func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), err
}

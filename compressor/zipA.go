package compressor

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"

	"github.com/huacnlee/gobackup/logger"
)

// Zip .zip compressor
//
// type: zip
type ZipA struct {
	Base
}

func (ctx *ZipA) perform() (archivePath string, err error) {
	filePath := ctx.archiveFilePath(".zip")

	err = zipDir(ctx.model.DumpPath, filePath)
	if err == nil {
		archivePath = filePath
		return
	}
	return
}

func zipDir(dir, zipFile string) error {

	fz, err := os.Create(zipFile)
	if err != nil {
		logger.Error("Create zip file failed: %s\n", err.Error())
		return err
	}
	defer fz.Close()

	w := zip.NewWriter(fz)
	defer w.Close()

	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fDest, err := w.Create(path[len(dir)+1:])
			if err != nil {
				logger.Error("Create failed: %s\n", err.Error())
				return nil
			}
			fSrc, err := os.Open(path)
			if err != nil {
				logger.Error("Open failed: %s\n", err.Error())
				return nil
			}
			defer fSrc.Close()
			_, err = io.Copy(fDest, fSrc)
			if err != nil {
				logger.Error("Copy failed: %s\n", err.Error())
				return nil
			}
		}
		return nil
	})
}

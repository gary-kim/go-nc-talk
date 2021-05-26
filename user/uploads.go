package user

import (
	"errors"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"

	"gomod.garykim.dev/nc-talk/constants"

	"github.com/studio-b12/gowebdav"
)

var (
	duplicateSuffixRegex = regexp.MustCompile(` \((\d+)\)$`)

	// ErrIsNotDirectory is returned if a path expected to be a directory is not a directory
	ErrIsNotDirectory = errors.New("is not a directory")
)

// UploadFile uploads the given data to the talk attachments
// directory and returns the filepath it was uploaded at.
//
// If the directory does not exist, the function will create the directory.
//
// Provide only the name of the file, not the full path. If a full path is provided,
// only the basename will be used. Use user.TalkUser.UploadFileAtPath for uploading
// to a specific directory.
func (t *TalkUser) UploadFile(data *[]byte, name string) (finalPath string, err error) {
	capabilities, err := t.Capabilities()
	if err != nil {
		return
	}
	finalPath = path.Clean(capabilities.AttachmentsFolder + "/" + path.Base(name))
	finalPath, err = t.uploadFilePath(finalPath)
	if err != nil {
		return
	}
	err = t.UploadFileAtPath(data, finalPath)
	return
}

// UploadFileAtPath uploads the given data at the given path.
func (t *TalkUser) UploadFileAtPath(data *[]byte, givenPath string) error {
	c := t.getWebdavClient()
	err := c.Write(givenPath, *data, 0644)
	if err != nil {
		if ferr, ok := err.(*os.PathError); !ok || strings.HasPrefix(ferr.Err.Error(), "404") {
			err = c.MkdirAll(path.Dir(givenPath), 0644)
			if err != nil {
				return err
			}
			err = c.Write(givenPath, *data, 0644)
		}
	}
	return err
}

// uploadFilePath provides a unique path to upload a file with the given path.
func (t *TalkUser) uploadFilePath(name string) (finalPath string, err error) {
	c := t.getWebdavClient()

	capabilities, err := t.Capabilities()
	if err != nil {
		return
	}

	statInfo, err := c.Stat(capabilities.AttachmentsFolder)
	if err != nil {
		// The error may be because it does not exist
		if strings.HasPrefix(err.(*os.PathError).Err.Error(), "404 Not Found - PROPFIND") {
			// Directory does not exist. Return the given path directly.
			return name, nil
		}
		return
	}
	if !statInfo.IsDir() {
		return "", ErrIsNotDirectory
	}

	files, err := c.ReadDir(capabilities.AttachmentsFolder)
	if err != nil {
		return
	}

	filename := path.Base(name)
	if !includesFileName(files, filename) {
		return filename, nil
	}

	extension := path.Ext(filename)
	basename := strings.TrimSuffix(filename, extension)

	suffix := duplicateSuffixRegex.Find([]byte(basename))
	nameWithoutSuffix := strings.TrimSuffix(basename, string(suffix))

	// Loop until a unique path is found
	for i := 2; true; i++ {
		uniqueName := nameWithoutSuffix + " (" + strconv.Itoa(i) + ")" + extension
		if !includesFileName(files, uniqueName) {
			finalPath = path.Dir(name) + "/" + uniqueName
			return
		}
	}
	return "", errors.New("should never reach")
}

func includesFileName(filelist []os.FileInfo, name string) bool {
	for _, v := range filelist {
		if v.Name() == name {
			return true
		}
	}
	return false
}

func (t *TalkUser) getWebdavClient() *gowebdav.Client {
	if t.webdavclient == nil {
		url := t.NextcloudURL + constants.RemoteDavEndpoint(t.User, "files")
		t.webdavclient = gowebdav.NewClient(url, t.User, t.Pass)
	}
	return t.webdavclient
}

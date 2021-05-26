package user

import (
	"fmt"
	"net"
	"net/http"
	"strconv"
	"testing"

	"github.com/monaco-io/request"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/webdav"
)

func startTestServer() (int, error) {
	wd := &webdav.Handler{
		Prefix:     "/remote.php/dav/files/testuser/",
		FileSystem: webdav.NewMemFS(),
		LockSystem: webdav.NewMemLS(),
	}

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return 0, err
	}
	port := listener.Addr().(*net.TCPAddr).Port
	go func() {
		fmt.Println("Serving on port: " + strconv.Itoa(port))
		err := http.Serve(listener, wd)
		if err != nil {
			panic(err)
		}
	}()
	return port, err
}

func TestFileUploads(t *testing.T) {
	port, err := startTestServer()
	assert.NoError(t, err, "Starting server")

	user, err := NewUser("http://localhost:"+strconv.Itoa(port), "testuser", "password", nil)
	assert.NoError(t, err, "Creating a new user")

	user.capabilities = &Capabilities{
		AttachmentsFolder:  "/Talk",
		AttachmentsAllowed: true,
	}

	someBytes := &[]byte{21, 65, 56, 57, 2, 95, 100, 85}

	filepath, err := user.UploadFile(someBytes, "testfile.txt")
	assert.NoError(t, err, "creating testfile.txt")
	assert.Equal(t, "/Talk/testfile.txt", filepath, "ensuring file is in the expected location")

	c := user.RequestClient(request.Client{
		URL: "/remote.php/dav/files/testuser/Talk/testfile.txt",
	})
	resp, err := c.Do()
	assert.NoError(t, err, "downloading file to check it is identical")
	assert.Equal(t, someBytes, &resp.Data, "checking that the data is the same after being downloaded")

	filepath, err = user.UploadFile(someBytes, "testfile.txt")
	assert.NoError(t, err, "creating a second testfile.txt")
	assert.Equal(t, "/Talk/testfile (2).txt", filepath, "ensuring file is in the expected location")

	c = user.RequestClient(request.Client{
		URL: "/remote.php/dav/files/testuser/Talk/testfile (2).txt",
	})
	resp, err = c.Do()
	assert.NoError(t, err, "downloading file to check it is identical")
	assert.Equal(t, someBytes, &resp.Data, "checking that the data is the same after being downloaded")

	filepath, err = user.UploadFile(someBytes, "testfile.txt")
	assert.NoError(t, err, "creating a third testfile.txt")
	assert.Equal(t, "/Talk/testfile (3).txt", filepath, "ensuring file is in the expected location")

	c = user.RequestClient(request.Client{
		URL: "/remote.php/dav/files/testuser/Talk/testfile (3).txt",
	})
	resp, err = c.Do()
	assert.NoError(t, err, "downloading file to check it is identical")
	assert.Equal(t, someBytes, &resp.Data, "checking that the data is the same after being downloaded")

	filepath, err = user.UploadFile(someBytes, "testfile (3).txt")
	assert.NoError(t, err, "creating a file with an already existing increment: testfile (2).txt")
	assert.Equal(t, "/Talk/testfile (4).txt", filepath, "ensuring file is in the expected location")

	err = user.UploadFileAtPath(someBytes, "/asdf/asdf/testing.txt")
	assert.NoError(t, err, "attempting upload with UploadFileAtPath")

	c = user.RequestClient(request.Client{
		URL: "/remote.php/dav/files/testuser/asdf/asdf/testing.txt",
	})
	resp, err = c.Do()
	assert.NoError(t, err, "downloading file to check it is identical")
	assert.Equal(t, someBytes, &resp.Data, "checking that the data is the same after being downloaded")

	// Sanity check
	c = user.RequestClient(request.Client{
		URL: "/remote.php/dav/files/testuser/asdf/asdf/testing bla.txt",
	})
	resp, err = c.Do()
	assert.NoError(t, err, "downloading file to check it is identical")
	assert.Equal(t, http.StatusNotFound, resp.StatusCode(), "making sure a file that does not exist is not found")
}

// Wrapper of golang's standard logger.
// Provide feature of preparing log folder and attaching stdout and file stream to logger's stream with just two functions.
package justlog // import "github.com/dictor/justlog"

import (
	"github.com/kardianos/osext"
	"io"
	"log"
	"os"
	"time"
)

// This variable is setted as executed binary's directory path after SetPath() function is called.
var ExePath string

// Retrieve executed binary's directory path and prepare log folder.
// Log folder name is 'log' and it's created under retrieved path when it doesn't exist.
func SetPath() (string, error) {
	path, err := osext.ExecutableFolder()
	ExePath = path
	path += "/log"
	prepareDirectory(path)
	return path, err
}

// Helper function for simple code.
// It cause panic when err isn't nil.
func MustPath(path string, err error) string {
	if err != nil {
		log.Panic(err)
	}
	return path
}

// Open file handler with name as current time string and attach file handler to logger's stream.
func SetStream(path string) (*os.File, error) {
	fpLog, err := os.OpenFile(path+"/"+time.Now().Format("2006-01-02T15_04_05")+".txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0775)
	if err != nil {
		return nil, err
	} else {
		multiWriter := io.MultiWriter(fpLog, os.Stdout)
		log.SetOutput(multiWriter)
	}
	return fpLog, nil
}

// Helper function for simple code.
// It cause panic when err isn't nil.
func MustStream(s *os.File, err error) *os.File {
	if err != nil {
		log.Panic(err)
	}
	return s
}

// Create all recursive directory when it doesn't exist.
func prepareDirectory(dir ...string) {
	for _, val := range dir {
		if _, err := os.Stat(val); os.IsNotExist(err) {
			os.Mkdir(val, 0775)
		}
	}
}

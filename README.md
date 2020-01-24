# justlog
**justlog** is log helper library for golang.
* Prepare directory for log file in same directory with executed binary 
* Set stream in standard go logger to os.Stdout and log file

## How to use
Import library
```
import (
	"github.com/dictor/justlog"
)
```

**Main code**
```
log_path := justlog.MustPath(justlog.SetPath())
defer (justlog.MustStream(justlog.SetStream(log_path))).Close()
log.Println("Hello world!!")
```

Log file will be created like below tree.
```
├── binary
└── log
    └── 2020-01-24T14_19_04.txt
```

In log file and stdout,
```
2020/01/24 14:19:04 Hello world!!
```

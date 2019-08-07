# Code Style

We follow the code style used to develop the Go language, inspiring by the [code review of Go code](https://github.com/golang/go/wiki/CodeReviewComments)

Complimentary you should read [Effective Go](https://golang.org/doc/effective_go.html)

## Anything not covered in the above resources

- Functions or methods are ordered by the dependency relationship, such that the most dependent function or method 
should be at the top. In the following example, ExecCmdDirBytes is the most fundamental function, it's called by ExecCmdDir, and ExecCmdDir is also called by ExecCmd:

    ```go
     // ExecCmdDirBytes executes system command in given directory
     // and return stdout, stderr in bytes type, along with possible error.
     func ExecCmdDirBytes(dir, cmdName string, args ...string) ([]byte, []byte, error) {
        ...
     }
    
     // ExecCmdDir executes system command in given directory
     // and return stdout, stderr in string type, along with possible error.
     func ExecCmdDir(dir, cmdName string, args ...string) (string, string, error) {
        bufOut, bufErr, err := ExecCmdDirBytes(dir, cmdName, args...)
        return string(bufOut), string(bufErr), err
     }
    
     // ExecCmd executes system command
     // and return stdout, stderr in string type, along with possible error.
     func ExecCmd(cmdName string, args ...string) (string, string, error) {
        return ExecCmdDir("", cmdName, args...)
     }
     ```
 
 - Methods of struct should be put after struct definition, and order them by the order of fields they mostly operate
  on:
 
    ```go
    type Webhook struct { ... }
    func (w *Webhook) GetEvent() { ... }
    func (w *Webhook) SaveEvent() error { ... }
    func (w *Webhook) HasPushEvent() bool { ... }
    ```
 
 - If a struct has operational functions, should basically follow the `CRUD` order:
 
 	```go
 	func CreateWebhook(w *Webhook) error { ... }
 	func GetWebhookById(hookId int64) (*Webhook, error) { ... }
 	func UpdateWebhook(w *Webhook) error { ... }
 	func DeleteWebhook(hookId int64) error { ... }
 	```
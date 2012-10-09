package gopdf
import (
  "bytes"
  "os/exec"
  "net/http"
  "fmt"
  "flag"
  "runtime"
)
func PDF(input string) bytes.Buffer {
    cmd := exec.Command("wkhtmltopdf", "-", "-")
    cmd.Stdin = bytes.NewBufferString(input)
    var out bytes.Buffer
    cmd.Stdout = &out
    go cmd.Run()
    return out
}

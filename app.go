package gopdf
import (
  "bytes"
  "os/exec"
)
func PDF(input string) bytes.Buffer {
    cmd := exec.Command("wkhtmltopdf", "-", "-")
    cmd.Stdin = bytes.NewBufferString(input)
    var out bytes.Buffer
    cmd.Stdout = &out
    cmd.Run()
    return out
}

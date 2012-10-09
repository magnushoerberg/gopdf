package gopdf
import (
  "bytes"
  "os/exec"
  "log"
)
func PDF(input string) bytes.Buffer {
    wkhtmltopdfPath := "/usr/local/bin/"
    cmd := exec.Command(wkhtmltopdfPath + "wkhtmltopdf", "-", "-")
    cmd.Stdin = bytes.NewBufferString(input)
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
    return out
}

package utils 

import (
	"fmt"
	"net/http"
	mand "math/rand"
	"os"
	"time"
	"io"
)

func Check(e error) {
    if e != nil {
        panic(e)
    }
}

func Generate(length int) string {
	mand.Seed(time.Now().UnixNano())
	charset :=  "abcdefghijklmnopqrstuvwxyzABCDEFGHIKLMNOPQRSTUVWXYZ1234567890_!@#"
	pwd := ""
	for i := 0; i<length; i++ {
		pwd += string(charset[mand.Intn(len(charset))]) // [0,n)
	}
	return pwd
}

func PadRight(str string, length int) string {
	for len(str) < length {
	   str = str + "0"
	}
	return str
 }

 func DownloadFile(filepath string, url string) (err error) {

	out, err := os.Create(filepath)
	if err != nil  {
	  return err
	}

	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
	  return err
	}
	defer resp.Body.Close()
  
	// Check server response
	if resp.StatusCode != http.StatusOK {
	  return fmt.Errorf("bad status: %s", resp.Status)
	}
  
	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil  {
	  return err
	}
  
	return nil
  }

  func Remove(s []string, i int) []string {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"

	"github.com/Unknwon/com"
	"gopkg.in/cheggaaa/pb.v2"
	"gopkg.in/yaml.v2"
)

const packageFile = "package.yaml"
const packagesDir = "./packages/"

func main() {
	var err error
	var data []byte
	if data, err = ioutil.ReadFile(packageFile); err != nil {
		panic(err)
	}

	var list = make(map[string]map[string]string)
	if err = yaml.Unmarshal(data, &list); err != nil {
		panic(err)
	}

	for key, val := range list {
		var filename, url string

		if strings.Contains(val["filename"], "%s") {
			filename = fmt.Sprintf(val["filename"], val["version"])
		} else {
			filename = val["filename"]
		}

		if strings.Contains(val["url"], "%s") {
			url = fmt.Sprintf(val["url"], val["version"])
		} else {
			url = val["url"]
		}

		if val["skip"] != "true" {
			for {
				if _, err = os.Stat(filename); err != nil {
					downloadFile(filename, url, val["size"])
				}
				var verify bool
				if verify, err = hashFile(filename, val["sha256"]); err != nil {
					fmt.Printf("Downloaded file hash with error: %v, try again.\n", err)
				} else if !verify {
					fmt.Printf("Downloaded %s verified failed, try again.\n", filename)
					if err = os.Remove(filename); err != nil {
						panic(err)
					}
				} else {
					break
				}
			}
			fmt.Printf("Download %s success.\n", key)
		} else {
			if err = com.Copy(packagesDir+filename, filename); err != nil {
				fmt.Errorf("Copy %s to here with error: %v.\n", key, err)
				continue
			}
			var verify bool
			if verify, err = hashFile(filename, val["sha256"]); err != nil {
				fmt.Printf("Downloaded file hash with error: %v, try again.\n", err)
				return
			} else if !verify {
				fmt.Printf("Downloaded %s verified failed, try again.\n", filename)
				return
			}
		}

		fmt.Printf("Verify %s success.\n", key)

		var info os.FileInfo
		if info, err = os.Stat(filename); err != nil {
			fmt.Printf("Get file info failed with error: %v", err)
		} else {
			val["size"] = strconv.FormatInt(info.Size(), 10)
		}

		var targetDir = strings.TrimSuffix(filename, ".tar"+path.Ext(filename))
		if com.IsFile(targetDir) {
			if err = os.RemoveAll(targetDir); err != nil {
				fmt.Errorf("Remove %s with error: %v\n", targetDir, err)
			}
		}

		var uncompressFlag = "zxf"
		if path.Ext(filename) == ".bz2" {
			uncompressFlag = "xjf"
		} else if path.Ext(filename) == ".xz" {
			uncompressFlag = "xJf"
		}
		if err = exec.Command("tar", uncompressFlag, filename).Run(); err != nil {
			fmt.Errorf("Uncompress %s with error: %v.\n", key, err)
		} else {
			fmt.Printf("Uncompress %s success.\n", key)
		}

		if com.IsDir(key) {
			if err = os.RemoveAll(key); err != nil {
				fmt.Errorf("Remove %s with error: %v.\n", key, err)
			}
		}

		if err = os.Rename(targetDir, key); err != nil {
			fmt.Printf("Rename %s to %s with error: %v.\n", targetDir, key, err)
		}
	}
	var outDate []byte
	if outDate, err = yaml.Marshal(list); err != nil {
		panic(err)
	} else if err = ioutil.WriteFile(packageFile, outDate, 0644); err != nil {
		panic(err)
	}

}

func downloadFile(file, url, size string) (err error) {
	var out *os.File
	if out, err = os.Create(file); err != nil {
		return
	}
	defer out.Close()

	var resp *http.Response
	if resp, err = http.Get(url); err != nil {
		return
	}
	defer resp.Body.Close()

	var totalLen int64
	if len(resp.Header["Content-Length"]) == 0 {
		if size != "" {
			if totalLen, err = strconv.ParseInt(size, 10, 64); err != nil {
				return
			}
		} else {
			totalLen = math.MaxInt64
		}
	} else {
		if totalLen, err = strconv.ParseInt(resp.Header["Content-Length"][0], 10, 64); err != nil {
			return
		}
	}

	bar := pb.Start64(totalLen)
	defer bar.Finish()
	bar.Set("prefix", file+": ")

	var buf = make([]byte, 10240)
	var n int
	for {
		n, err = resp.Body.Read(buf)
		if n != 0 {
			if _, err = out.Write(buf[:n]); err != nil {
				return
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
		bar.Add(n)
	}
	return
}

func hashFile(file, hash string) (verify bool, err error) {
	if !com.IsFile(file) {
		err = fmt.Errorf("No such a file: %s", file)
		return
	}
	var fileObj *os.File
	if fileObj, err = os.Open(file); err != nil {
		return
	}
	var data = make([]byte, 10240)
	var n int
	h := sha256.New()
	for {
		n, err = fileObj.Read(data)
		if n != 0 {
			h.Write(data[:n])
		}
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			return
		}
	}
	verify = hex.EncodeToString(h.Sum(nil)) == hash
	return
}

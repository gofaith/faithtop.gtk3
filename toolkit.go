package faithtop

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/StevenZack/tools/fileToolkit"
	"github.com/StevenZack/tools/netToolkit"
	"github.com/StevenZack/tools/strToolkit"
)

var cacheDir string

func GetCacheDir() string {
	if cacheDir != "" {
		return cacheDir
	}
	sep := string(os.PathSeparator)
	homeDir := fileToolkit.GetHomeDir()
	if runtime.GOOS == "linux" {
		cacheDir = strToolkit.Getrpath(homeDir) + ".cache/faithtop/"
		return cacheDir
	} else if runtime.GOOS == "windows" {
		cacheDir = strToolkit.Getrpath(homeDir) + "AppData" + sep + "Local" + sep + "faithtop" + sep
		return cacheDir
	} else {
		cacheDir = fileToolkit.GetCurrentExecPath() + ".cache" + sep
		return cacheDir
	}
}
func SetCacheDir(dir string) {
	cacheDir = strToolkit.Getrpath(dir)
}

func CacheNetFile(url, cacheDir string, callback func(string)) {
	f := strToolkit.Getrpath(cacheDir) + Url2cachePath(url)
	if _, e := os.Stat(f); e != nil {
		e = netToolkit.DownloadFile(url, f)
		if e != nil {
			fmt.Println(e)
			return
		}
	}
	callback(f)
}
func Url2cachePath(url string) string {
	rUrl := GetRealUrl(url)
	s := strings.Replace(rUrl, "://", "/", -1)
	sep := string(os.PathSeparator)
	s = strings.Replace(s, "/", sep, -1)
	if strings.HasSuffix(s, sep) {
		return s[:len(s)-1]
	}
	return s
}
func GetRealUrl(url string) string {
	for i := 0; i < len(url); i++ {
		item := url[i : i+1]
		if item == "?" || item == "#" {
			return url[:i]
		}
	}
	return url
}

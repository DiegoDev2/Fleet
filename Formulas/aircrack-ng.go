// Copyright (C) 2024 Fleet Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package formulas

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sync"
)

func InstallAircrackNg() {
	switch runtime.GOOS {
	case "darwin":
		installAircrackNgMac()
	default:
		fmt.Println("This script only supports macOS (darwin).")
	}
}

func installAircrackNgMac() {
	fmt.Println("Starting Aircrack-ng installation 🚀")

	var wg sync.WaitGroup
	errChan := make(chan error, 1)

	wg.Add(1)
	go func() {
		defer wg.Done()
		zipURL := "https://github.com/aircrack-ng/aircrack-ng/archive/refs/heads/master.zip"
		zipFile := "aircrack-ng.zip"

		if err := downloadFile(zipFile, zipURL); err != nil {
			errChan <- fmt.Errorf("error downloading aircrack-ng: %v", err)
			return
		}
		defer os.Remove(zipFile)

		if err := unzip(zipFile, "aircrack-ng"); err != nil {
			errChan <- fmt.Errorf("error extracting aircrack-ng: %v", err)
			return
		}

		if err := os.Chdir("aircrack-ng/aircrack-ng-master"); err != nil {
			errChan <- fmt.Errorf("error changing to aircrack-ng directory: %v", err)
			return
		}
	}()

	go func() {
		wg.Wait()
		close(errChan)
	}()

	if err := <-errChan; err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Building aircrack-ng...")
	if err := runCommand("autoreconf", "-i"); err != nil {
		fmt.Println("Error running autoreconf:", err)
		return
	}
	if err := runCommand("./configure", "--with-experimental"); err != nil {
		fmt.Println("Error running configure:", err)
		return
	}
	if err := runCommand("make"); err != nil {
		fmt.Println("Error running make:", err)
		return
	}
	if err := runCommand("sudo", "make", "install"); err != nil {
		fmt.Println("Error running make install:", err)
		return
	}

	fmt.Println("aircrack-ng installed successfully 🎉")
}

func downloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func unzip(src string, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		fPath := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(fPath, os.ModePerm); err != nil {
				return err
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fPath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, rc)
		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}
	}
	return nil
}

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

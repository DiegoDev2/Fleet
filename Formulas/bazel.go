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
//
// package formulas


package formulas

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sync"
)

func InstallBazel() {
	switch runtime.GOOS {
	case "darwin":
		installBazelMac()
	default:
		redBold.Println("This script only supports macOS (darwin).")
	}
}

func installBazelMac() {
	var wg sync.WaitGroup
	errCh := make(chan error, 4)

	boldGreen.Println("Starting Bazel installation 🚀")

	wg.Add(1)
	go func() {
		defer wg.Done()
		yellow.Println("Downloading Bazel...")

		download := exec.Command("curl", "-fLO", "https://github.com/bazelbuild/bazel/releases/download/5.0.0/bazel-5.0.0-installer-darwin-x86_64.sh")
		if err := download.Run(); err != nil {
			errCh <- fmt.Errorf("error downloading Bazel: %w", err)
			return
		}
		yellow.Println("Download completed.")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		yellow.Println("Setting execute permissions for Bazel installer...")

		if _, err := os.Stat("bazel-5.0.0-installer-darwin-x86_64.sh"); os.IsNotExist(err) {
			errCh <- fmt.Errorf("Bazel installer file not found: %w", err)
			return
		}

		chmod := exec.Command("chmod", "+x", "bazel-5.0.0-installer-darwin-x86_64.sh")
		if err := chmod.Run(); err != nil {
			errCh <- fmt.Errorf("error setting permissions for Bazel installer: %w", err)
			return
		}
		yellow.Println("Permissions set.")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		yellow.Println("Installing Bazel...")

		if _, err := os.Stat("bazel-5.0.0-installer-darwin-x86_64.sh"); os.IsNotExist(err) {
			errCh <- fmt.Errorf("Bazel installer file not found: %w", err)
			return
		}

		install := exec.Command("bash", "bazel-5.0.0-installer-darwin-x86_64.sh")
		if err := install.Run(); err != nil {
			errCh <- fmt.Errorf("error installing Bazel: %w", err)
			return
		}
		yellow.Println("Installation completed.")
	}()

	wg.Wait()
	close(errCh)

	for err := range errCh {
		redBold.Println(err)
		return
	}

	boldGreen.Println("Bazel installed successfully 🎉")
}

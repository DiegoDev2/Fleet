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
	"fmt"
	"os/exec"
	"runtime"
)


func InstallBat(){

 switch runtime.GOOS {
  case "darwin":
    installBazelMac()
  default:
    fmt.Println("OS No Support")
 }
 
}

func InstallBatMac(){
  url := "https://github.com/sharkdp/bat/releases/download/v0.24.0/bat-v0.24.0-x86_64-apple-darwin.tar.gz"
  download := exec.Command("curl", "-L", url, "-o", "bat.tar.gz")
  download.Run()
 }

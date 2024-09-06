package main

// LaunchctlCompletion - Bash completion for Launchctl
// Homepage: https://github.com/CamJN/launchctl-completion

import (
	"fmt"
	
	"os/exec"
)

func installLaunchctlCompletion() {
	// Método 1: Descargar y extraer .tar.gz
	launchctlcompletion_tar_url := "https://github.com/CamJN/launchctl-completion/archive/refs/tags/v1.0.tar.gz"
	launchctlcompletion_cmd_tar := exec.Command("curl", "-L", launchctlcompletion_tar_url, "-o", "package.tar.gz")
	err := launchctlcompletion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	launchctlcompletion_zip_url := "https://github.com/CamJN/launchctl-completion/archive/refs/tags/v1.0.zip"
	launchctlcompletion_cmd_zip := exec.Command("curl", "-L", launchctlcompletion_zip_url, "-o", "package.zip")
	err = launchctlcompletion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	launchctlcompletion_bin_url := "https://github.com/CamJN/launchctl-completion/archive/refs/tags/v1.0.bin"
	launchctlcompletion_cmd_bin := exec.Command("curl", "-L", launchctlcompletion_bin_url, "-o", "binary.bin")
	err = launchctlcompletion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	launchctlcompletion_src_url := "https://github.com/CamJN/launchctl-completion/archive/refs/tags/v1.0.src.tar.gz"
	launchctlcompletion_cmd_src := exec.Command("curl", "-L", launchctlcompletion_src_url, "-o", "source.tar.gz")
	err = launchctlcompletion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	launchctlcompletion_cmd_direct := exec.Command("./binary")
	err = launchctlcompletion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

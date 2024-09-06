package main

// DarkMode - Control the macOS dark mode from the command-line
// Homepage: https://github.com/sindresorhus/dark-mode

import (
	"fmt"
	
	"os/exec"
)

func installDarkMode() {
	// Método 1: Descargar y extraer .tar.gz
	darkmode_tar_url := "https://github.com/sindresorhus/dark-mode/archive/refs/tags/v3.0.2.tar.gz"
	darkmode_cmd_tar := exec.Command("curl", "-L", darkmode_tar_url, "-o", "package.tar.gz")
	err := darkmode_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	darkmode_zip_url := "https://github.com/sindresorhus/dark-mode/archive/refs/tags/v3.0.2.zip"
	darkmode_cmd_zip := exec.Command("curl", "-L", darkmode_zip_url, "-o", "package.zip")
	err = darkmode_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	darkmode_bin_url := "https://github.com/sindresorhus/dark-mode/archive/refs/tags/v3.0.2.bin"
	darkmode_cmd_bin := exec.Command("curl", "-L", darkmode_bin_url, "-o", "binary.bin")
	err = darkmode_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	darkmode_src_url := "https://github.com/sindresorhus/dark-mode/archive/refs/tags/v3.0.2.src.tar.gz"
	darkmode_cmd_src := exec.Command("curl", "-L", darkmode_src_url, "-o", "source.tar.gz")
	err = darkmode_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	darkmode_cmd_direct := exec.Command("./binary")
	err = darkmode_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

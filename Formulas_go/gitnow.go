package main

// GitNow - Light, temporary commits for git
// Homepage: https://github.com/iwata/git-now

import (
	"fmt"
	
	"os/exec"
)

func installGitNow() {
	// Método 1: Descargar y extraer .tar.gz
	gitnow_tar_url := "https://github.com/iwata/git-now/archive/refs/tags/v0.1.1.0.tar.gz"
	gitnow_cmd_tar := exec.Command("curl", "-L", gitnow_tar_url, "-o", "package.tar.gz")
	err := gitnow_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitnow_zip_url := "https://github.com/iwata/git-now/archive/refs/tags/v0.1.1.0.zip"
	gitnow_cmd_zip := exec.Command("curl", "-L", gitnow_zip_url, "-o", "package.zip")
	err = gitnow_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitnow_bin_url := "https://github.com/iwata/git-now/archive/refs/tags/v0.1.1.0.bin"
	gitnow_cmd_bin := exec.Command("curl", "-L", gitnow_bin_url, "-o", "binary.bin")
	err = gitnow_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitnow_src_url := "https://github.com/iwata/git-now/archive/refs/tags/v0.1.1.0.src.tar.gz"
	gitnow_cmd_src := exec.Command("curl", "-L", gitnow_src_url, "-o", "source.tar.gz")
	err = gitnow_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitnow_cmd_direct := exec.Command("./binary")
	err = gitnow_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gnu-getopt")
exec.Command("latte", "install", "gnu-getopt")
}

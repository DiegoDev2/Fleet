package main

// Gws - Manage workspaces composed of git repositories
// Homepage: https://streakycobra.github.io/gws/

import (
	"fmt"
	
	"os/exec"
)

func installGws() {
	// Método 1: Descargar y extraer .tar.gz
	gws_tar_url := "https://github.com/StreakyCobra/gws/archive/refs/tags/0.2.0.tar.gz"
	gws_cmd_tar := exec.Command("curl", "-L", gws_tar_url, "-o", "package.tar.gz")
	err := gws_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gws_zip_url := "https://github.com/StreakyCobra/gws/archive/refs/tags/0.2.0.zip"
	gws_cmd_zip := exec.Command("curl", "-L", gws_zip_url, "-o", "package.zip")
	err = gws_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gws_bin_url := "https://github.com/StreakyCobra/gws/archive/refs/tags/0.2.0.bin"
	gws_cmd_bin := exec.Command("curl", "-L", gws_bin_url, "-o", "binary.bin")
	err = gws_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gws_src_url := "https://github.com/StreakyCobra/gws/archive/refs/tags/0.2.0.src.tar.gz"
	gws_cmd_src := exec.Command("curl", "-L", gws_src_url, "-o", "source.tar.gz")
	err = gws_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gws_cmd_direct := exec.Command("./binary")
	err = gws_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bash")
exec.Command("latte", "install", "bash")
}

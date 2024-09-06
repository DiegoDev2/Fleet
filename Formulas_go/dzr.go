package main

// Dzr - Command-line Deezer.com player
// Homepage: https://github.com/yne/dzr

import (
	"fmt"
	
	"os/exec"
)

func installDzr() {
	// Método 1: Descargar y extraer .tar.gz
	dzr_tar_url := "https://github.com/yne/dzr/archive/refs/tags/240817.tar.gz"
	dzr_cmd_tar := exec.Command("curl", "-L", dzr_tar_url, "-o", "package.tar.gz")
	err := dzr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dzr_zip_url := "https://github.com/yne/dzr/archive/refs/tags/240817.zip"
	dzr_cmd_zip := exec.Command("curl", "-L", dzr_zip_url, "-o", "package.zip")
	err = dzr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dzr_bin_url := "https://github.com/yne/dzr/archive/refs/tags/240817.bin"
	dzr_cmd_bin := exec.Command("curl", "-L", dzr_bin_url, "-o", "binary.bin")
	err = dzr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dzr_src_url := "https://github.com/yne/dzr/archive/refs/tags/240817.src.tar.gz"
	dzr_cmd_src := exec.Command("curl", "-L", dzr_src_url, "-o", "source.tar.gz")
	err = dzr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dzr_cmd_direct := exec.Command("./binary")
	err = dzr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: dialog")
exec.Command("latte", "install", "dialog")
	fmt.Println("Instalando dependencia: jq")
exec.Command("latte", "install", "jq")
	fmt.Println("Instalando dependencia: mpv")
exec.Command("latte", "install", "mpv")
}

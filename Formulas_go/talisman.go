package main

// Talisman - Tool to detect and prevent secrets from getting checked in
// Homepage: https://thoughtworks.github.io/talisman/

import (
	"fmt"
	
	"os/exec"
)

func installTalisman() {
	// Método 1: Descargar y extraer .tar.gz
	talisman_tar_url := "https://github.com/thoughtworks/talisman/archive/refs/tags/v1.32.0.tar.gz"
	talisman_cmd_tar := exec.Command("curl", "-L", talisman_tar_url, "-o", "package.tar.gz")
	err := talisman_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	talisman_zip_url := "https://github.com/thoughtworks/talisman/archive/refs/tags/v1.32.0.zip"
	talisman_cmd_zip := exec.Command("curl", "-L", talisman_zip_url, "-o", "package.zip")
	err = talisman_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	talisman_bin_url := "https://github.com/thoughtworks/talisman/archive/refs/tags/v1.32.0.bin"
	talisman_cmd_bin := exec.Command("curl", "-L", talisman_bin_url, "-o", "binary.bin")
	err = talisman_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	talisman_src_url := "https://github.com/thoughtworks/talisman/archive/refs/tags/v1.32.0.src.tar.gz"
	talisman_cmd_src := exec.Command("curl", "-L", talisman_src_url, "-o", "source.tar.gz")
	err = talisman_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	talisman_cmd_direct := exec.Command("./binary")
	err = talisman_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}

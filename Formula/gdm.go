package main

// Gdm - Go Dependency Manager (gdm)
// Homepage: https://github.com/sparrc/gdm

import (
	"fmt"
	
	"os/exec"
)

func installGdm() {
	// Método 1: Descargar y extraer .tar.gz
	gdm_tar_url := "https://github.com/sparrc/gdm/archive/refs/tags/1.4.tar.gz"
	gdm_cmd_tar := exec.Command("curl", "-L", gdm_tar_url, "-o", "package.tar.gz")
	err := gdm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gdm_zip_url := "https://github.com/sparrc/gdm/archive/refs/tags/1.4.zip"
	gdm_cmd_zip := exec.Command("curl", "-L", gdm_zip_url, "-o", "package.zip")
	err = gdm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gdm_bin_url := "https://github.com/sparrc/gdm/archive/refs/tags/1.4.bin"
	gdm_cmd_bin := exec.Command("curl", "-L", gdm_bin_url, "-o", "binary.bin")
	err = gdm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gdm_src_url := "https://github.com/sparrc/gdm/archive/refs/tags/1.4.src.tar.gz"
	gdm_cmd_src := exec.Command("curl", "-L", gdm_src_url, "-o", "source.tar.gz")
	err = gdm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gdm_cmd_direct := exec.Command("./binary")
	err = gdm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}

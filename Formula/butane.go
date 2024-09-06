package main

// Butane - Translates human-readable Butane Configs into machine-readable Ignition Configs
// Homepage: https://github.com/coreos/butane

import (
	"fmt"
	
	"os/exec"
)

func installButane() {
	// Método 1: Descargar y extraer .tar.gz
	butane_tar_url := "https://github.com/coreos/butane/archive/refs/tags/v0.21.0.tar.gz"
	butane_cmd_tar := exec.Command("curl", "-L", butane_tar_url, "-o", "package.tar.gz")
	err := butane_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	butane_zip_url := "https://github.com/coreos/butane/archive/refs/tags/v0.21.0.zip"
	butane_cmd_zip := exec.Command("curl", "-L", butane_zip_url, "-o", "package.zip")
	err = butane_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	butane_bin_url := "https://github.com/coreos/butane/archive/refs/tags/v0.21.0.bin"
	butane_cmd_bin := exec.Command("curl", "-L", butane_bin_url, "-o", "binary.bin")
	err = butane_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	butane_src_url := "https://github.com/coreos/butane/archive/refs/tags/v0.21.0.src.tar.gz"
	butane_cmd_src := exec.Command("curl", "-L", butane_src_url, "-o", "source.tar.gz")
	err = butane_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	butane_cmd_direct := exec.Command("./binary")
	err = butane_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}

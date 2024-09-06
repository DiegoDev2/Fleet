package main

// Wapm - WebAssembly Package Manager (CLI)
// Homepage: https://wapm.io/

import (
	"fmt"
	
	"os/exec"
)

func installWapm() {
	// Método 1: Descargar y extraer .tar.gz
	wapm_tar_url := "https://github.com/wasmerio/wapm-cli/archive/refs/tags/v0.5.9.tar.gz"
	wapm_cmd_tar := exec.Command("curl", "-L", wapm_tar_url, "-o", "package.tar.gz")
	err := wapm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wapm_zip_url := "https://github.com/wasmerio/wapm-cli/archive/refs/tags/v0.5.9.zip"
	wapm_cmd_zip := exec.Command("curl", "-L", wapm_zip_url, "-o", "package.zip")
	err = wapm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wapm_bin_url := "https://github.com/wasmerio/wapm-cli/archive/refs/tags/v0.5.9.bin"
	wapm_cmd_bin := exec.Command("curl", "-L", wapm_bin_url, "-o", "binary.bin")
	err = wapm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wapm_src_url := "https://github.com/wasmerio/wapm-cli/archive/refs/tags/v0.5.9.src.tar.gz"
	wapm_cmd_src := exec.Command("curl", "-L", wapm_src_url, "-o", "source.tar.gz")
	err = wapm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wapm_cmd_direct := exec.Command("./binary")
	err = wapm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: wasmer")
exec.Command("latte", "install", "wasmer")
}

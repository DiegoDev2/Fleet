package main

// Dvm - Docker Version Manager
// Homepage: https://github.com/howtowhale/dvm

import (
	"fmt"
	
	"os/exec"
)

func installDvm() {
	// Método 1: Descargar y extraer .tar.gz
	dvm_tar_url := "https://github.com/howtowhale/dvm/archive/refs/tags/1.0.3.tar.gz"
	dvm_cmd_tar := exec.Command("curl", "-L", dvm_tar_url, "-o", "package.tar.gz")
	err := dvm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dvm_zip_url := "https://github.com/howtowhale/dvm/archive/refs/tags/1.0.3.zip"
	dvm_cmd_zip := exec.Command("curl", "-L", dvm_zip_url, "-o", "package.zip")
	err = dvm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dvm_bin_url := "https://github.com/howtowhale/dvm/archive/refs/tags/1.0.3.bin"
	dvm_cmd_bin := exec.Command("curl", "-L", dvm_bin_url, "-o", "binary.bin")
	err = dvm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dvm_src_url := "https://github.com/howtowhale/dvm/archive/refs/tags/1.0.3.src.tar.gz"
	dvm_cmd_src := exec.Command("curl", "-L", dvm_src_url, "-o", "source.tar.gz")
	err = dvm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dvm_cmd_direct := exec.Command("./binary")
	err = dvm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}

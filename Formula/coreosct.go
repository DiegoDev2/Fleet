package main

// CoreosCt - Convert a Container Linux Config into Ignition
// Homepage: https://flatcar-linux.org/docs/latest/provisioning/config-transpiler/

import (
	"fmt"
	
	"os/exec"
)

func installCoreosCt() {
	// Método 1: Descargar y extraer .tar.gz
	coreosct_tar_url := "https://github.com/flatcar/container-linux-config-transpiler/archive/refs/tags/v0.9.4.tar.gz"
	coreosct_cmd_tar := exec.Command("curl", "-L", coreosct_tar_url, "-o", "package.tar.gz")
	err := coreosct_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	coreosct_zip_url := "https://github.com/flatcar/container-linux-config-transpiler/archive/refs/tags/v0.9.4.zip"
	coreosct_cmd_zip := exec.Command("curl", "-L", coreosct_zip_url, "-o", "package.zip")
	err = coreosct_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	coreosct_bin_url := "https://github.com/flatcar/container-linux-config-transpiler/archive/refs/tags/v0.9.4.bin"
	coreosct_cmd_bin := exec.Command("curl", "-L", coreosct_bin_url, "-o", "binary.bin")
	err = coreosct_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	coreosct_src_url := "https://github.com/flatcar/container-linux-config-transpiler/archive/refs/tags/v0.9.4.src.tar.gz"
	coreosct_cmd_src := exec.Command("curl", "-L", coreosct_src_url, "-o", "source.tar.gz")
	err = coreosct_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	coreosct_cmd_direct := exec.Command("./binary")
	err = coreosct_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}

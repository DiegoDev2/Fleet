package main

// Wskdeploy - Apache OpenWhisk project deployment utility
// Homepage: https://openwhisk.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installWskdeploy() {
	// Método 1: Descargar y extraer .tar.gz
	wskdeploy_tar_url := "https://github.com/apache/openwhisk-wskdeploy/archive/refs/tags/1.2.0.tar.gz"
	wskdeploy_cmd_tar := exec.Command("curl", "-L", wskdeploy_tar_url, "-o", "package.tar.gz")
	err := wskdeploy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wskdeploy_zip_url := "https://github.com/apache/openwhisk-wskdeploy/archive/refs/tags/1.2.0.zip"
	wskdeploy_cmd_zip := exec.Command("curl", "-L", wskdeploy_zip_url, "-o", "package.zip")
	err = wskdeploy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wskdeploy_bin_url := "https://github.com/apache/openwhisk-wskdeploy/archive/refs/tags/1.2.0.bin"
	wskdeploy_cmd_bin := exec.Command("curl", "-L", wskdeploy_bin_url, "-o", "binary.bin")
	err = wskdeploy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wskdeploy_src_url := "https://github.com/apache/openwhisk-wskdeploy/archive/refs/tags/1.2.0.src.tar.gz"
	wskdeploy_cmd_src := exec.Command("curl", "-L", wskdeploy_src_url, "-o", "source.tar.gz")
	err = wskdeploy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wskdeploy_cmd_direct := exec.Command("./binary")
	err = wskdeploy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}

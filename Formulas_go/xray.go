package main

// Xray - Platform for building proxies to bypass network restrictions
// Homepage: https://xtls.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installXray() {
	// Método 1: Descargar y extraer .tar.gz
	xray_tar_url := "https://github.com/XTLS/Xray-core/archive/refs/tags/v1.8.24.tar.gz"
	xray_cmd_tar := exec.Command("curl", "-L", xray_tar_url, "-o", "package.tar.gz")
	err := xray_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xray_zip_url := "https://github.com/XTLS/Xray-core/archive/refs/tags/v1.8.24.zip"
	xray_cmd_zip := exec.Command("curl", "-L", xray_zip_url, "-o", "package.zip")
	err = xray_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xray_bin_url := "https://github.com/XTLS/Xray-core/archive/refs/tags/v1.8.24.bin"
	xray_cmd_bin := exec.Command("curl", "-L", xray_bin_url, "-o", "binary.bin")
	err = xray_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xray_src_url := "https://github.com/XTLS/Xray-core/archive/refs/tags/v1.8.24.src.tar.gz"
	xray_cmd_src := exec.Command("curl", "-L", xray_src_url, "-o", "source.tar.gz")
	err = xray_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xray_cmd_direct := exec.Command("./binary")
	err = xray_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}

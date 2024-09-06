package main

// XbeeComm - XBee communication libraries and utilities
// Homepage: https://github.com/guyzmo/xbee-comm

import (
	"fmt"
	
	"os/exec"
)

func installXbeeComm() {
	// Método 1: Descargar y extraer .tar.gz
	xbeecomm_tar_url := "https://github.com/guyzmo/xbee-comm/archive/refs/tags/v1.5.tar.gz"
	xbeecomm_cmd_tar := exec.Command("curl", "-L", xbeecomm_tar_url, "-o", "package.tar.gz")
	err := xbeecomm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xbeecomm_zip_url := "https://github.com/guyzmo/xbee-comm/archive/refs/tags/v1.5.zip"
	xbeecomm_cmd_zip := exec.Command("curl", "-L", xbeecomm_zip_url, "-o", "package.zip")
	err = xbeecomm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xbeecomm_bin_url := "https://github.com/guyzmo/xbee-comm/archive/refs/tags/v1.5.bin"
	xbeecomm_cmd_bin := exec.Command("curl", "-L", xbeecomm_bin_url, "-o", "binary.bin")
	err = xbeecomm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xbeecomm_src_url := "https://github.com/guyzmo/xbee-comm/archive/refs/tags/v1.5.src.tar.gz"
	xbeecomm_cmd_src := exec.Command("curl", "-L", xbeecomm_src_url, "-o", "source.tar.gz")
	err = xbeecomm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xbeecomm_cmd_direct := exec.Command("./binary")
	err = xbeecomm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
}

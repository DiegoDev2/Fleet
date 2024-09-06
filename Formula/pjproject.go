package main

// Pjproject - C library for multimedia protocols such as SIP, SDP, RTP and more
// Homepage: https://www.pjsip.org/

import (
	"fmt"
	
	"os/exec"
)

func installPjproject() {
	// Método 1: Descargar y extraer .tar.gz
	pjproject_tar_url := "https://github.com/pjsip/pjproject/archive/refs/tags/2.14.1.tar.gz"
	pjproject_cmd_tar := exec.Command("curl", "-L", pjproject_tar_url, "-o", "package.tar.gz")
	err := pjproject_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pjproject_zip_url := "https://github.com/pjsip/pjproject/archive/refs/tags/2.14.1.zip"
	pjproject_cmd_zip := exec.Command("curl", "-L", pjproject_zip_url, "-o", "package.zip")
	err = pjproject_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pjproject_bin_url := "https://github.com/pjsip/pjproject/archive/refs/tags/2.14.1.bin"
	pjproject_cmd_bin := exec.Command("curl", "-L", pjproject_bin_url, "-o", "binary.bin")
	err = pjproject_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pjproject_src_url := "https://github.com/pjsip/pjproject/archive/refs/tags/2.14.1.src.tar.gz"
	pjproject_cmd_src := exec.Command("curl", "-L", pjproject_src_url, "-o", "source.tar.gz")
	err = pjproject_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pjproject_cmd_direct := exec.Command("./binary")
	err = pjproject_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}

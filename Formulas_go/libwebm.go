package main

// Libwebm - WebM container
// Homepage: https://www.webmproject.org/code/

import (
	"fmt"
	
	"os/exec"
)

func installLibwebm() {
	// Método 1: Descargar y extraer .tar.gz
	libwebm_tar_url := "https://github.com/webmproject/libwebm/archive/refs/tags/libwebm-1.0.0.31.tar.gz"
	libwebm_cmd_tar := exec.Command("curl", "-L", libwebm_tar_url, "-o", "package.tar.gz")
	err := libwebm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libwebm_zip_url := "https://github.com/webmproject/libwebm/archive/refs/tags/libwebm-1.0.0.31.zip"
	libwebm_cmd_zip := exec.Command("curl", "-L", libwebm_zip_url, "-o", "package.zip")
	err = libwebm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libwebm_bin_url := "https://github.com/webmproject/libwebm/archive/refs/tags/libwebm-1.0.0.31.bin"
	libwebm_cmd_bin := exec.Command("curl", "-L", libwebm_bin_url, "-o", "binary.bin")
	err = libwebm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libwebm_src_url := "https://github.com/webmproject/libwebm/archive/refs/tags/libwebm-1.0.0.31.src.tar.gz"
	libwebm_cmd_src := exec.Command("curl", "-L", libwebm_src_url, "-o", "source.tar.gz")
	err = libwebm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libwebm_cmd_direct := exec.Command("./binary")
	err = libwebm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}

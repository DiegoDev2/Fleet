package main

// H2o - HTTP server with support for HTTP/1.x and HTTP/2
// Homepage: https://github.com/h2o/h2o/

import (
	"fmt"
	
	"os/exec"
)

func installH2o() {
	// Método 1: Descargar y extraer .tar.gz
	h2o_tar_url := "https://github.com/h2o/h2o/archive/refs/tags/v2.2.6.tar.gz"
	h2o_cmd_tar := exec.Command("curl", "-L", h2o_tar_url, "-o", "package.tar.gz")
	err := h2o_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	h2o_zip_url := "https://github.com/h2o/h2o/archive/refs/tags/v2.2.6.zip"
	h2o_cmd_zip := exec.Command("curl", "-L", h2o_zip_url, "-o", "package.zip")
	err = h2o_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	h2o_bin_url := "https://github.com/h2o/h2o/archive/refs/tags/v2.2.6.bin"
	h2o_cmd_bin := exec.Command("curl", "-L", h2o_bin_url, "-o", "binary.bin")
	err = h2o_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	h2o_src_url := "https://github.com/h2o/h2o/archive/refs/tags/v2.2.6.src.tar.gz"
	h2o_cmd_src := exec.Command("curl", "-L", h2o_src_url, "-o", "source.tar.gz")
	err = h2o_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	h2o_cmd_direct := exec.Command("./binary")
	err = h2o_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}

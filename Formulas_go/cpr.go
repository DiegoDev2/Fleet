package main

// Cpr - C++ Requests, a spiritual port of Python Requests
// Homepage: https://docs.libcpr.org/

import (
	"fmt"
	
	"os/exec"
)

func installCpr() {
	// Método 1: Descargar y extraer .tar.gz
	cpr_tar_url := "https://github.com/libcpr/cpr/archive/refs/tags/1.10.5.tar.gz"
	cpr_cmd_tar := exec.Command("curl", "-L", cpr_tar_url, "-o", "package.tar.gz")
	err := cpr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cpr_zip_url := "https://github.com/libcpr/cpr/archive/refs/tags/1.10.5.zip"
	cpr_cmd_zip := exec.Command("curl", "-L", cpr_zip_url, "-o", "package.zip")
	err = cpr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cpr_bin_url := "https://github.com/libcpr/cpr/archive/refs/tags/1.10.5.bin"
	cpr_cmd_bin := exec.Command("curl", "-L", cpr_bin_url, "-o", "binary.bin")
	err = cpr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cpr_src_url := "https://github.com/libcpr/cpr/archive/refs/tags/1.10.5.src.tar.gz"
	cpr_cmd_src := exec.Command("curl", "-L", cpr_src_url, "-o", "source.tar.gz")
	err = cpr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cpr_cmd_direct := exec.Command("./binary")
	err = cpr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}

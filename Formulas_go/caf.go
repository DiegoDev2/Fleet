package main

// Caf - Implementation of the Actor Model for C++
// Homepage: https://www.actor-framework.org/

import (
	"fmt"
	
	"os/exec"
)

func installCaf() {
	// Método 1: Descargar y extraer .tar.gz
	caf_tar_url := "https://github.com/actor-framework/actor-framework/archive/refs/tags/1.0.1.tar.gz"
	caf_cmd_tar := exec.Command("curl", "-L", caf_tar_url, "-o", "package.tar.gz")
	err := caf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	caf_zip_url := "https://github.com/actor-framework/actor-framework/archive/refs/tags/1.0.1.zip"
	caf_cmd_zip := exec.Command("curl", "-L", caf_zip_url, "-o", "package.zip")
	err = caf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	caf_bin_url := "https://github.com/actor-framework/actor-framework/archive/refs/tags/1.0.1.bin"
	caf_cmd_bin := exec.Command("curl", "-L", caf_bin_url, "-o", "binary.bin")
	err = caf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	caf_src_url := "https://github.com/actor-framework/actor-framework/archive/refs/tags/1.0.1.src.tar.gz"
	caf_cmd_src := exec.Command("curl", "-L", caf_src_url, "-o", "source.tar.gz")
	err = caf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	caf_cmd_direct := exec.Command("./binary")
	err = caf_cmd_direct.Run()
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

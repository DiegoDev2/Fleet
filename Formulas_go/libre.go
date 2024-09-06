package main

// Libre - Toolkit library for asynchronous network I/O with protocol stacks
// Homepage: https://github.com/baresip/re

import (
	"fmt"
	
	"os/exec"
)

func installLibre() {
	// Método 1: Descargar y extraer .tar.gz
	libre_tar_url := "https://github.com/baresip/re/archive/refs/tags/v3.15.0.tar.gz"
	libre_cmd_tar := exec.Command("curl", "-L", libre_tar_url, "-o", "package.tar.gz")
	err := libre_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libre_zip_url := "https://github.com/baresip/re/archive/refs/tags/v3.15.0.zip"
	libre_cmd_zip := exec.Command("curl", "-L", libre_zip_url, "-o", "package.zip")
	err = libre_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libre_bin_url := "https://github.com/baresip/re/archive/refs/tags/v3.15.0.bin"
	libre_cmd_bin := exec.Command("curl", "-L", libre_bin_url, "-o", "binary.bin")
	err = libre_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libre_src_url := "https://github.com/baresip/re/archive/refs/tags/v3.15.0.src.tar.gz"
	libre_cmd_src := exec.Command("curl", "-L", libre_src_url, "-o", "source.tar.gz")
	err = libre_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libre_cmd_direct := exec.Command("./binary")
	err = libre_cmd_direct.Run()
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

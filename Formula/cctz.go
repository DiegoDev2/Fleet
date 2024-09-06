package main

// Cctz - C++ library for translating between absolute and civil times
// Homepage: https://github.com/google/cctz

import (
	"fmt"
	
	"os/exec"
)

func installCctz() {
	// Método 1: Descargar y extraer .tar.gz
	cctz_tar_url := "https://github.com/google/cctz/archive/refs/tags/v2.4.tar.gz"
	cctz_cmd_tar := exec.Command("curl", "-L", cctz_tar_url, "-o", "package.tar.gz")
	err := cctz_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cctz_zip_url := "https://github.com/google/cctz/archive/refs/tags/v2.4.zip"
	cctz_cmd_zip := exec.Command("curl", "-L", cctz_zip_url, "-o", "package.zip")
	err = cctz_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cctz_bin_url := "https://github.com/google/cctz/archive/refs/tags/v2.4.bin"
	cctz_cmd_bin := exec.Command("curl", "-L", cctz_bin_url, "-o", "binary.bin")
	err = cctz_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cctz_src_url := "https://github.com/google/cctz/archive/refs/tags/v2.4.src.tar.gz"
	cctz_cmd_src := exec.Command("curl", "-L", cctz_src_url, "-o", "source.tar.gz")
	err = cctz_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cctz_cmd_direct := exec.Command("./binary")
	err = cctz_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}

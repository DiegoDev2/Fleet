package main

// Cmrc - CMake Resource Compiler
// Homepage: https://github.com/vector-of-bool/cmrc

import (
	"fmt"
	
	"os/exec"
)

func installCmrc() {
	// Método 1: Descargar y extraer .tar.gz
	cmrc_tar_url := "https://github.com/vector-of-bool/cmrc/archive/refs/tags/2.0.1.tar.gz"
	cmrc_cmd_tar := exec.Command("curl", "-L", cmrc_tar_url, "-o", "package.tar.gz")
	err := cmrc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cmrc_zip_url := "https://github.com/vector-of-bool/cmrc/archive/refs/tags/2.0.1.zip"
	cmrc_cmd_zip := exec.Command("curl", "-L", cmrc_zip_url, "-o", "package.zip")
	err = cmrc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cmrc_bin_url := "https://github.com/vector-of-bool/cmrc/archive/refs/tags/2.0.1.bin"
	cmrc_cmd_bin := exec.Command("curl", "-L", cmrc_bin_url, "-o", "binary.bin")
	err = cmrc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cmrc_src_url := "https://github.com/vector-of-bool/cmrc/archive/refs/tags/2.0.1.src.tar.gz"
	cmrc_cmd_src := exec.Command("curl", "-L", cmrc_src_url, "-o", "source.tar.gz")
	err = cmrc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cmrc_cmd_direct := exec.Command("./binary")
	err = cmrc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}

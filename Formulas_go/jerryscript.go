package main

// Jerryscript - Ultra-lightweight JavaScript engine for the Internet of Things
// Homepage: https://jerryscript.net

import (
	"fmt"
	
	"os/exec"
)

func installJerryscript() {
	// Método 1: Descargar y extraer .tar.gz
	jerryscript_tar_url := "https://github.com/jerryscript-project/jerryscript/archive/refs/tags/v2.4.0.tar.gz"
	jerryscript_cmd_tar := exec.Command("curl", "-L", jerryscript_tar_url, "-o", "package.tar.gz")
	err := jerryscript_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jerryscript_zip_url := "https://github.com/jerryscript-project/jerryscript/archive/refs/tags/v2.4.0.zip"
	jerryscript_cmd_zip := exec.Command("curl", "-L", jerryscript_zip_url, "-o", "package.zip")
	err = jerryscript_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jerryscript_bin_url := "https://github.com/jerryscript-project/jerryscript/archive/refs/tags/v2.4.0.bin"
	jerryscript_cmd_bin := exec.Command("curl", "-L", jerryscript_bin_url, "-o", "binary.bin")
	err = jerryscript_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jerryscript_src_url := "https://github.com/jerryscript-project/jerryscript/archive/refs/tags/v2.4.0.src.tar.gz"
	jerryscript_cmd_src := exec.Command("curl", "-L", jerryscript_src_url, "-o", "source.tar.gz")
	err = jerryscript_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jerryscript_cmd_direct := exec.Command("./binary")
	err = jerryscript_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}

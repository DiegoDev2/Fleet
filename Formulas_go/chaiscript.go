package main

// Chaiscript - Easy to use embedded scripting language for C++
// Homepage: https://chaiscript.com/

import (
	"fmt"
	
	"os/exec"
)

func installChaiscript() {
	// Método 1: Descargar y extraer .tar.gz
	chaiscript_tar_url := "https://github.com/ChaiScript/ChaiScript/archive/refs/tags/v6.1.0.tar.gz"
	chaiscript_cmd_tar := exec.Command("curl", "-L", chaiscript_tar_url, "-o", "package.tar.gz")
	err := chaiscript_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chaiscript_zip_url := "https://github.com/ChaiScript/ChaiScript/archive/refs/tags/v6.1.0.zip"
	chaiscript_cmd_zip := exec.Command("curl", "-L", chaiscript_zip_url, "-o", "package.zip")
	err = chaiscript_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chaiscript_bin_url := "https://github.com/ChaiScript/ChaiScript/archive/refs/tags/v6.1.0.bin"
	chaiscript_cmd_bin := exec.Command("curl", "-L", chaiscript_bin_url, "-o", "binary.bin")
	err = chaiscript_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chaiscript_src_url := "https://github.com/ChaiScript/ChaiScript/archive/refs/tags/v6.1.0.src.tar.gz"
	chaiscript_cmd_src := exec.Command("curl", "-L", chaiscript_src_url, "-o", "source.tar.gz")
	err = chaiscript_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chaiscript_cmd_direct := exec.Command("./binary")
	err = chaiscript_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}

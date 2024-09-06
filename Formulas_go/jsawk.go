package main

// Jsawk - Like awk, but for JSON, using JavaScript objects and arrays
// Homepage: https://github.com/micha/jsawk

import (
	"fmt"
	
	"os/exec"
)

func installJsawk() {
	// Método 1: Descargar y extraer .tar.gz
	jsawk_tar_url := "https://github.com/micha/jsawk/archive/refs/tags/1.4.tar.gz"
	jsawk_cmd_tar := exec.Command("curl", "-L", jsawk_tar_url, "-o", "package.tar.gz")
	err := jsawk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jsawk_zip_url := "https://github.com/micha/jsawk/archive/refs/tags/1.4.zip"
	jsawk_cmd_zip := exec.Command("curl", "-L", jsawk_zip_url, "-o", "package.zip")
	err = jsawk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jsawk_bin_url := "https://github.com/micha/jsawk/archive/refs/tags/1.4.bin"
	jsawk_cmd_bin := exec.Command("curl", "-L", jsawk_bin_url, "-o", "binary.bin")
	err = jsawk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jsawk_src_url := "https://github.com/micha/jsawk/archive/refs/tags/1.4.src.tar.gz"
	jsawk_cmd_src := exec.Command("curl", "-L", jsawk_src_url, "-o", "source.tar.gz")
	err = jsawk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jsawk_cmd_direct := exec.Command("./binary")
	err = jsawk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: spidermonkey")
exec.Command("latte", "install", "spidermonkey")
}

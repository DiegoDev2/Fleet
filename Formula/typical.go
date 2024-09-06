package main

// Typical - Data interchange with algebraic data types
// Homepage: https://github.com/stepchowfun/typical

import (
	"fmt"
	
	"os/exec"
)

func installTypical() {
	// Método 1: Descargar y extraer .tar.gz
	typical_tar_url := "https://github.com/stepchowfun/typical/archive/refs/tags/v0.12.1.tar.gz"
	typical_cmd_tar := exec.Command("curl", "-L", typical_tar_url, "-o", "package.tar.gz")
	err := typical_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	typical_zip_url := "https://github.com/stepchowfun/typical/archive/refs/tags/v0.12.1.zip"
	typical_cmd_zip := exec.Command("curl", "-L", typical_zip_url, "-o", "package.zip")
	err = typical_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	typical_bin_url := "https://github.com/stepchowfun/typical/archive/refs/tags/v0.12.1.bin"
	typical_cmd_bin := exec.Command("curl", "-L", typical_bin_url, "-o", "binary.bin")
	err = typical_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	typical_src_url := "https://github.com/stepchowfun/typical/archive/refs/tags/v0.12.1.src.tar.gz"
	typical_cmd_src := exec.Command("curl", "-L", typical_src_url, "-o", "source.tar.gz")
	err = typical_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	typical_cmd_direct := exec.Command("./binary")
	err = typical_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}

package main

// Form - Symbolic manipulation system
// Homepage: https://www.nikhef.nl/~form/

import (
	"fmt"
	
	"os/exec"
)

func installForm() {
	// Método 1: Descargar y extraer .tar.gz
	form_tar_url := "https://github.com/vermaseren/form/releases/download/v4.3.1/form-4.3.1.tar.gz"
	form_cmd_tar := exec.Command("curl", "-L", form_tar_url, "-o", "package.tar.gz")
	err := form_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	form_zip_url := "https://github.com/vermaseren/form/releases/download/v4.3.1/form-4.3.1.zip"
	form_cmd_zip := exec.Command("curl", "-L", form_zip_url, "-o", "package.zip")
	err = form_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	form_bin_url := "https://github.com/vermaseren/form/releases/download/v4.3.1/form-4.3.1.bin"
	form_cmd_bin := exec.Command("curl", "-L", form_bin_url, "-o", "binary.bin")
	err = form_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	form_src_url := "https://github.com/vermaseren/form/releases/download/v4.3.1/form-4.3.1.src.tar.gz"
	form_cmd_src := exec.Command("curl", "-L", form_src_url, "-o", "source.tar.gz")
	err = form_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	form_cmd_direct := exec.Command("./binary")
	err = form_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
}

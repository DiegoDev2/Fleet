package main

// Oil - Bash-compatible Unix shell with more consistent syntax and semantics
// Homepage: https://www.oilshell.org/

import (
	"fmt"
	
	"os/exec"
)

func installOil() {
	// Método 1: Descargar y extraer .tar.gz
	oil_tar_url := "https://www.oilshell.org/download/oil-0.23.0.tar.gz"
	oil_cmd_tar := exec.Command("curl", "-L", oil_tar_url, "-o", "package.tar.gz")
	err := oil_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	oil_zip_url := "https://www.oilshell.org/download/oil-0.23.0.zip"
	oil_cmd_zip := exec.Command("curl", "-L", oil_zip_url, "-o", "package.zip")
	err = oil_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	oil_bin_url := "https://www.oilshell.org/download/oil-0.23.0.bin"
	oil_cmd_bin := exec.Command("curl", "-L", oil_bin_url, "-o", "binary.bin")
	err = oil_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	oil_src_url := "https://www.oilshell.org/download/oil-0.23.0.src.tar.gz"
	oil_cmd_src := exec.Command("curl", "-L", oil_src_url, "-o", "source.tar.gz")
	err = oil_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	oil_cmd_direct := exec.Command("./binary")
	err = oil_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}

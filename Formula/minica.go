package main

// Minica - Small, simple certificate authority
// Homepage: https://github.com/jsha/minica

import (
	"fmt"
	
	"os/exec"
)

func installMinica() {
	// Método 1: Descargar y extraer .tar.gz
	minica_tar_url := "https://github.com/jsha/minica/archive/refs/tags/v1.1.0.tar.gz"
	minica_cmd_tar := exec.Command("curl", "-L", minica_tar_url, "-o", "package.tar.gz")
	err := minica_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	minica_zip_url := "https://github.com/jsha/minica/archive/refs/tags/v1.1.0.zip"
	minica_cmd_zip := exec.Command("curl", "-L", minica_zip_url, "-o", "package.zip")
	err = minica_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	minica_bin_url := "https://github.com/jsha/minica/archive/refs/tags/v1.1.0.bin"
	minica_cmd_bin := exec.Command("curl", "-L", minica_bin_url, "-o", "binary.bin")
	err = minica_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	minica_src_url := "https://github.com/jsha/minica/archive/refs/tags/v1.1.0.src.tar.gz"
	minica_cmd_src := exec.Command("curl", "-L", minica_src_url, "-o", "source.tar.gz")
	err = minica_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	minica_cmd_direct := exec.Command("./binary")
	err = minica_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}

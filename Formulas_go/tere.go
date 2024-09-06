package main

// Tere - Terminal file explorer
// Homepage: https://github.com/mgunyho/tere

import (
	"fmt"
	
	"os/exec"
)

func installTere() {
	// Método 1: Descargar y extraer .tar.gz
	tere_tar_url := "https://github.com/mgunyho/tere/archive/refs/tags/v1.5.1.tar.gz"
	tere_cmd_tar := exec.Command("curl", "-L", tere_tar_url, "-o", "package.tar.gz")
	err := tere_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tere_zip_url := "https://github.com/mgunyho/tere/archive/refs/tags/v1.5.1.zip"
	tere_cmd_zip := exec.Command("curl", "-L", tere_zip_url, "-o", "package.zip")
	err = tere_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tere_bin_url := "https://github.com/mgunyho/tere/archive/refs/tags/v1.5.1.bin"
	tere_cmd_bin := exec.Command("curl", "-L", tere_bin_url, "-o", "binary.bin")
	err = tere_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tere_src_url := "https://github.com/mgunyho/tere/archive/refs/tags/v1.5.1.src.tar.gz"
	tere_cmd_src := exec.Command("curl", "-L", tere_src_url, "-o", "source.tar.gz")
	err = tere_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tere_cmd_direct := exec.Command("./binary")
	err = tere_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}

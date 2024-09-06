package main

// Ares - Automated decoding of encrypted text
// Homepage: https://github.com/bee-san/Ares

import (
	"fmt"
	
	"os/exec"
)

func installAres() {
	// Método 1: Descargar y extraer .tar.gz
	ares_tar_url := "https://github.com/bee-san/Ares/archive/refs/tags/v0.10.0.tar.gz"
	ares_cmd_tar := exec.Command("curl", "-L", ares_tar_url, "-o", "package.tar.gz")
	err := ares_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ares_zip_url := "https://github.com/bee-san/Ares/archive/refs/tags/v0.10.0.zip"
	ares_cmd_zip := exec.Command("curl", "-L", ares_zip_url, "-o", "package.zip")
	err = ares_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ares_bin_url := "https://github.com/bee-san/Ares/archive/refs/tags/v0.10.0.bin"
	ares_cmd_bin := exec.Command("curl", "-L", ares_bin_url, "-o", "binary.bin")
	err = ares_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ares_src_url := "https://github.com/bee-san/Ares/archive/refs/tags/v0.10.0.src.tar.gz"
	ares_cmd_src := exec.Command("curl", "-L", ares_src_url, "-o", "source.tar.gz")
	err = ares_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ares_cmd_direct := exec.Command("./binary")
	err = ares_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}

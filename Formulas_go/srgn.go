package main

// Srgn - Code surgeon for precise text and code transplantation
// Homepage: https://github.com/alexpovel/srgn

import (
	"fmt"
	
	"os/exec"
)

func installSrgn() {
	// Método 1: Descargar y extraer .tar.gz
	srgn_tar_url := "https://github.com/alexpovel/srgn/archive/refs/tags/srgn-v0.13.1.tar.gz"
	srgn_cmd_tar := exec.Command("curl", "-L", srgn_tar_url, "-o", "package.tar.gz")
	err := srgn_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	srgn_zip_url := "https://github.com/alexpovel/srgn/archive/refs/tags/srgn-v0.13.1.zip"
	srgn_cmd_zip := exec.Command("curl", "-L", srgn_zip_url, "-o", "package.zip")
	err = srgn_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	srgn_bin_url := "https://github.com/alexpovel/srgn/archive/refs/tags/srgn-v0.13.1.bin"
	srgn_cmd_bin := exec.Command("curl", "-L", srgn_bin_url, "-o", "binary.bin")
	err = srgn_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	srgn_src_url := "https://github.com/alexpovel/srgn/archive/refs/tags/srgn-v0.13.1.src.tar.gz"
	srgn_cmd_src := exec.Command("curl", "-L", srgn_src_url, "-o", "source.tar.gz")
	err = srgn_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	srgn_cmd_direct := exec.Command("./binary")
	err = srgn_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}

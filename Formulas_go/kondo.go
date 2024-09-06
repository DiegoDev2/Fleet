package main

// Kondo - Save disk space by cleaning non-essential files from software projects
// Homepage: https://github.com/tbillington/kondo

import (
	"fmt"
	
	"os/exec"
)

func installKondo() {
	// Método 1: Descargar y extraer .tar.gz
	kondo_tar_url := "https://github.com/tbillington/kondo/archive/refs/tags/v0.8.tar.gz"
	kondo_cmd_tar := exec.Command("curl", "-L", kondo_tar_url, "-o", "package.tar.gz")
	err := kondo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kondo_zip_url := "https://github.com/tbillington/kondo/archive/refs/tags/v0.8.zip"
	kondo_cmd_zip := exec.Command("curl", "-L", kondo_zip_url, "-o", "package.zip")
	err = kondo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kondo_bin_url := "https://github.com/tbillington/kondo/archive/refs/tags/v0.8.bin"
	kondo_cmd_bin := exec.Command("curl", "-L", kondo_bin_url, "-o", "binary.bin")
	err = kondo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kondo_src_url := "https://github.com/tbillington/kondo/archive/refs/tags/v0.8.src.tar.gz"
	kondo_cmd_src := exec.Command("curl", "-L", kondo_src_url, "-o", "source.tar.gz")
	err = kondo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kondo_cmd_direct := exec.Command("./binary")
	err = kondo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}

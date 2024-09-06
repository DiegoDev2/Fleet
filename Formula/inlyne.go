package main

// Inlyne - GPU powered yet browserless tool to help you quickly view markdown files
// Homepage: https://github.com/Inlyne-Project/inlyne

import (
	"fmt"
	
	"os/exec"
)

func installInlyne() {
	// Método 1: Descargar y extraer .tar.gz
	inlyne_tar_url := "https://github.com/Inlyne-Project/inlyne/archive/refs/tags/v0.4.3.tar.gz"
	inlyne_cmd_tar := exec.Command("curl", "-L", inlyne_tar_url, "-o", "package.tar.gz")
	err := inlyne_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	inlyne_zip_url := "https://github.com/Inlyne-Project/inlyne/archive/refs/tags/v0.4.3.zip"
	inlyne_cmd_zip := exec.Command("curl", "-L", inlyne_zip_url, "-o", "package.zip")
	err = inlyne_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	inlyne_bin_url := "https://github.com/Inlyne-Project/inlyne/archive/refs/tags/v0.4.3.bin"
	inlyne_cmd_bin := exec.Command("curl", "-L", inlyne_bin_url, "-o", "binary.bin")
	err = inlyne_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	inlyne_src_url := "https://github.com/Inlyne-Project/inlyne/archive/refs/tags/v0.4.3.src.tar.gz"
	inlyne_cmd_src := exec.Command("curl", "-L", inlyne_src_url, "-o", "source.tar.gz")
	err = inlyne_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	inlyne_cmd_direct := exec.Command("./binary")
	err = inlyne_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}

package main

// Mask - CLI task runner defined by a simple markdown file
// Homepage: https://github.com/jakedeichert/mask/

import (
	"fmt"
	
	"os/exec"
)

func installMask() {
	// Método 1: Descargar y extraer .tar.gz
	mask_tar_url := "https://github.com/jacobdeichert/mask/archive/refs/tags/mask/0.11.4.tar.gz"
	mask_cmd_tar := exec.Command("curl", "-L", mask_tar_url, "-o", "package.tar.gz")
	err := mask_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mask_zip_url := "https://github.com/jacobdeichert/mask/archive/refs/tags/mask/0.11.4.zip"
	mask_cmd_zip := exec.Command("curl", "-L", mask_zip_url, "-o", "package.zip")
	err = mask_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mask_bin_url := "https://github.com/jacobdeichert/mask/archive/refs/tags/mask/0.11.4.bin"
	mask_cmd_bin := exec.Command("curl", "-L", mask_bin_url, "-o", "binary.bin")
	err = mask_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mask_src_url := "https://github.com/jacobdeichert/mask/archive/refs/tags/mask/0.11.4.src.tar.gz"
	mask_cmd_src := exec.Command("curl", "-L", mask_src_url, "-o", "source.tar.gz")
	err = mask_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mask_cmd_direct := exec.Command("./binary")
	err = mask_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}

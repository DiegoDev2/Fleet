package main

// FileFormula - Utility to determine file types
// Homepage: https://darwinsys.com/file/

import (
	"fmt"
	
	"os/exec"
)

func installFileFormula() {
	// Método 1: Descargar y extraer .tar.gz
	fileformula_tar_url := "https://astron.com/pub/file/file-5.45.tar.gz"
	fileformula_cmd_tar := exec.Command("curl", "-L", fileformula_tar_url, "-o", "package.tar.gz")
	err := fileformula_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fileformula_zip_url := "https://astron.com/pub/file/file-5.45.zip"
	fileformula_cmd_zip := exec.Command("curl", "-L", fileformula_zip_url, "-o", "package.zip")
	err = fileformula_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fileformula_bin_url := "https://astron.com/pub/file/file-5.45.bin"
	fileformula_cmd_bin := exec.Command("curl", "-L", fileformula_bin_url, "-o", "binary.bin")
	err = fileformula_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fileformula_src_url := "https://astron.com/pub/file/file-5.45.src.tar.gz"
	fileformula_cmd_src := exec.Command("curl", "-L", fileformula_src_url, "-o", "source.tar.gz")
	err = fileformula_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fileformula_cmd_direct := exec.Command("./binary")
	err = fileformula_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libmagic")
	exec.Command("latte", "install", "libmagic").Run()
}

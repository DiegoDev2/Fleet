package main

// Croc - Securely send things from one computer to another
// Homepage: https://github.com/schollz/croc

import (
	"fmt"
	
	"os/exec"
)

func installCroc() {
	// Método 1: Descargar y extraer .tar.gz
	croc_tar_url := "https://github.com/schollz/croc/archive/refs/tags/v10.0.12.tar.gz"
	croc_cmd_tar := exec.Command("curl", "-L", croc_tar_url, "-o", "package.tar.gz")
	err := croc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	croc_zip_url := "https://github.com/schollz/croc/archive/refs/tags/v10.0.12.zip"
	croc_cmd_zip := exec.Command("curl", "-L", croc_zip_url, "-o", "package.zip")
	err = croc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	croc_bin_url := "https://github.com/schollz/croc/archive/refs/tags/v10.0.12.bin"
	croc_cmd_bin := exec.Command("curl", "-L", croc_bin_url, "-o", "binary.bin")
	err = croc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	croc_src_url := "https://github.com/schollz/croc/archive/refs/tags/v10.0.12.src.tar.gz"
	croc_cmd_src := exec.Command("curl", "-L", croc_src_url, "-o", "source.tar.gz")
	err = croc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	croc_cmd_direct := exec.Command("./binary")
	err = croc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}

package main

// Openvi - Portable OpenBSD vi for UNIX systems
// Homepage: https://github.com/johnsonjh/OpenVi

import (
	"fmt"
	
	"os/exec"
)

func installOpenvi() {
	// Método 1: Descargar y extraer .tar.gz
	openvi_tar_url := "https://github.com/johnsonjh/OpenVi/archive/refs/tags/7.5.29.tar.gz"
	openvi_cmd_tar := exec.Command("curl", "-L", openvi_tar_url, "-o", "package.tar.gz")
	err := openvi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openvi_zip_url := "https://github.com/johnsonjh/OpenVi/archive/refs/tags/7.5.29.zip"
	openvi_cmd_zip := exec.Command("curl", "-L", openvi_zip_url, "-o", "package.zip")
	err = openvi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openvi_bin_url := "https://github.com/johnsonjh/OpenVi/archive/refs/tags/7.5.29.bin"
	openvi_cmd_bin := exec.Command("curl", "-L", openvi_bin_url, "-o", "binary.bin")
	err = openvi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openvi_src_url := "https://github.com/johnsonjh/OpenVi/archive/refs/tags/7.5.29.src.tar.gz"
	openvi_cmd_src := exec.Command("curl", "-L", openvi_src_url, "-o", "source.tar.gz")
	err = openvi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openvi_cmd_direct := exec.Command("./binary")
	err = openvi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ncurses")
exec.Command("latte", "install", "ncurses")
}

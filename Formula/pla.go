package main

// Pla - Tool for building Gantt charts in PNG, EPS, PDF or SVG format
// Homepage: https://www.arpalert.org/pla.html

import (
	"fmt"
	
	"os/exec"
)

func installPla() {
	// Método 1: Descargar y extraer .tar.gz
	pla_tar_url := "https://github.com/thierry-f-78/pla/archive/refs/tags/1.3.tar.gz"
	pla_cmd_tar := exec.Command("curl", "-L", pla_tar_url, "-o", "package.tar.gz")
	err := pla_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pla_zip_url := "https://github.com/thierry-f-78/pla/archive/refs/tags/1.3.zip"
	pla_cmd_zip := exec.Command("curl", "-L", pla_zip_url, "-o", "package.zip")
	err = pla_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pla_bin_url := "https://github.com/thierry-f-78/pla/archive/refs/tags/1.3.bin"
	pla_cmd_bin := exec.Command("curl", "-L", pla_bin_url, "-o", "binary.bin")
	err = pla_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pla_src_url := "https://github.com/thierry-f-78/pla/archive/refs/tags/1.3.src.tar.gz"
	pla_cmd_src := exec.Command("curl", "-L", pla_src_url, "-o", "source.tar.gz")
	err = pla_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pla_cmd_direct := exec.Command("./binary")
	err = pla_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
}

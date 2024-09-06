package main

// Riemann - Event stream processor
// Homepage: https://riemann.io/

import (
	"fmt"
	
	"os/exec"
)

func installRiemann() {
	// Método 1: Descargar y extraer .tar.gz
	riemann_tar_url := "https://github.com/riemann/riemann/releases/download/0.3.11/riemann-0.3.11.tar.bz2"
	riemann_cmd_tar := exec.Command("curl", "-L", riemann_tar_url, "-o", "package.tar.gz")
	err := riemann_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	riemann_zip_url := "https://github.com/riemann/riemann/releases/download/0.3.11/riemann-0.3.11.tar.bz2"
	riemann_cmd_zip := exec.Command("curl", "-L", riemann_zip_url, "-o", "package.zip")
	err = riemann_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	riemann_bin_url := "https://github.com/riemann/riemann/releases/download/0.3.11/riemann-0.3.11.tar.bz2"
	riemann_cmd_bin := exec.Command("curl", "-L", riemann_bin_url, "-o", "binary.bin")
	err = riemann_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	riemann_src_url := "https://github.com/riemann/riemann/releases/download/0.3.11/riemann-0.3.11.tar.bz2"
	riemann_cmd_src := exec.Command("curl", "-L", riemann_src_url, "-o", "source.tar.gz")
	err = riemann_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	riemann_cmd_direct := exec.Command("./binary")
	err = riemann_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}

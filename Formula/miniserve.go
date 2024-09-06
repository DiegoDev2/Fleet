package main

// Miniserve - High performance static file server
// Homepage: https://github.com/svenstaro/miniserve

import (
	"fmt"
	
	"os/exec"
)

func installMiniserve() {
	// Método 1: Descargar y extraer .tar.gz
	miniserve_tar_url := "https://github.com/svenstaro/miniserve/archive/refs/tags/v0.27.1.tar.gz"
	miniserve_cmd_tar := exec.Command("curl", "-L", miniserve_tar_url, "-o", "package.tar.gz")
	err := miniserve_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	miniserve_zip_url := "https://github.com/svenstaro/miniserve/archive/refs/tags/v0.27.1.zip"
	miniserve_cmd_zip := exec.Command("curl", "-L", miniserve_zip_url, "-o", "package.zip")
	err = miniserve_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	miniserve_bin_url := "https://github.com/svenstaro/miniserve/archive/refs/tags/v0.27.1.bin"
	miniserve_cmd_bin := exec.Command("curl", "-L", miniserve_bin_url, "-o", "binary.bin")
	err = miniserve_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	miniserve_src_url := "https://github.com/svenstaro/miniserve/archive/refs/tags/v0.27.1.src.tar.gz"
	miniserve_cmd_src := exec.Command("curl", "-L", miniserve_src_url, "-o", "source.tar.gz")
	err = miniserve_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	miniserve_cmd_direct := exec.Command("./binary")
	err = miniserve_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}

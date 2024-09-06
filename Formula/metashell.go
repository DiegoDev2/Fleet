package main

// Metashell - Metaprogramming shell for C++ templates
// Homepage: http://metashell.org

import (
	"fmt"
	
	"os/exec"
)

func installMetashell() {
	// Método 1: Descargar y extraer .tar.gz
	metashell_tar_url := "https://github.com/metashell/metashell/archive/refs/tags/v5.0.0.tar.gz"
	metashell_cmd_tar := exec.Command("curl", "-L", metashell_tar_url, "-o", "package.tar.gz")
	err := metashell_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	metashell_zip_url := "https://github.com/metashell/metashell/archive/refs/tags/v5.0.0.zip"
	metashell_cmd_zip := exec.Command("curl", "-L", metashell_zip_url, "-o", "package.zip")
	err = metashell_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	metashell_bin_url := "https://github.com/metashell/metashell/archive/refs/tags/v5.0.0.bin"
	metashell_cmd_bin := exec.Command("curl", "-L", metashell_bin_url, "-o", "binary.bin")
	err = metashell_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	metashell_src_url := "https://github.com/metashell/metashell/archive/refs/tags/v5.0.0.src.tar.gz"
	metashell_cmd_src := exec.Command("curl", "-L", metashell_src_url, "-o", "source.tar.gz")
	err = metashell_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	metashell_cmd_direct := exec.Command("./binary")
	err = metashell_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}

package main

// Csfml - SMFL bindings for C
// Homepage: https://www.sfml-dev.org/

import (
	"fmt"
	
	"os/exec"
)

func installCsfml() {
	// Método 1: Descargar y extraer .tar.gz
	csfml_tar_url := "https://github.com/SFML/CSFML/archive/refs/tags/2.6.1.tar.gz"
	csfml_cmd_tar := exec.Command("curl", "-L", csfml_tar_url, "-o", "package.tar.gz")
	err := csfml_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	csfml_zip_url := "https://github.com/SFML/CSFML/archive/refs/tags/2.6.1.zip"
	csfml_cmd_zip := exec.Command("curl", "-L", csfml_zip_url, "-o", "package.zip")
	err = csfml_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	csfml_bin_url := "https://github.com/SFML/CSFML/archive/refs/tags/2.6.1.bin"
	csfml_cmd_bin := exec.Command("curl", "-L", csfml_bin_url, "-o", "binary.bin")
	err = csfml_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	csfml_src_url := "https://github.com/SFML/CSFML/archive/refs/tags/2.6.1.src.tar.gz"
	csfml_cmd_src := exec.Command("curl", "-L", csfml_src_url, "-o", "source.tar.gz")
	err = csfml_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	csfml_cmd_direct := exec.Command("./binary")
	err = csfml_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: sfml")
	exec.Command("latte", "install", "sfml").Run()
}

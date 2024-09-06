package main

// Libjuice - UDP Interactive Connectivity Establishment (ICE) library
// Homepage: https://github.com/paullouisageneau/libjuice

import (
	"fmt"
	
	"os/exec"
)

func installLibjuice() {
	// Método 1: Descargar y extraer .tar.gz
	libjuice_tar_url := "https://github.com/paullouisageneau/libjuice/archive/refs/tags/v1.5.3.tar.gz"
	libjuice_cmd_tar := exec.Command("curl", "-L", libjuice_tar_url, "-o", "package.tar.gz")
	err := libjuice_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libjuice_zip_url := "https://github.com/paullouisageneau/libjuice/archive/refs/tags/v1.5.3.zip"
	libjuice_cmd_zip := exec.Command("curl", "-L", libjuice_zip_url, "-o", "package.zip")
	err = libjuice_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libjuice_bin_url := "https://github.com/paullouisageneau/libjuice/archive/refs/tags/v1.5.3.bin"
	libjuice_cmd_bin := exec.Command("curl", "-L", libjuice_bin_url, "-o", "binary.bin")
	err = libjuice_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libjuice_src_url := "https://github.com/paullouisageneau/libjuice/archive/refs/tags/v1.5.3.src.tar.gz"
	libjuice_cmd_src := exec.Command("curl", "-L", libjuice_src_url, "-o", "source.tar.gz")
	err = libjuice_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libjuice_cmd_direct := exec.Command("./binary")
	err = libjuice_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}

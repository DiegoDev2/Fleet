package main

// Scnlib - Scanf for modern C++
// Homepage: https://scnlib.dev

import (
	"fmt"
	
	"os/exec"
)

func installScnlib() {
	// Método 1: Descargar y extraer .tar.gz
	scnlib_tar_url := "https://github.com/eliaskosunen/scnlib/archive/refs/tags/v3.0.1.tar.gz"
	scnlib_cmd_tar := exec.Command("curl", "-L", scnlib_tar_url, "-o", "package.tar.gz")
	err := scnlib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	scnlib_zip_url := "https://github.com/eliaskosunen/scnlib/archive/refs/tags/v3.0.1.zip"
	scnlib_cmd_zip := exec.Command("curl", "-L", scnlib_zip_url, "-o", "package.zip")
	err = scnlib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	scnlib_bin_url := "https://github.com/eliaskosunen/scnlib/archive/refs/tags/v3.0.1.bin"
	scnlib_cmd_bin := exec.Command("curl", "-L", scnlib_bin_url, "-o", "binary.bin")
	err = scnlib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	scnlib_src_url := "https://github.com/eliaskosunen/scnlib/archive/refs/tags/v3.0.1.src.tar.gz"
	scnlib_cmd_src := exec.Command("curl", "-L", scnlib_src_url, "-o", "source.tar.gz")
	err = scnlib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	scnlib_cmd_direct := exec.Command("./binary")
	err = scnlib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: simdutf")
	exec.Command("latte", "install", "simdutf").Run()
}

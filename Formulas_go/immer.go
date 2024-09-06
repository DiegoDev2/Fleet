package main

// Immer - Library of persistent and immutable data structures written in C++
// Homepage: https://sinusoid.es/immer/

import (
	"fmt"
	
	"os/exec"
)

func installImmer() {
	// Método 1: Descargar y extraer .tar.gz
	immer_tar_url := "https://github.com/arximboldi/immer/archive/refs/tags/v0.8.1.tar.gz"
	immer_cmd_tar := exec.Command("curl", "-L", immer_tar_url, "-o", "package.tar.gz")
	err := immer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	immer_zip_url := "https://github.com/arximboldi/immer/archive/refs/tags/v0.8.1.zip"
	immer_cmd_zip := exec.Command("curl", "-L", immer_zip_url, "-o", "package.zip")
	err = immer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	immer_bin_url := "https://github.com/arximboldi/immer/archive/refs/tags/v0.8.1.bin"
	immer_cmd_bin := exec.Command("curl", "-L", immer_bin_url, "-o", "binary.bin")
	err = immer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	immer_src_url := "https://github.com/arximboldi/immer/archive/refs/tags/v0.8.1.src.tar.gz"
	immer_cmd_src := exec.Command("curl", "-L", immer_src_url, "-o", "source.tar.gz")
	err = immer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	immer_cmd_direct := exec.Command("./binary")
	err = immer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}

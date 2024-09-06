package main

// Onednn - Basic building blocks for deep learning applications
// Homepage: https://www.oneapi.io/open-source/

import (
	"fmt"
	
	"os/exec"
)

func installOnednn() {
	// Método 1: Descargar y extraer .tar.gz
	onednn_tar_url := "https://github.com/oneapi-src/oneDNN/archive/refs/tags/v3.5.3.tar.gz"
	onednn_cmd_tar := exec.Command("curl", "-L", onednn_tar_url, "-o", "package.tar.gz")
	err := onednn_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	onednn_zip_url := "https://github.com/oneapi-src/oneDNN/archive/refs/tags/v3.5.3.zip"
	onednn_cmd_zip := exec.Command("curl", "-L", onednn_zip_url, "-o", "package.zip")
	err = onednn_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	onednn_bin_url := "https://github.com/oneapi-src/oneDNN/archive/refs/tags/v3.5.3.bin"
	onednn_cmd_bin := exec.Command("curl", "-L", onednn_bin_url, "-o", "binary.bin")
	err = onednn_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	onednn_src_url := "https://github.com/oneapi-src/oneDNN/archive/refs/tags/v3.5.3.src.tar.gz"
	onednn_cmd_src := exec.Command("curl", "-L", onednn_src_url, "-o", "source.tar.gz")
	err = onednn_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	onednn_cmd_direct := exec.Command("./binary")
	err = onednn_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: doxygen")
exec.Command("latte", "install", "doxygen")
}

package main

// Woff2 - Utilities to create and convert Web Open Font File (WOFF) files
// Homepage: https://github.com/google/woff2

import (
	"fmt"
	
	"os/exec"
)

func installWoff2() {
	// Método 1: Descargar y extraer .tar.gz
	woff2_tar_url := "https://github.com/google/woff2/archive/refs/tags/v1.0.2.tar.gz"
	woff2_cmd_tar := exec.Command("curl", "-L", woff2_tar_url, "-o", "package.tar.gz")
	err := woff2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	woff2_zip_url := "https://github.com/google/woff2/archive/refs/tags/v1.0.2.zip"
	woff2_cmd_zip := exec.Command("curl", "-L", woff2_zip_url, "-o", "package.zip")
	err = woff2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	woff2_bin_url := "https://github.com/google/woff2/archive/refs/tags/v1.0.2.bin"
	woff2_cmd_bin := exec.Command("curl", "-L", woff2_bin_url, "-o", "binary.bin")
	err = woff2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	woff2_src_url := "https://github.com/google/woff2/archive/refs/tags/v1.0.2.src.tar.gz"
	woff2_cmd_src := exec.Command("curl", "-L", woff2_src_url, "-o", "source.tar.gz")
	err = woff2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	woff2_cmd_direct := exec.Command("./binary")
	err = woff2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: brotli")
exec.Command("latte", "install", "brotli")
}

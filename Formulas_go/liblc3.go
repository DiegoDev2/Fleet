package main

// Liblc3 - Low Complexity Communication Codec library and tools
// Homepage: https://github.com/google/liblc3

import (
	"fmt"
	
	"os/exec"
)

func installLiblc3() {
	// Método 1: Descargar y extraer .tar.gz
	liblc3_tar_url := "https://github.com/google/liblc3/archive/refs/tags/v1.1.1.tar.gz"
	liblc3_cmd_tar := exec.Command("curl", "-L", liblc3_tar_url, "-o", "package.tar.gz")
	err := liblc3_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	liblc3_zip_url := "https://github.com/google/liblc3/archive/refs/tags/v1.1.1.zip"
	liblc3_cmd_zip := exec.Command("curl", "-L", liblc3_zip_url, "-o", "package.zip")
	err = liblc3_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	liblc3_bin_url := "https://github.com/google/liblc3/archive/refs/tags/v1.1.1.bin"
	liblc3_cmd_bin := exec.Command("curl", "-L", liblc3_bin_url, "-o", "binary.bin")
	err = liblc3_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	liblc3_src_url := "https://github.com/google/liblc3/archive/refs/tags/v1.1.1.src.tar.gz"
	liblc3_cmd_src := exec.Command("curl", "-L", liblc3_src_url, "-o", "source.tar.gz")
	err = liblc3_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	liblc3_cmd_direct := exec.Command("./binary")
	err = liblc3_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
}

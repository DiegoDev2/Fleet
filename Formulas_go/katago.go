package main

// Katago - Neural Network Go engine with no human-provided knowledge
// Homepage: https://github.com/lightvector/KataGo

import (
	"fmt"
	
	"os/exec"
)

func installKatago() {
	// Método 1: Descargar y extraer .tar.gz
	katago_tar_url := "https://github.com/lightvector/KataGo/archive/refs/tags/v1.15.3.tar.gz"
	katago_cmd_tar := exec.Command("curl", "-L", katago_tar_url, "-o", "package.tar.gz")
	err := katago_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	katago_zip_url := "https://github.com/lightvector/KataGo/archive/refs/tags/v1.15.3.zip"
	katago_cmd_zip := exec.Command("curl", "-L", katago_zip_url, "-o", "package.zip")
	err = katago_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	katago_bin_url := "https://github.com/lightvector/KataGo/archive/refs/tags/v1.15.3.bin"
	katago_cmd_bin := exec.Command("curl", "-L", katago_bin_url, "-o", "binary.bin")
	err = katago_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	katago_src_url := "https://github.com/lightvector/KataGo/archive/refs/tags/v1.15.3.src.tar.gz"
	katago_cmd_src := exec.Command("curl", "-L", katago_src_url, "-o", "source.tar.gz")
	err = katago_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	katago_cmd_direct := exec.Command("./binary")
	err = katago_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: libzip")
exec.Command("latte", "install", "libzip")
	fmt.Println("Instalando dependencia: eigen")
exec.Command("latte", "install", "eigen")
}

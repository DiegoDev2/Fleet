package main

// Opencc - Simplified-traditional Chinese conversion tool
// Homepage: https://github.com/BYVoid/OpenCC

import (
	"fmt"
	
	"os/exec"
)

func installOpencc() {
	// Método 1: Descargar y extraer .tar.gz
	opencc_tar_url := "https://github.com/BYVoid/OpenCC/archive/refs/tags/ver.1.1.9.tar.gz"
	opencc_cmd_tar := exec.Command("curl", "-L", opencc_tar_url, "-o", "package.tar.gz")
	err := opencc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	opencc_zip_url := "https://github.com/BYVoid/OpenCC/archive/refs/tags/ver.1.1.9.zip"
	opencc_cmd_zip := exec.Command("curl", "-L", opencc_zip_url, "-o", "package.zip")
	err = opencc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	opencc_bin_url := "https://github.com/BYVoid/OpenCC/archive/refs/tags/ver.1.1.9.bin"
	opencc_cmd_bin := exec.Command("curl", "-L", opencc_bin_url, "-o", "binary.bin")
	err = opencc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	opencc_src_url := "https://github.com/BYVoid/OpenCC/archive/refs/tags/ver.1.1.9.src.tar.gz"
	opencc_cmd_src := exec.Command("curl", "-L", opencc_src_url, "-o", "source.tar.gz")
	err = opencc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	opencc_cmd_direct := exec.Command("./binary")
	err = opencc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}

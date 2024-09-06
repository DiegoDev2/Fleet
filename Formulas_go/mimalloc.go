package main

// Mimalloc - Compact general purpose allocator
// Homepage: https://github.com/microsoft/mimalloc

import (
	"fmt"
	
	"os/exec"
)

func installMimalloc() {
	// Método 1: Descargar y extraer .tar.gz
	mimalloc_tar_url := "https://github.com/microsoft/mimalloc/archive/refs/tags/v2.1.7.tar.gz"
	mimalloc_cmd_tar := exec.Command("curl", "-L", mimalloc_tar_url, "-o", "package.tar.gz")
	err := mimalloc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mimalloc_zip_url := "https://github.com/microsoft/mimalloc/archive/refs/tags/v2.1.7.zip"
	mimalloc_cmd_zip := exec.Command("curl", "-L", mimalloc_zip_url, "-o", "package.zip")
	err = mimalloc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mimalloc_bin_url := "https://github.com/microsoft/mimalloc/archive/refs/tags/v2.1.7.bin"
	mimalloc_cmd_bin := exec.Command("curl", "-L", mimalloc_bin_url, "-o", "binary.bin")
	err = mimalloc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mimalloc_src_url := "https://github.com/microsoft/mimalloc/archive/refs/tags/v2.1.7.src.tar.gz"
	mimalloc_cmd_src := exec.Command("curl", "-L", mimalloc_src_url, "-o", "source.tar.gz")
	err = mimalloc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mimalloc_cmd_direct := exec.Command("./binary")
	err = mimalloc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}

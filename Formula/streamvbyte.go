package main

// Streamvbyte - Fast integer compression in C
// Homepage: https://github.com/lemire/streamvbyte

import (
	"fmt"
	
	"os/exec"
)

func installStreamvbyte() {
	// Método 1: Descargar y extraer .tar.gz
	streamvbyte_tar_url := "https://github.com/lemire/streamvbyte/archive/refs/tags/v1.0.0.tar.gz"
	streamvbyte_cmd_tar := exec.Command("curl", "-L", streamvbyte_tar_url, "-o", "package.tar.gz")
	err := streamvbyte_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	streamvbyte_zip_url := "https://github.com/lemire/streamvbyte/archive/refs/tags/v1.0.0.zip"
	streamvbyte_cmd_zip := exec.Command("curl", "-L", streamvbyte_zip_url, "-o", "package.zip")
	err = streamvbyte_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	streamvbyte_bin_url := "https://github.com/lemire/streamvbyte/archive/refs/tags/v1.0.0.bin"
	streamvbyte_cmd_bin := exec.Command("curl", "-L", streamvbyte_bin_url, "-o", "binary.bin")
	err = streamvbyte_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	streamvbyte_src_url := "https://github.com/lemire/streamvbyte/archive/refs/tags/v1.0.0.src.tar.gz"
	streamvbyte_cmd_src := exec.Command("curl", "-L", streamvbyte_src_url, "-o", "source.tar.gz")
	err = streamvbyte_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	streamvbyte_cmd_direct := exec.Command("./binary")
	err = streamvbyte_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}

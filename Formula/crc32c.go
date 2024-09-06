package main

// Crc32c - Implementation of CRC32C with CPU-specific acceleration
// Homepage: https://github.com/google/crc32c

import (
	"fmt"
	
	"os/exec"
)

func installCrc32c() {
	// Método 1: Descargar y extraer .tar.gz
	crc32c_tar_url := "https://github.com/google/crc32c/archive/refs/tags/1.1.2.tar.gz"
	crc32c_cmd_tar := exec.Command("curl", "-L", crc32c_tar_url, "-o", "package.tar.gz")
	err := crc32c_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	crc32c_zip_url := "https://github.com/google/crc32c/archive/refs/tags/1.1.2.zip"
	crc32c_cmd_zip := exec.Command("curl", "-L", crc32c_zip_url, "-o", "package.zip")
	err = crc32c_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	crc32c_bin_url := "https://github.com/google/crc32c/archive/refs/tags/1.1.2.bin"
	crc32c_cmd_bin := exec.Command("curl", "-L", crc32c_bin_url, "-o", "binary.bin")
	err = crc32c_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	crc32c_src_url := "https://github.com/google/crc32c/archive/refs/tags/1.1.2.src.tar.gz"
	crc32c_cmd_src := exec.Command("curl", "-L", crc32c_src_url, "-o", "source.tar.gz")
	err = crc32c_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	crc32c_cmd_direct := exec.Command("./binary")
	err = crc32c_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}

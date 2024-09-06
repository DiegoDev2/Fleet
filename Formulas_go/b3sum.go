package main

// B3sum - Command-line implementation of the BLAKE3 cryptographic hash function
// Homepage: https://github.com/BLAKE3-team/BLAKE3

import (
	"fmt"
	
	"os/exec"
)

func installB3sum() {
	// Método 1: Descargar y extraer .tar.gz
	b3sum_tar_url := "https://github.com/BLAKE3-team/BLAKE3/archive/refs/tags/1.5.4.tar.gz"
	b3sum_cmd_tar := exec.Command("curl", "-L", b3sum_tar_url, "-o", "package.tar.gz")
	err := b3sum_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	b3sum_zip_url := "https://github.com/BLAKE3-team/BLAKE3/archive/refs/tags/1.5.4.zip"
	b3sum_cmd_zip := exec.Command("curl", "-L", b3sum_zip_url, "-o", "package.zip")
	err = b3sum_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	b3sum_bin_url := "https://github.com/BLAKE3-team/BLAKE3/archive/refs/tags/1.5.4.bin"
	b3sum_cmd_bin := exec.Command("curl", "-L", b3sum_bin_url, "-o", "binary.bin")
	err = b3sum_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	b3sum_src_url := "https://github.com/BLAKE3-team/BLAKE3/archive/refs/tags/1.5.4.src.tar.gz"
	b3sum_cmd_src := exec.Command("curl", "-L", b3sum_src_url, "-o", "source.tar.gz")
	err = b3sum_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	b3sum_cmd_direct := exec.Command("./binary")
	err = b3sum_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}

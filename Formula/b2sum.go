package main

// B2sum - BLAKE2 b2sum reference binary
// Homepage: https://github.com/BLAKE2/BLAKE2

import (
	"fmt"
	
	"os/exec"
)

func installB2sum() {
	// Método 1: Descargar y extraer .tar.gz
	b2sum_tar_url := "https://github.com/BLAKE2/BLAKE2/archive/refs/tags/20190724.tar.gz"
	b2sum_cmd_tar := exec.Command("curl", "-L", b2sum_tar_url, "-o", "package.tar.gz")
	err := b2sum_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	b2sum_zip_url := "https://github.com/BLAKE2/BLAKE2/archive/refs/tags/20190724.zip"
	b2sum_cmd_zip := exec.Command("curl", "-L", b2sum_zip_url, "-o", "package.zip")
	err = b2sum_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	b2sum_bin_url := "https://github.com/BLAKE2/BLAKE2/archive/refs/tags/20190724.bin"
	b2sum_cmd_bin := exec.Command("curl", "-L", b2sum_bin_url, "-o", "binary.bin")
	err = b2sum_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	b2sum_src_url := "https://github.com/BLAKE2/BLAKE2/archive/refs/tags/20190724.src.tar.gz"
	b2sum_cmd_src := exec.Command("curl", "-L", b2sum_src_url, "-o", "source.tar.gz")
	err = b2sum_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	b2sum_cmd_direct := exec.Command("./binary")
	err = b2sum_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

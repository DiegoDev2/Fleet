package main

// Nq - Unix command-line queue utility
// Homepage: https://github.com/leahneukirchen/nq

import (
	"fmt"
	
	"os/exec"
)

func installNq() {
	// Método 1: Descargar y extraer .tar.gz
	nq_tar_url := "https://github.com/leahneukirchen/nq/archive/refs/tags/v1.0.tar.gz"
	nq_cmd_tar := exec.Command("curl", "-L", nq_tar_url, "-o", "package.tar.gz")
	err := nq_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nq_zip_url := "https://github.com/leahneukirchen/nq/archive/refs/tags/v1.0.zip"
	nq_cmd_zip := exec.Command("curl", "-L", nq_zip_url, "-o", "package.zip")
	err = nq_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nq_bin_url := "https://github.com/leahneukirchen/nq/archive/refs/tags/v1.0.bin"
	nq_cmd_bin := exec.Command("curl", "-L", nq_bin_url, "-o", "binary.bin")
	err = nq_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nq_src_url := "https://github.com/leahneukirchen/nq/archive/refs/tags/v1.0.src.tar.gz"
	nq_cmd_src := exec.Command("curl", "-L", nq_src_url, "-o", "source.tar.gz")
	err = nq_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nq_cmd_direct := exec.Command("./binary")
	err = nq_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

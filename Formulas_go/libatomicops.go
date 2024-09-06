package main

// LibatomicOps - Implementations for atomic memory update operations
// Homepage: https://github.com/ivmai/libatomic_ops/

import (
	"fmt"
	
	"os/exec"
)

func installLibatomicOps() {
	// Método 1: Descargar y extraer .tar.gz
	libatomicops_tar_url := "https://github.com/ivmai/libatomic_ops/releases/download/v7.8.2/libatomic_ops-7.8.2.tar.gz"
	libatomicops_cmd_tar := exec.Command("curl", "-L", libatomicops_tar_url, "-o", "package.tar.gz")
	err := libatomicops_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libatomicops_zip_url := "https://github.com/ivmai/libatomic_ops/releases/download/v7.8.2/libatomic_ops-7.8.2.zip"
	libatomicops_cmd_zip := exec.Command("curl", "-L", libatomicops_zip_url, "-o", "package.zip")
	err = libatomicops_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libatomicops_bin_url := "https://github.com/ivmai/libatomic_ops/releases/download/v7.8.2/libatomic_ops-7.8.2.bin"
	libatomicops_cmd_bin := exec.Command("curl", "-L", libatomicops_bin_url, "-o", "binary.bin")
	err = libatomicops_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libatomicops_src_url := "https://github.com/ivmai/libatomic_ops/releases/download/v7.8.2/libatomic_ops-7.8.2.src.tar.gz"
	libatomicops_cmd_src := exec.Command("curl", "-L", libatomicops_src_url, "-o", "source.tar.gz")
	err = libatomicops_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libatomicops_cmd_direct := exec.Command("./binary")
	err = libatomicops_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

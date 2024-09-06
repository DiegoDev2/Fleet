package main

// Tetra - Tetragon CLI to observe, manage and troubleshoot Tetragon instances
// Homepage: https://github.com/cilium/tetragon

import (
	"fmt"
	
	"os/exec"
)

func installTetra() {
	// Método 1: Descargar y extraer .tar.gz
	tetra_tar_url := "https://github.com/cilium/tetragon/archive/refs/tags/v1.2.0.tar.gz"
	tetra_cmd_tar := exec.Command("curl", "-L", tetra_tar_url, "-o", "package.tar.gz")
	err := tetra_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tetra_zip_url := "https://github.com/cilium/tetragon/archive/refs/tags/v1.2.0.zip"
	tetra_cmd_zip := exec.Command("curl", "-L", tetra_zip_url, "-o", "package.zip")
	err = tetra_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tetra_bin_url := "https://github.com/cilium/tetragon/archive/refs/tags/v1.2.0.bin"
	tetra_cmd_bin := exec.Command("curl", "-L", tetra_bin_url, "-o", "binary.bin")
	err = tetra_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tetra_src_url := "https://github.com/cilium/tetragon/archive/refs/tags/v1.2.0.src.tar.gz"
	tetra_cmd_src := exec.Command("curl", "-L", tetra_src_url, "-o", "source.tar.gz")
	err = tetra_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tetra_cmd_direct := exec.Command("./binary")
	err = tetra_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}

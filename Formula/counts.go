package main

// Counts - Tool for ad hoc profiling
// Homepage: https://github.com/nnethercote/counts

import (
	"fmt"
	
	"os/exec"
)

func installCounts() {
	// Método 1: Descargar y extraer .tar.gz
	counts_tar_url := "https://github.com/nnethercote/counts/archive/refs/tags/1.0.4.tar.gz"
	counts_cmd_tar := exec.Command("curl", "-L", counts_tar_url, "-o", "package.tar.gz")
	err := counts_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	counts_zip_url := "https://github.com/nnethercote/counts/archive/refs/tags/1.0.4.zip"
	counts_cmd_zip := exec.Command("curl", "-L", counts_zip_url, "-o", "package.zip")
	err = counts_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	counts_bin_url := "https://github.com/nnethercote/counts/archive/refs/tags/1.0.4.bin"
	counts_cmd_bin := exec.Command("curl", "-L", counts_bin_url, "-o", "binary.bin")
	err = counts_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	counts_src_url := "https://github.com/nnethercote/counts/archive/refs/tags/1.0.4.src.tar.gz"
	counts_cmd_src := exec.Command("curl", "-L", counts_src_url, "-o", "source.tar.gz")
	err = counts_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	counts_cmd_direct := exec.Command("./binary")
	err = counts_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}

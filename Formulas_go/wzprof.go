package main

// Wzprof - Profiling for Wazero
// Homepage: https://github.com/dispatchrun/wzprof

import (
	"fmt"
	
	"os/exec"
)

func installWzprof() {
	// Método 1: Descargar y extraer .tar.gz
	wzprof_tar_url := "https://github.com/dispatchrun/wzprof/archive/refs/tags/v0.2.0.tar.gz"
	wzprof_cmd_tar := exec.Command("curl", "-L", wzprof_tar_url, "-o", "package.tar.gz")
	err := wzprof_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wzprof_zip_url := "https://github.com/dispatchrun/wzprof/archive/refs/tags/v0.2.0.zip"
	wzprof_cmd_zip := exec.Command("curl", "-L", wzprof_zip_url, "-o", "package.zip")
	err = wzprof_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wzprof_bin_url := "https://github.com/dispatchrun/wzprof/archive/refs/tags/v0.2.0.bin"
	wzprof_cmd_bin := exec.Command("curl", "-L", wzprof_bin_url, "-o", "binary.bin")
	err = wzprof_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wzprof_src_url := "https://github.com/dispatchrun/wzprof/archive/refs/tags/v0.2.0.src.tar.gz"
	wzprof_cmd_src := exec.Command("curl", "-L", wzprof_src_url, "-o", "source.tar.gz")
	err = wzprof_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wzprof_cmd_direct := exec.Command("./binary")
	err = wzprof_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}

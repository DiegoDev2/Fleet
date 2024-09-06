package main

// Procs - Modern replacement for ps written by Rust
// Homepage: https://github.com/dalance/procs

import (
	"fmt"
	
	"os/exec"
)

func installProcs() {
	// Método 1: Descargar y extraer .tar.gz
	procs_tar_url := "https://github.com/dalance/procs/archive/refs/tags/v0.14.6.tar.gz"
	procs_cmd_tar := exec.Command("curl", "-L", procs_tar_url, "-o", "package.tar.gz")
	err := procs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	procs_zip_url := "https://github.com/dalance/procs/archive/refs/tags/v0.14.6.zip"
	procs_cmd_zip := exec.Command("curl", "-L", procs_zip_url, "-o", "package.zip")
	err = procs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	procs_bin_url := "https://github.com/dalance/procs/archive/refs/tags/v0.14.6.bin"
	procs_cmd_bin := exec.Command("curl", "-L", procs_bin_url, "-o", "binary.bin")
	err = procs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	procs_src_url := "https://github.com/dalance/procs/archive/refs/tags/v0.14.6.src.tar.gz"
	procs_cmd_src := exec.Command("curl", "-L", procs_src_url, "-o", "source.tar.gz")
	err = procs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	procs_cmd_direct := exec.Command("./binary")
	err = procs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}

package main

// Progress - Coreutils progress viewer
// Homepage: https://github.com/Xfennec/progress

import (
	"fmt"
	
	"os/exec"
)

func installProgress() {
	// Método 1: Descargar y extraer .tar.gz
	progress_tar_url := "https://github.com/Xfennec/progress/archive/refs/tags/v0.17.tar.gz"
	progress_cmd_tar := exec.Command("curl", "-L", progress_tar_url, "-o", "package.tar.gz")
	err := progress_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	progress_zip_url := "https://github.com/Xfennec/progress/archive/refs/tags/v0.17.zip"
	progress_cmd_zip := exec.Command("curl", "-L", progress_zip_url, "-o", "package.zip")
	err = progress_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	progress_bin_url := "https://github.com/Xfennec/progress/archive/refs/tags/v0.17.bin"
	progress_cmd_bin := exec.Command("curl", "-L", progress_bin_url, "-o", "binary.bin")
	err = progress_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	progress_src_url := "https://github.com/Xfennec/progress/archive/refs/tags/v0.17.src.tar.gz"
	progress_cmd_src := exec.Command("curl", "-L", progress_src_url, "-o", "source.tar.gz")
	err = progress_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	progress_cmd_direct := exec.Command("./binary")
	err = progress_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}

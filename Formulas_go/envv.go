package main

// Envv - Shell-independent handling of environment variables
// Homepage: https://github.com/jakewendt/envv

import (
	"fmt"
	
	"os/exec"
)

func installEnvv() {
	// Método 1: Descargar y extraer .tar.gz
	envv_tar_url := "https://github.com/jakewendt/envv/archive/refs/tags/v1.7.tar.gz"
	envv_cmd_tar := exec.Command("curl", "-L", envv_tar_url, "-o", "package.tar.gz")
	err := envv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	envv_zip_url := "https://github.com/jakewendt/envv/archive/refs/tags/v1.7.zip"
	envv_cmd_zip := exec.Command("curl", "-L", envv_zip_url, "-o", "package.zip")
	err = envv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	envv_bin_url := "https://github.com/jakewendt/envv/archive/refs/tags/v1.7.bin"
	envv_cmd_bin := exec.Command("curl", "-L", envv_bin_url, "-o", "binary.bin")
	err = envv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	envv_src_url := "https://github.com/jakewendt/envv/archive/refs/tags/v1.7.src.tar.gz"
	envv_cmd_src := exec.Command("curl", "-L", envv_src_url, "-o", "source.tar.gz")
	err = envv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	envv_cmd_direct := exec.Command("./binary")
	err = envv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

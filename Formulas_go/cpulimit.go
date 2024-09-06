package main

// Cpulimit - CPU usage limiter
// Homepage: https://github.com/opsengine/cpulimit

import (
	"fmt"
	
	"os/exec"
)

func installCpulimit() {
	// Método 1: Descargar y extraer .tar.gz
	cpulimit_tar_url := "https://github.com/opsengine/cpulimit/archive/refs/tags/v0.2.tar.gz"
	cpulimit_cmd_tar := exec.Command("curl", "-L", cpulimit_tar_url, "-o", "package.tar.gz")
	err := cpulimit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cpulimit_zip_url := "https://github.com/opsengine/cpulimit/archive/refs/tags/v0.2.zip"
	cpulimit_cmd_zip := exec.Command("curl", "-L", cpulimit_zip_url, "-o", "package.zip")
	err = cpulimit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cpulimit_bin_url := "https://github.com/opsengine/cpulimit/archive/refs/tags/v0.2.bin"
	cpulimit_cmd_bin := exec.Command("curl", "-L", cpulimit_bin_url, "-o", "binary.bin")
	err = cpulimit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cpulimit_src_url := "https://github.com/opsengine/cpulimit/archive/refs/tags/v0.2.src.tar.gz"
	cpulimit_cmd_src := exec.Command("curl", "-L", cpulimit_src_url, "-o", "source.tar.gz")
	err = cpulimit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cpulimit_cmd_direct := exec.Command("./binary")
	err = cpulimit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

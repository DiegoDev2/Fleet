package main

// Retry - Repeat a command until the command succeeds
// Homepage: https://github.com/minfrin/retry

import (
	"fmt"
	
	"os/exec"
)

func installRetry() {
	// Método 1: Descargar y extraer .tar.gz
	retry_tar_url := "https://github.com/minfrin/retry/releases/download/retry-1.0.5/retry-1.0.5.tar.bz2"
	retry_cmd_tar := exec.Command("curl", "-L", retry_tar_url, "-o", "package.tar.gz")
	err := retry_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	retry_zip_url := "https://github.com/minfrin/retry/releases/download/retry-1.0.5/retry-1.0.5.tar.bz2"
	retry_cmd_zip := exec.Command("curl", "-L", retry_zip_url, "-o", "package.zip")
	err = retry_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	retry_bin_url := "https://github.com/minfrin/retry/releases/download/retry-1.0.5/retry-1.0.5.tar.bz2"
	retry_cmd_bin := exec.Command("curl", "-L", retry_bin_url, "-o", "binary.bin")
	err = retry_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	retry_src_url := "https://github.com/minfrin/retry/releases/download/retry-1.0.5/retry-1.0.5.tar.bz2"
	retry_cmd_src := exec.Command("curl", "-L", retry_src_url, "-o", "source.tar.gz")
	err = retry_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	retry_cmd_direct := exec.Command("./binary")
	err = retry_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

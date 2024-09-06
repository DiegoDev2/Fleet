package main

// Shmux - Execute the same command on many hosts in parallel
// Homepage: https://github.com/shmux/shmux

import (
	"fmt"
	
	"os/exec"
)

func installShmux() {
	// Método 1: Descargar y extraer .tar.gz
	shmux_tar_url := "https://github.com/shmux/shmux/archive/refs/tags/v1.0.3.tar.gz"
	shmux_cmd_tar := exec.Command("curl", "-L", shmux_tar_url, "-o", "package.tar.gz")
	err := shmux_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	shmux_zip_url := "https://github.com/shmux/shmux/archive/refs/tags/v1.0.3.zip"
	shmux_cmd_zip := exec.Command("curl", "-L", shmux_zip_url, "-o", "package.zip")
	err = shmux_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	shmux_bin_url := "https://github.com/shmux/shmux/archive/refs/tags/v1.0.3.bin"
	shmux_cmd_bin := exec.Command("curl", "-L", shmux_bin_url, "-o", "binary.bin")
	err = shmux_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	shmux_src_url := "https://github.com/shmux/shmux/archive/refs/tags/v1.0.3.src.tar.gz"
	shmux_cmd_src := exec.Command("curl", "-L", shmux_src_url, "-o", "source.tar.gz")
	err = shmux_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	shmux_cmd_direct := exec.Command("./binary")
	err = shmux_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

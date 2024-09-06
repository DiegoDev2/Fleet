package main

// Timelimit - Limit a process's absolute execution time
// Homepage: https://devel.ringlet.net/sysutils/timelimit/

import (
	"fmt"
	
	"os/exec"
)

func installTimelimit() {
	// Método 1: Descargar y extraer .tar.gz
	timelimit_tar_url := "https://devel.ringlet.net/files/sys/timelimit/timelimit-1.9.2.tar.gz"
	timelimit_cmd_tar := exec.Command("curl", "-L", timelimit_tar_url, "-o", "package.tar.gz")
	err := timelimit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	timelimit_zip_url := "https://devel.ringlet.net/files/sys/timelimit/timelimit-1.9.2.zip"
	timelimit_cmd_zip := exec.Command("curl", "-L", timelimit_zip_url, "-o", "package.zip")
	err = timelimit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	timelimit_bin_url := "https://devel.ringlet.net/files/sys/timelimit/timelimit-1.9.2.bin"
	timelimit_cmd_bin := exec.Command("curl", "-L", timelimit_bin_url, "-o", "binary.bin")
	err = timelimit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	timelimit_src_url := "https://devel.ringlet.net/files/sys/timelimit/timelimit-1.9.2.src.tar.gz"
	timelimit_cmd_src := exec.Command("curl", "-L", timelimit_src_url, "-o", "source.tar.gz")
	err = timelimit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	timelimit_cmd_direct := exec.Command("./binary")
	err = timelimit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

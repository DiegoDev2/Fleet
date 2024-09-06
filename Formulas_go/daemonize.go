package main

// Daemonize - Run a command as a UNIX daemon
// Homepage: https://software.clapper.org/daemonize/

import (
	"fmt"
	
	"os/exec"
)

func installDaemonize() {
	// Método 1: Descargar y extraer .tar.gz
	daemonize_tar_url := "https://github.com/bmc/daemonize/archive/refs/tags/release-1.7.8.tar.gz"
	daemonize_cmd_tar := exec.Command("curl", "-L", daemonize_tar_url, "-o", "package.tar.gz")
	err := daemonize_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	daemonize_zip_url := "https://github.com/bmc/daemonize/archive/refs/tags/release-1.7.8.zip"
	daemonize_cmd_zip := exec.Command("curl", "-L", daemonize_zip_url, "-o", "package.zip")
	err = daemonize_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	daemonize_bin_url := "https://github.com/bmc/daemonize/archive/refs/tags/release-1.7.8.bin"
	daemonize_cmd_bin := exec.Command("curl", "-L", daemonize_bin_url, "-o", "binary.bin")
	err = daemonize_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	daemonize_src_url := "https://github.com/bmc/daemonize/archive/refs/tags/release-1.7.8.src.tar.gz"
	daemonize_cmd_src := exec.Command("curl", "-L", daemonize_src_url, "-o", "source.tar.gz")
	err = daemonize_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	daemonize_cmd_direct := exec.Command("./binary")
	err = daemonize_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

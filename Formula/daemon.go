package main

// Daemon - Turn other processes into daemons
// Homepage: https://libslack.org/daemon/

import (
	"fmt"
	
	"os/exec"
)

func installDaemon() {
	// Método 1: Descargar y extraer .tar.gz
	daemon_tar_url := "https://libslack.org/daemon/download/daemon-0.8.4.tar.gz"
	daemon_cmd_tar := exec.Command("curl", "-L", daemon_tar_url, "-o", "package.tar.gz")
	err := daemon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	daemon_zip_url := "https://libslack.org/daemon/download/daemon-0.8.4.zip"
	daemon_cmd_zip := exec.Command("curl", "-L", daemon_zip_url, "-o", "package.zip")
	err = daemon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	daemon_bin_url := "https://libslack.org/daemon/download/daemon-0.8.4.bin"
	daemon_cmd_bin := exec.Command("curl", "-L", daemon_bin_url, "-o", "binary.bin")
	err = daemon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	daemon_src_url := "https://libslack.org/daemon/download/daemon-0.8.4.src.tar.gz"
	daemon_cmd_src := exec.Command("curl", "-L", daemon_src_url, "-o", "source.tar.gz")
	err = daemon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	daemon_cmd_direct := exec.Command("./binary")
	err = daemon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

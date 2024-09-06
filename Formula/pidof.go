package main

// Pidof - Display the PID number for a given process name
// Homepage: http://www.nightproductions.net/cli.htm

import (
	"fmt"
	
	"os/exec"
)

func installPidof() {
	// Método 1: Descargar y extraer .tar.gz
	pidof_tar_url := "http://www.nightproductions.net/downloads/pidof_source.tar.gz"
	pidof_cmd_tar := exec.Command("curl", "-L", pidof_tar_url, "-o", "package.tar.gz")
	err := pidof_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pidof_zip_url := "http://www.nightproductions.net/downloads/pidof_source.zip"
	pidof_cmd_zip := exec.Command("curl", "-L", pidof_zip_url, "-o", "package.zip")
	err = pidof_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pidof_bin_url := "http://www.nightproductions.net/downloads/pidof_source.bin"
	pidof_cmd_bin := exec.Command("curl", "-L", pidof_bin_url, "-o", "binary.bin")
	err = pidof_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pidof_src_url := "http://www.nightproductions.net/downloads/pidof_source.src.tar.gz"
	pidof_cmd_src := exec.Command("curl", "-L", pidof_src_url, "-o", "source.tar.gz")
	err = pidof_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pidof_cmd_direct := exec.Command("./binary")
	err = pidof_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

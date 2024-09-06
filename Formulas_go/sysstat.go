package main

// Sysstat - Performance monitoring tools for Linux
// Homepage: https://github.com/sysstat/sysstat

import (
	"fmt"
	
	"os/exec"
)

func installSysstat() {
	// Método 1: Descargar y extraer .tar.gz
	sysstat_tar_url := "https://github.com/sysstat/sysstat/archive/refs/tags/v12.7.6.tar.gz"
	sysstat_cmd_tar := exec.Command("curl", "-L", sysstat_tar_url, "-o", "package.tar.gz")
	err := sysstat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sysstat_zip_url := "https://github.com/sysstat/sysstat/archive/refs/tags/v12.7.6.zip"
	sysstat_cmd_zip := exec.Command("curl", "-L", sysstat_zip_url, "-o", "package.zip")
	err = sysstat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sysstat_bin_url := "https://github.com/sysstat/sysstat/archive/refs/tags/v12.7.6.bin"
	sysstat_cmd_bin := exec.Command("curl", "-L", sysstat_bin_url, "-o", "binary.bin")
	err = sysstat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sysstat_src_url := "https://github.com/sysstat/sysstat/archive/refs/tags/v12.7.6.src.tar.gz"
	sysstat_cmd_src := exec.Command("curl", "-L", sysstat_src_url, "-o", "source.tar.gz")
	err = sysstat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sysstat_cmd_direct := exec.Command("./binary")
	err = sysstat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

package main

// Ifstat - Tool to report network interface bandwidth
// Homepage: http://gael.roualland.free.fr/ifstat/

import (
	"fmt"
	
	"os/exec"
)

func installIfstat() {
	// Método 1: Descargar y extraer .tar.gz
	ifstat_tar_url := "http://gael.roualland.free.fr/ifstat/ifstat-1.1.tar.gz"
	ifstat_cmd_tar := exec.Command("curl", "-L", ifstat_tar_url, "-o", "package.tar.gz")
	err := ifstat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ifstat_zip_url := "http://gael.roualland.free.fr/ifstat/ifstat-1.1.zip"
	ifstat_cmd_zip := exec.Command("curl", "-L", ifstat_zip_url, "-o", "package.zip")
	err = ifstat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ifstat_bin_url := "http://gael.roualland.free.fr/ifstat/ifstat-1.1.bin"
	ifstat_cmd_bin := exec.Command("curl", "-L", ifstat_bin_url, "-o", "binary.bin")
	err = ifstat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ifstat_src_url := "http://gael.roualland.free.fr/ifstat/ifstat-1.1.src.tar.gz"
	ifstat_cmd_src := exec.Command("curl", "-L", ifstat_src_url, "-o", "source.tar.gz")
	err = ifstat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ifstat_cmd_direct := exec.Command("./binary")
	err = ifstat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

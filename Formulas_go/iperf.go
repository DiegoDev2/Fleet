package main

// Iperf - Tool to measure maximum TCP and UDP bandwidth
// Homepage: https://sourceforge.net/projects/iperf2/

import (
	"fmt"
	
	"os/exec"
)

func installIperf() {
	// Método 1: Descargar y extraer .tar.gz
	iperf_tar_url := "https://downloads.sourceforge.net/project/iperf2/iperf-2.2.0.tar.gz"
	iperf_cmd_tar := exec.Command("curl", "-L", iperf_tar_url, "-o", "package.tar.gz")
	err := iperf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	iperf_zip_url := "https://downloads.sourceforge.net/project/iperf2/iperf-2.2.0.zip"
	iperf_cmd_zip := exec.Command("curl", "-L", iperf_zip_url, "-o", "package.zip")
	err = iperf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	iperf_bin_url := "https://downloads.sourceforge.net/project/iperf2/iperf-2.2.0.bin"
	iperf_cmd_bin := exec.Command("curl", "-L", iperf_bin_url, "-o", "binary.bin")
	err = iperf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	iperf_src_url := "https://downloads.sourceforge.net/project/iperf2/iperf-2.2.0.src.tar.gz"
	iperf_cmd_src := exec.Command("curl", "-L", iperf_src_url, "-o", "source.tar.gz")
	err = iperf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	iperf_cmd_direct := exec.Command("./binary")
	err = iperf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

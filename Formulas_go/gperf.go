package main

// Gperf - Perfect hash function generator
// Homepage: https://www.gnu.org/software/gperf/

import (
	"fmt"
	
	"os/exec"
)

func installGperf() {
	// Método 1: Descargar y extraer .tar.gz
	gperf_tar_url := "https://ftp.gnu.org/gnu/gperf/gperf-3.1.tar.gz"
	gperf_cmd_tar := exec.Command("curl", "-L", gperf_tar_url, "-o", "package.tar.gz")
	err := gperf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gperf_zip_url := "https://ftp.gnu.org/gnu/gperf/gperf-3.1.zip"
	gperf_cmd_zip := exec.Command("curl", "-L", gperf_zip_url, "-o", "package.zip")
	err = gperf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gperf_bin_url := "https://ftp.gnu.org/gnu/gperf/gperf-3.1.bin"
	gperf_cmd_bin := exec.Command("curl", "-L", gperf_bin_url, "-o", "binary.bin")
	err = gperf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gperf_src_url := "https://ftp.gnu.org/gnu/gperf/gperf-3.1.src.tar.gz"
	gperf_cmd_src := exec.Command("curl", "-L", gperf_src_url, "-o", "source.tar.gz")
	err = gperf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gperf_cmd_direct := exec.Command("./binary")
	err = gperf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

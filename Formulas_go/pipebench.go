package main

// Pipebench - Measure the speed of STDIN/STDOUT communication
// Homepage: https://www.habets.pp.se/synscan/programs_pipebench.html

import (
	"fmt"
	
	"os/exec"
)

func installPipebench() {
	// Método 1: Descargar y extraer .tar.gz
	pipebench_tar_url := "http://www.habets.pp.se/synscan/files/pipebench-0.40.tar.gz"
	pipebench_cmd_tar := exec.Command("curl", "-L", pipebench_tar_url, "-o", "package.tar.gz")
	err := pipebench_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pipebench_zip_url := "http://www.habets.pp.se/synscan/files/pipebench-0.40.zip"
	pipebench_cmd_zip := exec.Command("curl", "-L", pipebench_zip_url, "-o", "package.zip")
	err = pipebench_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pipebench_bin_url := "http://www.habets.pp.se/synscan/files/pipebench-0.40.bin"
	pipebench_cmd_bin := exec.Command("curl", "-L", pipebench_bin_url, "-o", "binary.bin")
	err = pipebench_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pipebench_src_url := "http://www.habets.pp.se/synscan/files/pipebench-0.40.src.tar.gz"
	pipebench_cmd_src := exec.Command("curl", "-L", pipebench_src_url, "-o", "source.tar.gz")
	err = pipebench_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pipebench_cmd_direct := exec.Command("./binary")
	err = pipebench_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

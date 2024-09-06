package main

// Minised - Smaller, cheaper, faster SED implementation
// Homepage: https://www.exactcode.com/opensource/minised/

import (
	"fmt"
	
	"os/exec"
)

func installMinised() {
	// Método 1: Descargar y extraer .tar.gz
	minised_tar_url := "https://dl.exactcode.de/oss/minised/minised-1.16.tar.gz"
	minised_cmd_tar := exec.Command("curl", "-L", minised_tar_url, "-o", "package.tar.gz")
	err := minised_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	minised_zip_url := "https://dl.exactcode.de/oss/minised/minised-1.16.zip"
	minised_cmd_zip := exec.Command("curl", "-L", minised_zip_url, "-o", "package.zip")
	err = minised_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	minised_bin_url := "https://dl.exactcode.de/oss/minised/minised-1.16.bin"
	minised_cmd_bin := exec.Command("curl", "-L", minised_bin_url, "-o", "binary.bin")
	err = minised_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	minised_src_url := "https://dl.exactcode.de/oss/minised/minised-1.16.src.tar.gz"
	minised_cmd_src := exec.Command("curl", "-L", minised_src_url, "-o", "source.tar.gz")
	err = minised_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	minised_cmd_direct := exec.Command("./binary")
	err = minised_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

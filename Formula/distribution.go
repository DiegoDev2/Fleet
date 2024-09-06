package main

// Distribution - Create ASCII graphical histograms in the terminal
// Homepage: https://github.com/time-less-ness/distribution

import (
	"fmt"
	
	"os/exec"
)

func installDistribution() {
	// Método 1: Descargar y extraer .tar.gz
	distribution_tar_url := "https://github.com/time-less-ness/distribution/archive/refs/tags/1.3.tar.gz"
	distribution_cmd_tar := exec.Command("curl", "-L", distribution_tar_url, "-o", "package.tar.gz")
	err := distribution_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	distribution_zip_url := "https://github.com/time-less-ness/distribution/archive/refs/tags/1.3.zip"
	distribution_cmd_zip := exec.Command("curl", "-L", distribution_zip_url, "-o", "package.zip")
	err = distribution_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	distribution_bin_url := "https://github.com/time-less-ness/distribution/archive/refs/tags/1.3.bin"
	distribution_cmd_bin := exec.Command("curl", "-L", distribution_bin_url, "-o", "binary.bin")
	err = distribution_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	distribution_src_url := "https://github.com/time-less-ness/distribution/archive/refs/tags/1.3.src.tar.gz"
	distribution_cmd_src := exec.Command("curl", "-L", distribution_src_url, "-o", "source.tar.gz")
	err = distribution_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	distribution_cmd_direct := exec.Command("./binary")
	err = distribution_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

package main

// Swftools - SWF manipulation and generation tools
// Homepage: http://www.swftools.org/

import (
	"fmt"
	
	"os/exec"
)

func installSwftools() {
	// Método 1: Descargar y extraer .tar.gz
	swftools_tar_url := "http://www.swftools.org/swftools-0.9.2.tar.gz"
	swftools_cmd_tar := exec.Command("curl", "-L", swftools_tar_url, "-o", "package.tar.gz")
	err := swftools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	swftools_zip_url := "http://www.swftools.org/swftools-0.9.2.zip"
	swftools_cmd_zip := exec.Command("curl", "-L", swftools_zip_url, "-o", "package.zip")
	err = swftools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	swftools_bin_url := "http://www.swftools.org/swftools-0.9.2.bin"
	swftools_cmd_bin := exec.Command("curl", "-L", swftools_bin_url, "-o", "binary.bin")
	err = swftools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	swftools_src_url := "http://www.swftools.org/swftools-0.9.2.src.tar.gz"
	swftools_cmd_src := exec.Command("curl", "-L", swftools_src_url, "-o", "source.tar.gz")
	err = swftools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	swftools_cmd_direct := exec.Command("./binary")
	err = swftools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

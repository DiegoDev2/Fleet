package main

// Bedops - Set and statistical operations on genomic data of arbitrary scale
// Homepage: https://github.com/bedops/bedops

import (
	"fmt"
	
	"os/exec"
)

func installBedops() {
	// Método 1: Descargar y extraer .tar.gz
	bedops_tar_url := "https://github.com/bedops/bedops/archive/refs/tags/v2.4.41.tar.gz"
	bedops_cmd_tar := exec.Command("curl", "-L", bedops_tar_url, "-o", "package.tar.gz")
	err := bedops_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bedops_zip_url := "https://github.com/bedops/bedops/archive/refs/tags/v2.4.41.zip"
	bedops_cmd_zip := exec.Command("curl", "-L", bedops_zip_url, "-o", "package.zip")
	err = bedops_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bedops_bin_url := "https://github.com/bedops/bedops/archive/refs/tags/v2.4.41.bin"
	bedops_cmd_bin := exec.Command("curl", "-L", bedops_bin_url, "-o", "binary.bin")
	err = bedops_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bedops_src_url := "https://github.com/bedops/bedops/archive/refs/tags/v2.4.41.src.tar.gz"
	bedops_cmd_src := exec.Command("curl", "-L", bedops_src_url, "-o", "source.tar.gz")
	err = bedops_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bedops_cmd_direct := exec.Command("./binary")
	err = bedops_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: jansson")
	exec.Command("latte", "install", "jansson").Run()
}

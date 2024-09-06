package main

// Cheops - CHEss OPponent Simulator
// Homepage: https://logological.org/cheops

import (
	"fmt"
	
	"os/exec"
)

func installCheops() {
	// Método 1: Descargar y extraer .tar.gz
	cheops_tar_url := "https://files.nothingisreal.com/software/cheops/cheops-1.3.tar.bz2"
	cheops_cmd_tar := exec.Command("curl", "-L", cheops_tar_url, "-o", "package.tar.gz")
	err := cheops_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cheops_zip_url := "https://files.nothingisreal.com/software/cheops/cheops-1.3.tar.bz2"
	cheops_cmd_zip := exec.Command("curl", "-L", cheops_zip_url, "-o", "package.zip")
	err = cheops_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cheops_bin_url := "https://files.nothingisreal.com/software/cheops/cheops-1.3.tar.bz2"
	cheops_cmd_bin := exec.Command("curl", "-L", cheops_bin_url, "-o", "binary.bin")
	err = cheops_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cheops_src_url := "https://files.nothingisreal.com/software/cheops/cheops-1.3.tar.bz2"
	cheops_cmd_src := exec.Command("curl", "-L", cheops_src_url, "-o", "source.tar.gz")
	err = cheops_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cheops_cmd_direct := exec.Command("./binary")
	err = cheops_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

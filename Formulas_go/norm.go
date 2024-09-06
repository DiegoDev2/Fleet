package main

// Norm - NACK-Oriented Reliable Multicast
// Homepage: https://www.nrl.navy.mil/itd/ncs/products/norm

import (
	"fmt"
	
	"os/exec"
)

func installNorm() {
	// Método 1: Descargar y extraer .tar.gz
	norm_tar_url := "https://github.com/USNavalResearchLaboratory/norm/releases/download/v1.5.9/src-norm-1.5.9.tgz"
	norm_cmd_tar := exec.Command("curl", "-L", norm_tar_url, "-o", "package.tar.gz")
	err := norm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	norm_zip_url := "https://github.com/USNavalResearchLaboratory/norm/releases/download/v1.5.9/src-norm-1.5.9.tgz"
	norm_cmd_zip := exec.Command("curl", "-L", norm_zip_url, "-o", "package.zip")
	err = norm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	norm_bin_url := "https://github.com/USNavalResearchLaboratory/norm/releases/download/v1.5.9/src-norm-1.5.9.tgz"
	norm_cmd_bin := exec.Command("curl", "-L", norm_bin_url, "-o", "binary.bin")
	err = norm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	norm_src_url := "https://github.com/USNavalResearchLaboratory/norm/releases/download/v1.5.9/src-norm-1.5.9.tgz"
	norm_cmd_src := exec.Command("curl", "-L", norm_src_url, "-o", "source.tar.gz")
	err = norm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	norm_cmd_direct := exec.Command("./binary")
	err = norm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

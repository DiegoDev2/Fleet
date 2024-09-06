package main

// Ssdeep - Recursive piecewise hashing tool
// Homepage: https://ssdeep-project.github.io/ssdeep/

import (
	"fmt"
	
	"os/exec"
)

func installSsdeep() {
	// Método 1: Descargar y extraer .tar.gz
	ssdeep_tar_url := "https://github.com/ssdeep-project/ssdeep/releases/download/release-2.14.1/ssdeep-2.14.1.tar.gz"
	ssdeep_cmd_tar := exec.Command("curl", "-L", ssdeep_tar_url, "-o", "package.tar.gz")
	err := ssdeep_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ssdeep_zip_url := "https://github.com/ssdeep-project/ssdeep/releases/download/release-2.14.1/ssdeep-2.14.1.zip"
	ssdeep_cmd_zip := exec.Command("curl", "-L", ssdeep_zip_url, "-o", "package.zip")
	err = ssdeep_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ssdeep_bin_url := "https://github.com/ssdeep-project/ssdeep/releases/download/release-2.14.1/ssdeep-2.14.1.bin"
	ssdeep_cmd_bin := exec.Command("curl", "-L", ssdeep_bin_url, "-o", "binary.bin")
	err = ssdeep_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ssdeep_src_url := "https://github.com/ssdeep-project/ssdeep/releases/download/release-2.14.1/ssdeep-2.14.1.src.tar.gz"
	ssdeep_cmd_src := exec.Command("curl", "-L", ssdeep_src_url, "-o", "source.tar.gz")
	err = ssdeep_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ssdeep_cmd_direct := exec.Command("./binary")
	err = ssdeep_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

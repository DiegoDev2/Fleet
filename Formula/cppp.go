package main

// Cppp - Partial Preprocessor for C
// Homepage: https://www.muppetlabs.com/~breadbox/software/cppp.html

import (
	"fmt"
	
	"os/exec"
)

func installCppp() {
	// Método 1: Descargar y extraer .tar.gz
	cppp_tar_url := "https://www.muppetlabs.com/~breadbox/pub/software/cppp-2.9.tar.gz"
	cppp_cmd_tar := exec.Command("curl", "-L", cppp_tar_url, "-o", "package.tar.gz")
	err := cppp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cppp_zip_url := "https://www.muppetlabs.com/~breadbox/pub/software/cppp-2.9.zip"
	cppp_cmd_zip := exec.Command("curl", "-L", cppp_zip_url, "-o", "package.zip")
	err = cppp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cppp_bin_url := "https://www.muppetlabs.com/~breadbox/pub/software/cppp-2.9.bin"
	cppp_cmd_bin := exec.Command("curl", "-L", cppp_bin_url, "-o", "binary.bin")
	err = cppp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cppp_src_url := "https://www.muppetlabs.com/~breadbox/pub/software/cppp-2.9.src.tar.gz"
	cppp_cmd_src := exec.Command("curl", "-L", cppp_src_url, "-o", "source.tar.gz")
	err = cppp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cppp_cmd_direct := exec.Command("./binary")
	err = cppp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

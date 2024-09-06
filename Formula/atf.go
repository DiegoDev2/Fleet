package main

// Atf - Automated testing framework
// Homepage: https://github.com/freebsd/atf

import (
	"fmt"
	
	"os/exec"
)

func installAtf() {
	// Método 1: Descargar y extraer .tar.gz
	atf_tar_url := "https://github.com/freebsd/atf/releases/download/atf-0.21/atf-0.21.tar.gz"
	atf_cmd_tar := exec.Command("curl", "-L", atf_tar_url, "-o", "package.tar.gz")
	err := atf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	atf_zip_url := "https://github.com/freebsd/atf/releases/download/atf-0.21/atf-0.21.zip"
	atf_cmd_zip := exec.Command("curl", "-L", atf_zip_url, "-o", "package.zip")
	err = atf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	atf_bin_url := "https://github.com/freebsd/atf/releases/download/atf-0.21/atf-0.21.bin"
	atf_cmd_bin := exec.Command("curl", "-L", atf_bin_url, "-o", "binary.bin")
	err = atf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	atf_src_url := "https://github.com/freebsd/atf/releases/download/atf-0.21/atf-0.21.src.tar.gz"
	atf_cmd_src := exec.Command("curl", "-L", atf_src_url, "-o", "source.tar.gz")
	err = atf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	atf_cmd_direct := exec.Command("./binary")
	err = atf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

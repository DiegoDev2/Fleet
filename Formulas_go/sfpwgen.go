package main

// SfPwgen - Generate passwords using SecurityFoundation framework
// Homepage: https://github.com/anders/pwgen/

import (
	"fmt"
	
	"os/exec"
)

func installSfPwgen() {
	// Método 1: Descargar y extraer .tar.gz
	sfpwgen_tar_url := "https://github.com/anders/pwgen/archive/refs/tags/1.5.tar.gz"
	sfpwgen_cmd_tar := exec.Command("curl", "-L", sfpwgen_tar_url, "-o", "package.tar.gz")
	err := sfpwgen_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sfpwgen_zip_url := "https://github.com/anders/pwgen/archive/refs/tags/1.5.zip"
	sfpwgen_cmd_zip := exec.Command("curl", "-L", sfpwgen_zip_url, "-o", "package.zip")
	err = sfpwgen_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sfpwgen_bin_url := "https://github.com/anders/pwgen/archive/refs/tags/1.5.bin"
	sfpwgen_cmd_bin := exec.Command("curl", "-L", sfpwgen_bin_url, "-o", "binary.bin")
	err = sfpwgen_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sfpwgen_src_url := "https://github.com/anders/pwgen/archive/refs/tags/1.5.src.tar.gz"
	sfpwgen_cmd_src := exec.Command("curl", "-L", sfpwgen_src_url, "-o", "source.tar.gz")
	err = sfpwgen_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sfpwgen_cmd_direct := exec.Command("./binary")
	err = sfpwgen_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

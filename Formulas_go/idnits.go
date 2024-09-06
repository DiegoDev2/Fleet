package main

// Idnits - Looks for problems in internet draft formatting
// Homepage: https://github.com/ietf-tools/idnits

import (
	"fmt"
	
	"os/exec"
)

func installIdnits() {
	// Método 1: Descargar y extraer .tar.gz
	idnits_tar_url := "https://github.com/ietf-tools/idnits/archive/refs/tags/2.17.1.tar.gz"
	idnits_cmd_tar := exec.Command("curl", "-L", idnits_tar_url, "-o", "package.tar.gz")
	err := idnits_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	idnits_zip_url := "https://github.com/ietf-tools/idnits/archive/refs/tags/2.17.1.zip"
	idnits_cmd_zip := exec.Command("curl", "-L", idnits_zip_url, "-o", "package.zip")
	err = idnits_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	idnits_bin_url := "https://github.com/ietf-tools/idnits/archive/refs/tags/2.17.1.bin"
	idnits_cmd_bin := exec.Command("curl", "-L", idnits_bin_url, "-o", "binary.bin")
	err = idnits_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	idnits_src_url := "https://github.com/ietf-tools/idnits/archive/refs/tags/2.17.1.src.tar.gz"
	idnits_cmd_src := exec.Command("curl", "-L", idnits_src_url, "-o", "source.tar.gz")
	err = idnits_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	idnits_cmd_direct := exec.Command("./binary")
	err = idnits_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

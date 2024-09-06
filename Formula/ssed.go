package main

// Ssed - Super sed stream editor
// Homepage: https://sed.sourceforge.io/grabbag/ssed/

import (
	"fmt"
	
	"os/exec"
)

func installSsed() {
	// Método 1: Descargar y extraer .tar.gz
	ssed_tar_url := "https://sed.sourceforge.io/grabbag/ssed/sed-3.62.tar.gz"
	ssed_cmd_tar := exec.Command("curl", "-L", ssed_tar_url, "-o", "package.tar.gz")
	err := ssed_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ssed_zip_url := "https://sed.sourceforge.io/grabbag/ssed/sed-3.62.zip"
	ssed_cmd_zip := exec.Command("curl", "-L", ssed_zip_url, "-o", "package.zip")
	err = ssed_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ssed_bin_url := "https://sed.sourceforge.io/grabbag/ssed/sed-3.62.bin"
	ssed_cmd_bin := exec.Command("curl", "-L", ssed_bin_url, "-o", "binary.bin")
	err = ssed_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ssed_src_url := "https://sed.sourceforge.io/grabbag/ssed/sed-3.62.src.tar.gz"
	ssed_cmd_src := exec.Command("curl", "-L", ssed_src_url, "-o", "source.tar.gz")
	err = ssed_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ssed_cmd_direct := exec.Command("./binary")
	err = ssed_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

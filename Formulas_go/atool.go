package main

// Atool - Archival front-end
// Homepage: https://savannah.nongnu.org/projects/atool/

import (
	"fmt"
	
	"os/exec"
)

func installAtool() {
	// Método 1: Descargar y extraer .tar.gz
	atool_tar_url := "https://savannah.nongnu.org/download/atool/atool-0.39.0.tar.gz"
	atool_cmd_tar := exec.Command("curl", "-L", atool_tar_url, "-o", "package.tar.gz")
	err := atool_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	atool_zip_url := "https://savannah.nongnu.org/download/atool/atool-0.39.0.zip"
	atool_cmd_zip := exec.Command("curl", "-L", atool_zip_url, "-o", "package.zip")
	err = atool_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	atool_bin_url := "https://savannah.nongnu.org/download/atool/atool-0.39.0.bin"
	atool_cmd_bin := exec.Command("curl", "-L", atool_bin_url, "-o", "binary.bin")
	err = atool_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	atool_src_url := "https://savannah.nongnu.org/download/atool/atool-0.39.0.src.tar.gz"
	atool_cmd_src := exec.Command("curl", "-L", atool_src_url, "-o", "source.tar.gz")
	err = atool_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	atool_cmd_direct := exec.Command("./binary")
	err = atool_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

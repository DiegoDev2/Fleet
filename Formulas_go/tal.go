package main

// Tal - Align line endings if they match
// Homepage: https://thomasjensen.com/software/tal/

import (
	"fmt"
	
	"os/exec"
)

func installTal() {
	// Método 1: Descargar y extraer .tar.gz
	tal_tar_url := "https://thomasjensen.com/software/tal/tal-1.9.tar.gz"
	tal_cmd_tar := exec.Command("curl", "-L", tal_tar_url, "-o", "package.tar.gz")
	err := tal_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tal_zip_url := "https://thomasjensen.com/software/tal/tal-1.9.zip"
	tal_cmd_zip := exec.Command("curl", "-L", tal_zip_url, "-o", "package.zip")
	err = tal_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tal_bin_url := "https://thomasjensen.com/software/tal/tal-1.9.bin"
	tal_cmd_bin := exec.Command("curl", "-L", tal_bin_url, "-o", "binary.bin")
	err = tal_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tal_src_url := "https://thomasjensen.com/software/tal/tal-1.9.src.tar.gz"
	tal_cmd_src := exec.Command("curl", "-L", tal_src_url, "-o", "source.tar.gz")
	err = tal_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tal_cmd_direct := exec.Command("./binary")
	err = tal_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

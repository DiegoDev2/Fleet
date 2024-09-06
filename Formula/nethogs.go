package main

// Nethogs - Net top tool grouping bandwidth per process
// Homepage: https://raboof.github.io/nethogs/

import (
	"fmt"
	
	"os/exec"
)

func installNethogs() {
	// Método 1: Descargar y extraer .tar.gz
	nethogs_tar_url := "https://github.com/raboof/nethogs/archive/refs/tags/v0.8.7.tar.gz"
	nethogs_cmd_tar := exec.Command("curl", "-L", nethogs_tar_url, "-o", "package.tar.gz")
	err := nethogs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nethogs_zip_url := "https://github.com/raboof/nethogs/archive/refs/tags/v0.8.7.zip"
	nethogs_cmd_zip := exec.Command("curl", "-L", nethogs_zip_url, "-o", "package.zip")
	err = nethogs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nethogs_bin_url := "https://github.com/raboof/nethogs/archive/refs/tags/v0.8.7.bin"
	nethogs_cmd_bin := exec.Command("curl", "-L", nethogs_bin_url, "-o", "binary.bin")
	err = nethogs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nethogs_src_url := "https://github.com/raboof/nethogs/archive/refs/tags/v0.8.7.src.tar.gz"
	nethogs_cmd_src := exec.Command("curl", "-L", nethogs_src_url, "-o", "source.tar.gz")
	err = nethogs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nethogs_cmd_direct := exec.Command("./binary")
	err = nethogs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

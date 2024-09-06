package main

// Pit - Project manager from hell (integrates with Git)
// Homepage: https://github.com/michaeldv/pit

import (
	"fmt"
	
	"os/exec"
)

func installPit() {
	// Método 1: Descargar y extraer .tar.gz
	pit_tar_url := "https://github.com/michaeldv/pit/archive/refs/tags/0.1.0.tar.gz"
	pit_cmd_tar := exec.Command("curl", "-L", pit_tar_url, "-o", "package.tar.gz")
	err := pit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pit_zip_url := "https://github.com/michaeldv/pit/archive/refs/tags/0.1.0.zip"
	pit_cmd_zip := exec.Command("curl", "-L", pit_zip_url, "-o", "package.zip")
	err = pit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pit_bin_url := "https://github.com/michaeldv/pit/archive/refs/tags/0.1.0.bin"
	pit_cmd_bin := exec.Command("curl", "-L", pit_bin_url, "-o", "binary.bin")
	err = pit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pit_src_url := "https://github.com/michaeldv/pit/archive/refs/tags/0.1.0.src.tar.gz"
	pit_cmd_src := exec.Command("curl", "-L", pit_src_url, "-o", "source.tar.gz")
	err = pit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pit_cmd_direct := exec.Command("./binary")
	err = pit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

package main

// Angband - Dungeon exploration game
// Homepage: https://angband.github.io/angband/

import (
	"fmt"
	
	"os/exec"
)

func installAngband() {
	// Método 1: Descargar y extraer .tar.gz
	angband_tar_url := "https://github.com/angband/angband/releases/download/4.2.5/Angband-4.2.5.tar.gz"
	angband_cmd_tar := exec.Command("curl", "-L", angband_tar_url, "-o", "package.tar.gz")
	err := angband_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	angband_zip_url := "https://github.com/angband/angband/releases/download/4.2.5/Angband-4.2.5.zip"
	angband_cmd_zip := exec.Command("curl", "-L", angband_zip_url, "-o", "package.zip")
	err = angband_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	angband_bin_url := "https://github.com/angband/angband/releases/download/4.2.5/Angband-4.2.5.bin"
	angband_cmd_bin := exec.Command("curl", "-L", angband_bin_url, "-o", "binary.bin")
	err = angband_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	angband_src_url := "https://github.com/angband/angband/releases/download/4.2.5/Angband-4.2.5.src.tar.gz"
	angband_cmd_src := exec.Command("curl", "-L", angband_src_url, "-o", "source.tar.gz")
	err = angband_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	angband_cmd_direct := exec.Command("./binary")
	err = angband_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

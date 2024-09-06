package main

// Nbsdgames - Text-based modern games
// Homepage: https://github.com/abakh/nbsdgames

import (
	"fmt"
	
	"os/exec"
)

func installNbsdgames() {
	// Método 1: Descargar y extraer .tar.gz
	nbsdgames_tar_url := "https://github.com/abakh/nbsdgames/archive/refs/tags/v5.tar.gz"
	nbsdgames_cmd_tar := exec.Command("curl", "-L", nbsdgames_tar_url, "-o", "package.tar.gz")
	err := nbsdgames_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nbsdgames_zip_url := "https://github.com/abakh/nbsdgames/archive/refs/tags/v5.zip"
	nbsdgames_cmd_zip := exec.Command("curl", "-L", nbsdgames_zip_url, "-o", "package.zip")
	err = nbsdgames_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nbsdgames_bin_url := "https://github.com/abakh/nbsdgames/archive/refs/tags/v5.bin"
	nbsdgames_cmd_bin := exec.Command("curl", "-L", nbsdgames_bin_url, "-o", "binary.bin")
	err = nbsdgames_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nbsdgames_src_url := "https://github.com/abakh/nbsdgames/archive/refs/tags/v5.src.tar.gz"
	nbsdgames_cmd_src := exec.Command("curl", "-L", nbsdgames_src_url, "-o", "source.tar.gz")
	err = nbsdgames_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nbsdgames_cmd_direct := exec.Command("./binary")
	err = nbsdgames_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

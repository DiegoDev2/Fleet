package main

// Regldg - Regular expression grammar language dictionary generator
// Homepage: https://github.com/PatrickCronin/regldg

import (
	"fmt"
	
	"os/exec"
)

func installRegldg() {
	// Método 1: Descargar y extraer .tar.gz
	regldg_tar_url := "https://github.com/PatrickCronin/regldg/releases/download/v1.0.1/regldg-1.0.1.tar.gz"
	regldg_cmd_tar := exec.Command("curl", "-L", regldg_tar_url, "-o", "package.tar.gz")
	err := regldg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	regldg_zip_url := "https://github.com/PatrickCronin/regldg/releases/download/v1.0.1/regldg-1.0.1.zip"
	regldg_cmd_zip := exec.Command("curl", "-L", regldg_zip_url, "-o", "package.zip")
	err = regldg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	regldg_bin_url := "https://github.com/PatrickCronin/regldg/releases/download/v1.0.1/regldg-1.0.1.bin"
	regldg_cmd_bin := exec.Command("curl", "-L", regldg_bin_url, "-o", "binary.bin")
	err = regldg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	regldg_src_url := "https://github.com/PatrickCronin/regldg/releases/download/v1.0.1/regldg-1.0.1.src.tar.gz"
	regldg_cmd_src := exec.Command("curl", "-L", regldg_src_url, "-o", "source.tar.gz")
	err = regldg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	regldg_cmd_direct := exec.Command("./binary")
	err = regldg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

package main

// Phoon - Displays current or specified phase of the moon via ASCII art
// Homepage: https://www.acme.com/software/phoon/

import (
	"fmt"
	
	"os/exec"
)

func installPhoon() {
	// Método 1: Descargar y extraer .tar.gz
	phoon_tar_url := "https://www.acme.com/software/phoon/phoon_14Aug2014.tar.gz"
	phoon_cmd_tar := exec.Command("curl", "-L", phoon_tar_url, "-o", "package.tar.gz")
	err := phoon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	phoon_zip_url := "https://www.acme.com/software/phoon/phoon_14Aug2014.zip"
	phoon_cmd_zip := exec.Command("curl", "-L", phoon_zip_url, "-o", "package.zip")
	err = phoon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	phoon_bin_url := "https://www.acme.com/software/phoon/phoon_14Aug2014.bin"
	phoon_cmd_bin := exec.Command("curl", "-L", phoon_bin_url, "-o", "binary.bin")
	err = phoon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	phoon_src_url := "https://www.acme.com/software/phoon/phoon_14Aug2014.src.tar.gz"
	phoon_cmd_src := exec.Command("curl", "-L", phoon_src_url, "-o", "source.tar.gz")
	err = phoon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	phoon_cmd_direct := exec.Command("./binary")
	err = phoon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

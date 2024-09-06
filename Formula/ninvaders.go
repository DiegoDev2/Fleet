package main

// Ninvaders - Space Invaders in the terminal
// Homepage: https://ninvaders.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installNinvaders() {
	// Método 1: Descargar y extraer .tar.gz
	ninvaders_tar_url := "https://downloads.sourceforge.net/project/ninvaders/ninvaders/0.1.1/ninvaders-0.1.1.tar.gz"
	ninvaders_cmd_tar := exec.Command("curl", "-L", ninvaders_tar_url, "-o", "package.tar.gz")
	err := ninvaders_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ninvaders_zip_url := "https://downloads.sourceforge.net/project/ninvaders/ninvaders/0.1.1/ninvaders-0.1.1.zip"
	ninvaders_cmd_zip := exec.Command("curl", "-L", ninvaders_zip_url, "-o", "package.zip")
	err = ninvaders_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ninvaders_bin_url := "https://downloads.sourceforge.net/project/ninvaders/ninvaders/0.1.1/ninvaders-0.1.1.bin"
	ninvaders_cmd_bin := exec.Command("curl", "-L", ninvaders_bin_url, "-o", "binary.bin")
	err = ninvaders_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ninvaders_src_url := "https://downloads.sourceforge.net/project/ninvaders/ninvaders/0.1.1/ninvaders-0.1.1.src.tar.gz"
	ninvaders_cmd_src := exec.Command("curl", "-L", ninvaders_src_url, "-o", "source.tar.gz")
	err = ninvaders_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ninvaders_cmd_direct := exec.Command("./binary")
	err = ninvaders_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

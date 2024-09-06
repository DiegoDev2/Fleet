package main

// Fortune - Infamous electronic fortune-cookie generator
// Homepage: https://www.ibiblio.org/pub/linux/games/amusements/fortune/!INDEX.html

import (
	"fmt"
	
	"os/exec"
)

func installFortune() {
	// Método 1: Descargar y extraer .tar.gz
	fortune_tar_url := "https://www.ibiblio.org/pub/linux/games/amusements/fortune/fortune-mod-9708.tar.gz"
	fortune_cmd_tar := exec.Command("curl", "-L", fortune_tar_url, "-o", "package.tar.gz")
	err := fortune_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fortune_zip_url := "https://www.ibiblio.org/pub/linux/games/amusements/fortune/fortune-mod-9708.zip"
	fortune_cmd_zip := exec.Command("curl", "-L", fortune_zip_url, "-o", "package.zip")
	err = fortune_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fortune_bin_url := "https://www.ibiblio.org/pub/linux/games/amusements/fortune/fortune-mod-9708.bin"
	fortune_cmd_bin := exec.Command("curl", "-L", fortune_bin_url, "-o", "binary.bin")
	err = fortune_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fortune_src_url := "https://www.ibiblio.org/pub/linux/games/amusements/fortune/fortune-mod-9708.src.tar.gz"
	fortune_cmd_src := exec.Command("curl", "-L", fortune_src_url, "-o", "source.tar.gz")
	err = fortune_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fortune_cmd_direct := exec.Command("./binary")
	err = fortune_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

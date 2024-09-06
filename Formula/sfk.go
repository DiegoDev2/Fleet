package main

// Sfk - Command-line tools collection
// Homepage: http://stahlworks.com/dev/swiss-file-knife.html

import (
	"fmt"
	
	"os/exec"
)

func installSfk() {
	// Método 1: Descargar y extraer .tar.gz
	sfk_tar_url := "https://downloads.sourceforge.net/project/swissfileknife/1-swissfileknife/1.9.9.0/sfk-1.9.9.tar.gz"
	sfk_cmd_tar := exec.Command("curl", "-L", sfk_tar_url, "-o", "package.tar.gz")
	err := sfk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sfk_zip_url := "https://downloads.sourceforge.net/project/swissfileknife/1-swissfileknife/1.9.9.0/sfk-1.9.9.zip"
	sfk_cmd_zip := exec.Command("curl", "-L", sfk_zip_url, "-o", "package.zip")
	err = sfk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sfk_bin_url := "https://downloads.sourceforge.net/project/swissfileknife/1-swissfileknife/1.9.9.0/sfk-1.9.9.bin"
	sfk_cmd_bin := exec.Command("curl", "-L", sfk_bin_url, "-o", "binary.bin")
	err = sfk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sfk_src_url := "https://downloads.sourceforge.net/project/swissfileknife/1-swissfileknife/1.9.9.0/sfk-1.9.9.src.tar.gz"
	sfk_cmd_src := exec.Command("curl", "-L", sfk_src_url, "-o", "source.tar.gz")
	err = sfk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sfk_cmd_direct := exec.Command("./binary")
	err = sfk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

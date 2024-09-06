package main

// GitNumber - Use numbers for dealing with files in git
// Homepage: https://github.com/holygeek/git-number

import (
	"fmt"
	
	"os/exec"
)

func installGitNumber() {
	// Método 1: Descargar y extraer .tar.gz
	gitnumber_tar_url := "https://github.com/holygeek/git-number/archive/refs/tags/1.0.1.tar.gz"
	gitnumber_cmd_tar := exec.Command("curl", "-L", gitnumber_tar_url, "-o", "package.tar.gz")
	err := gitnumber_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitnumber_zip_url := "https://github.com/holygeek/git-number/archive/refs/tags/1.0.1.zip"
	gitnumber_cmd_zip := exec.Command("curl", "-L", gitnumber_zip_url, "-o", "package.zip")
	err = gitnumber_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitnumber_bin_url := "https://github.com/holygeek/git-number/archive/refs/tags/1.0.1.bin"
	gitnumber_cmd_bin := exec.Command("curl", "-L", gitnumber_bin_url, "-o", "binary.bin")
	err = gitnumber_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitnumber_src_url := "https://github.com/holygeek/git-number/archive/refs/tags/1.0.1.src.tar.gz"
	gitnumber_cmd_src := exec.Command("curl", "-L", gitnumber_src_url, "-o", "source.tar.gz")
	err = gitnumber_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitnumber_cmd_direct := exec.Command("./binary")
	err = gitnumber_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

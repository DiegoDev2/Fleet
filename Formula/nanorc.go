package main

// Nanorc - Improved Nano Syntax Highlighting Files
// Homepage: https://github.com/scopatz/nanorc

import (
	"fmt"
	
	"os/exec"
)

func installNanorc() {
	// Método 1: Descargar y extraer .tar.gz
	nanorc_tar_url := "https://github.com/scopatz/nanorc/releases/download/2020.10.10/nanorc-2020.10.10.tar.gz"
	nanorc_cmd_tar := exec.Command("curl", "-L", nanorc_tar_url, "-o", "package.tar.gz")
	err := nanorc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nanorc_zip_url := "https://github.com/scopatz/nanorc/releases/download/2020.10.10/nanorc-2020.10.10.zip"
	nanorc_cmd_zip := exec.Command("curl", "-L", nanorc_zip_url, "-o", "package.zip")
	err = nanorc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nanorc_bin_url := "https://github.com/scopatz/nanorc/releases/download/2020.10.10/nanorc-2020.10.10.bin"
	nanorc_cmd_bin := exec.Command("curl", "-L", nanorc_bin_url, "-o", "binary.bin")
	err = nanorc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nanorc_src_url := "https://github.com/scopatz/nanorc/releases/download/2020.10.10/nanorc-2020.10.10.src.tar.gz"
	nanorc_cmd_src := exec.Command("curl", "-L", nanorc_src_url, "-o", "source.tar.gz")
	err = nanorc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nanorc_cmd_direct := exec.Command("./binary")
	err = nanorc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

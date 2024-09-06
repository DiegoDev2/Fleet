package main

// Entr - Run arbitrary commands when files change
// Homepage: https://eradman.com/entrproject/

import (
	"fmt"
	
	"os/exec"
)

func installEntr() {
	// Método 1: Descargar y extraer .tar.gz
	entr_tar_url := "https://eradman.com/entrproject/code/entr-5.6.tar.gz"
	entr_cmd_tar := exec.Command("curl", "-L", entr_tar_url, "-o", "package.tar.gz")
	err := entr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	entr_zip_url := "https://eradman.com/entrproject/code/entr-5.6.zip"
	entr_cmd_zip := exec.Command("curl", "-L", entr_zip_url, "-o", "package.zip")
	err = entr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	entr_bin_url := "https://eradman.com/entrproject/code/entr-5.6.bin"
	entr_cmd_bin := exec.Command("curl", "-L", entr_bin_url, "-o", "binary.bin")
	err = entr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	entr_src_url := "https://eradman.com/entrproject/code/entr-5.6.src.tar.gz"
	entr_cmd_src := exec.Command("curl", "-L", entr_src_url, "-o", "source.tar.gz")
	err = entr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	entr_cmd_direct := exec.Command("./binary")
	err = entr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

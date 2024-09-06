package main

// Lout - Text formatting like TeX, but simpler
// Homepage: https://savannah.nongnu.org/projects/lout

import (
	"fmt"
	
	"os/exec"
)

func installLout() {
	// Método 1: Descargar y extraer .tar.gz
	lout_tar_url := "https://github.com/william8000/lout/archive/refs/tags/3.43.tar.gz"
	lout_cmd_tar := exec.Command("curl", "-L", lout_tar_url, "-o", "package.tar.gz")
	err := lout_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lout_zip_url := "https://github.com/william8000/lout/archive/refs/tags/3.43.zip"
	lout_cmd_zip := exec.Command("curl", "-L", lout_zip_url, "-o", "package.zip")
	err = lout_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lout_bin_url := "https://github.com/william8000/lout/archive/refs/tags/3.43.bin"
	lout_cmd_bin := exec.Command("curl", "-L", lout_bin_url, "-o", "binary.bin")
	err = lout_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lout_src_url := "https://github.com/william8000/lout/archive/refs/tags/3.43.src.tar.gz"
	lout_cmd_src := exec.Command("curl", "-L", lout_src_url, "-o", "source.tar.gz")
	err = lout_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lout_cmd_direct := exec.Command("./binary")
	err = lout_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

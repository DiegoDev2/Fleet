package main

// Willgit - William's miscellaneous git tools
// Homepage: https://github.com/DanielVartanov/willgit

import (
	"fmt"
	
	"os/exec"
)

func installWillgit() {
	// Método 1: Descargar y extraer .tar.gz
	willgit_tar_url := "https://github.com/DanielVartanov/willgit/archive/refs/tags/1.0.0.tar.gz"
	willgit_cmd_tar := exec.Command("curl", "-L", willgit_tar_url, "-o", "package.tar.gz")
	err := willgit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	willgit_zip_url := "https://github.com/DanielVartanov/willgit/archive/refs/tags/1.0.0.zip"
	willgit_cmd_zip := exec.Command("curl", "-L", willgit_zip_url, "-o", "package.zip")
	err = willgit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	willgit_bin_url := "https://github.com/DanielVartanov/willgit/archive/refs/tags/1.0.0.bin"
	willgit_cmd_bin := exec.Command("curl", "-L", willgit_bin_url, "-o", "binary.bin")
	err = willgit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	willgit_src_url := "https://github.com/DanielVartanov/willgit/archive/refs/tags/1.0.0.src.tar.gz"
	willgit_cmd_src := exec.Command("curl", "-L", willgit_src_url, "-o", "source.tar.gz")
	err = willgit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	willgit_cmd_direct := exec.Command("./binary")
	err = willgit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

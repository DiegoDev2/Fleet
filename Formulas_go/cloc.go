package main

// Cloc - Statistics utility to count lines of code
// Homepage: https://github.com/AlDanial/cloc/

import (
	"fmt"
	
	"os/exec"
)

func installCloc() {
	// Método 1: Descargar y extraer .tar.gz
	cloc_tar_url := "https://github.com/AlDanial/cloc/archive/refs/tags/v2.02.tar.gz"
	cloc_cmd_tar := exec.Command("curl", "-L", cloc_tar_url, "-o", "package.tar.gz")
	err := cloc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cloc_zip_url := "https://github.com/AlDanial/cloc/archive/refs/tags/v2.02.zip"
	cloc_cmd_zip := exec.Command("curl", "-L", cloc_zip_url, "-o", "package.zip")
	err = cloc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cloc_bin_url := "https://github.com/AlDanial/cloc/archive/refs/tags/v2.02.bin"
	cloc_cmd_bin := exec.Command("curl", "-L", cloc_bin_url, "-o", "binary.bin")
	err = cloc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cloc_src_url := "https://github.com/AlDanial/cloc/archive/refs/tags/v2.02.src.tar.gz"
	cloc_cmd_src := exec.Command("curl", "-L", cloc_src_url, "-o", "source.tar.gz")
	err = cloc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cloc_cmd_direct := exec.Command("./binary")
	err = cloc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

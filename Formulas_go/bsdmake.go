package main

// Bsdmake - BSD version of the Make build tool
// Homepage: https://opensource.apple.com/

import (
	"fmt"
	
	"os/exec"
)

func installBsdmake() {
	// Método 1: Descargar y extraer .tar.gz
	bsdmake_tar_url := "https://github.com/apple-oss-distributions/bsdmake/archive/refs/tags/bsdmake-24.tar.gz"
	bsdmake_cmd_tar := exec.Command("curl", "-L", bsdmake_tar_url, "-o", "package.tar.gz")
	err := bsdmake_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bsdmake_zip_url := "https://github.com/apple-oss-distributions/bsdmake/archive/refs/tags/bsdmake-24.zip"
	bsdmake_cmd_zip := exec.Command("curl", "-L", bsdmake_zip_url, "-o", "package.zip")
	err = bsdmake_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bsdmake_bin_url := "https://github.com/apple-oss-distributions/bsdmake/archive/refs/tags/bsdmake-24.bin"
	bsdmake_cmd_bin := exec.Command("curl", "-L", bsdmake_bin_url, "-o", "binary.bin")
	err = bsdmake_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bsdmake_src_url := "https://github.com/apple-oss-distributions/bsdmake/archive/refs/tags/bsdmake-24.src.tar.gz"
	bsdmake_cmd_src := exec.Command("curl", "-L", bsdmake_src_url, "-o", "source.tar.gz")
	err = bsdmake_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bsdmake_cmd_direct := exec.Command("./binary")
	err = bsdmake_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

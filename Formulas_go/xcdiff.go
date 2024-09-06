package main

// Xcdiff - Tool to diff xcodeproj files
// Homepage: https://github.com/bloomberg/xcdiff

import (
	"fmt"
	
	"os/exec"
)

func installXcdiff() {
	// Método 1: Descargar y extraer .tar.gz
	xcdiff_tar_url := "https://github.com/bloomberg/xcdiff.git"
	xcdiff_cmd_tar := exec.Command("curl", "-L", xcdiff_tar_url, "-o", "package.tar.gz")
	err := xcdiff_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xcdiff_zip_url := "https://github.com/bloomberg/xcdiff.git"
	xcdiff_cmd_zip := exec.Command("curl", "-L", xcdiff_zip_url, "-o", "package.zip")
	err = xcdiff_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xcdiff_bin_url := "https://github.com/bloomberg/xcdiff.git"
	xcdiff_cmd_bin := exec.Command("curl", "-L", xcdiff_bin_url, "-o", "binary.bin")
	err = xcdiff_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xcdiff_src_url := "https://github.com/bloomberg/xcdiff.git"
	xcdiff_cmd_src := exec.Command("curl", "-L", xcdiff_src_url, "-o", "source.tar.gz")
	err = xcdiff_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xcdiff_cmd_direct := exec.Command("./binary")
	err = xcdiff_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

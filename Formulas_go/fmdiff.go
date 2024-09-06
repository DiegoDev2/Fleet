package main

// Fmdiff - Use FileMerge as a diff command for Subversion and Mercurial
// Homepage: https://github.com/brunodefraine/fmscripts

import (
	"fmt"
	
	"os/exec"
)

func installFmdiff() {
	// Método 1: Descargar y extraer .tar.gz
	fmdiff_tar_url := "https://github.com/brunodefraine/fmscripts/archive/refs/tags/20150915.tar.gz"
	fmdiff_cmd_tar := exec.Command("curl", "-L", fmdiff_tar_url, "-o", "package.tar.gz")
	err := fmdiff_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fmdiff_zip_url := "https://github.com/brunodefraine/fmscripts/archive/refs/tags/20150915.zip"
	fmdiff_cmd_zip := exec.Command("curl", "-L", fmdiff_zip_url, "-o", "package.zip")
	err = fmdiff_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fmdiff_bin_url := "https://github.com/brunodefraine/fmscripts/archive/refs/tags/20150915.bin"
	fmdiff_cmd_bin := exec.Command("curl", "-L", fmdiff_bin_url, "-o", "binary.bin")
	err = fmdiff_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fmdiff_src_url := "https://github.com/brunodefraine/fmscripts/archive/refs/tags/20150915.src.tar.gz"
	fmdiff_cmd_src := exec.Command("curl", "-L", fmdiff_src_url, "-o", "source.tar.gz")
	err = fmdiff_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fmdiff_cmd_direct := exec.Command("./binary")
	err = fmdiff_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

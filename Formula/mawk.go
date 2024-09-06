package main

// Mawk - Interpreter for the AWK Programming Language
// Homepage: https://invisible-island.net/mawk/

import (
	"fmt"
	
	"os/exec"
)

func installMawk() {
	// Método 1: Descargar y extraer .tar.gz
	mawk_tar_url := "https://invisible-mirror.net/archives/mawk/mawk-1.3.4-20240827.tgz"
	mawk_cmd_tar := exec.Command("curl", "-L", mawk_tar_url, "-o", "package.tar.gz")
	err := mawk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mawk_zip_url := "https://invisible-mirror.net/archives/mawk/mawk-1.3.4-20240827.tgz"
	mawk_cmd_zip := exec.Command("curl", "-L", mawk_zip_url, "-o", "package.zip")
	err = mawk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mawk_bin_url := "https://invisible-mirror.net/archives/mawk/mawk-1.3.4-20240827.tgz"
	mawk_cmd_bin := exec.Command("curl", "-L", mawk_bin_url, "-o", "binary.bin")
	err = mawk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mawk_src_url := "https://invisible-mirror.net/archives/mawk/mawk-1.3.4-20240827.tgz"
	mawk_cmd_src := exec.Command("curl", "-L", mawk_src_url, "-o", "source.tar.gz")
	err = mawk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mawk_cmd_direct := exec.Command("./binary")
	err = mawk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

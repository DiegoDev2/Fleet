package main

// Tmpwatch - Find and remove files not accessed in a specified time
// Homepage: https://pagure.io/tmpwatch

import (
	"fmt"
	
	"os/exec"
)

func installTmpwatch() {
	// Método 1: Descargar y extraer .tar.gz
	tmpwatch_tar_url := "https://releases.pagure.org/tmpwatch/tmpwatch-2.11.tar.bz2"
	tmpwatch_cmd_tar := exec.Command("curl", "-L", tmpwatch_tar_url, "-o", "package.tar.gz")
	err := tmpwatch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tmpwatch_zip_url := "https://releases.pagure.org/tmpwatch/tmpwatch-2.11.tar.bz2"
	tmpwatch_cmd_zip := exec.Command("curl", "-L", tmpwatch_zip_url, "-o", "package.zip")
	err = tmpwatch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tmpwatch_bin_url := "https://releases.pagure.org/tmpwatch/tmpwatch-2.11.tar.bz2"
	tmpwatch_cmd_bin := exec.Command("curl", "-L", tmpwatch_bin_url, "-o", "binary.bin")
	err = tmpwatch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tmpwatch_src_url := "https://releases.pagure.org/tmpwatch/tmpwatch-2.11.tar.bz2"
	tmpwatch_cmd_src := exec.Command("curl", "-L", tmpwatch_src_url, "-o", "source.tar.gz")
	err = tmpwatch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tmpwatch_cmd_direct := exec.Command("./binary")
	err = tmpwatch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

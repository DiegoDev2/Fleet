package main

// Empty - Lightweight Expect-like PTY tool for shell scripts
// Homepage: https://empty.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installEmpty() {
	// Método 1: Descargar y extraer .tar.gz
	empty_tar_url := "https://downloads.sourceforge.net/project/empty/empty/empty-0.6.23c/empty-0.6.23c.tgz"
	empty_cmd_tar := exec.Command("curl", "-L", empty_tar_url, "-o", "package.tar.gz")
	err := empty_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	empty_zip_url := "https://downloads.sourceforge.net/project/empty/empty/empty-0.6.23c/empty-0.6.23c.tgz"
	empty_cmd_zip := exec.Command("curl", "-L", empty_zip_url, "-o", "package.zip")
	err = empty_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	empty_bin_url := "https://downloads.sourceforge.net/project/empty/empty/empty-0.6.23c/empty-0.6.23c.tgz"
	empty_cmd_bin := exec.Command("curl", "-L", empty_bin_url, "-o", "binary.bin")
	err = empty_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	empty_src_url := "https://downloads.sourceforge.net/project/empty/empty/empty-0.6.23c/empty-0.6.23c.tgz"
	empty_cmd_src := exec.Command("curl", "-L", empty_src_url, "-o", "source.tar.gz")
	err = empty_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	empty_cmd_direct := exec.Command("./binary")
	err = empty_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

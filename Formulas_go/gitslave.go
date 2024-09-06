package main

// Gitslave - Create group of related repos with one as superproject
// Homepage: https://gitslave.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installGitslave() {
	// Método 1: Descargar y extraer .tar.gz
	gitslave_tar_url := "https://downloads.sourceforge.net/project/gitslave/gitslave-2.0.2.tar.gz"
	gitslave_cmd_tar := exec.Command("curl", "-L", gitslave_tar_url, "-o", "package.tar.gz")
	err := gitslave_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitslave_zip_url := "https://downloads.sourceforge.net/project/gitslave/gitslave-2.0.2.zip"
	gitslave_cmd_zip := exec.Command("curl", "-L", gitslave_zip_url, "-o", "package.zip")
	err = gitslave_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitslave_bin_url := "https://downloads.sourceforge.net/project/gitslave/gitslave-2.0.2.bin"
	gitslave_cmd_bin := exec.Command("curl", "-L", gitslave_bin_url, "-o", "binary.bin")
	err = gitslave_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitslave_src_url := "https://downloads.sourceforge.net/project/gitslave/gitslave-2.0.2.src.tar.gz"
	gitslave_cmd_src := exec.Command("curl", "-L", gitslave_src_url, "-o", "source.tar.gz")
	err = gitslave_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitslave_cmd_direct := exec.Command("./binary")
	err = gitslave_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

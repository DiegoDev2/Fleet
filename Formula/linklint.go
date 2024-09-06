package main

// Linklint - Link checker and web site maintenance tool
// Homepage: http://linklint.org

import (
	"fmt"
	
	"os/exec"
)

func installLinklint() {
	// Método 1: Descargar y extraer .tar.gz
	linklint_tar_url := "http://linklint.org/download/linklint-2.3.5.tar.gz"
	linklint_cmd_tar := exec.Command("curl", "-L", linklint_tar_url, "-o", "package.tar.gz")
	err := linklint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	linklint_zip_url := "http://linklint.org/download/linklint-2.3.5.zip"
	linklint_cmd_zip := exec.Command("curl", "-L", linklint_zip_url, "-o", "package.zip")
	err = linklint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	linklint_bin_url := "http://linklint.org/download/linklint-2.3.5.bin"
	linklint_cmd_bin := exec.Command("curl", "-L", linklint_bin_url, "-o", "binary.bin")
	err = linklint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	linklint_src_url := "http://linklint.org/download/linklint-2.3.5.src.tar.gz"
	linklint_cmd_src := exec.Command("curl", "-L", linklint_src_url, "-o", "source.tar.gz")
	err = linklint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	linklint_cmd_direct := exec.Command("./binary")
	err = linklint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

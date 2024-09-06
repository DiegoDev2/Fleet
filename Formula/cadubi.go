package main

// Cadubi - Creative ASCII drawing utility
// Homepage: https://github.com/statico/cadubi/

import (
	"fmt"
	
	"os/exec"
)

func installCadubi() {
	// Método 1: Descargar y extraer .tar.gz
	cadubi_tar_url := "https://github.com/statico/cadubi/archive/refs/tags/v1.3.4.tar.gz"
	cadubi_cmd_tar := exec.Command("curl", "-L", cadubi_tar_url, "-o", "package.tar.gz")
	err := cadubi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cadubi_zip_url := "https://github.com/statico/cadubi/archive/refs/tags/v1.3.4.zip"
	cadubi_cmd_zip := exec.Command("curl", "-L", cadubi_zip_url, "-o", "package.zip")
	err = cadubi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cadubi_bin_url := "https://github.com/statico/cadubi/archive/refs/tags/v1.3.4.bin"
	cadubi_cmd_bin := exec.Command("curl", "-L", cadubi_bin_url, "-o", "binary.bin")
	err = cadubi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cadubi_src_url := "https://github.com/statico/cadubi/archive/refs/tags/v1.3.4.src.tar.gz"
	cadubi_cmd_src := exec.Command("curl", "-L", cadubi_src_url, "-o", "source.tar.gz")
	err = cadubi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cadubi_cmd_direct := exec.Command("./binary")
	err = cadubi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

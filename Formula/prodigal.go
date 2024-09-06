package main

// Prodigal - Microbial gene prediction
// Homepage: https://github.com/hyattpd/Prodigal

import (
	"fmt"
	
	"os/exec"
)

func installProdigal() {
	// Método 1: Descargar y extraer .tar.gz
	prodigal_tar_url := "https://github.com/hyattpd/Prodigal/archive/refs/tags/v2.6.3.tar.gz"
	prodigal_cmd_tar := exec.Command("curl", "-L", prodigal_tar_url, "-o", "package.tar.gz")
	err := prodigal_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	prodigal_zip_url := "https://github.com/hyattpd/Prodigal/archive/refs/tags/v2.6.3.zip"
	prodigal_cmd_zip := exec.Command("curl", "-L", prodigal_zip_url, "-o", "package.zip")
	err = prodigal_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	prodigal_bin_url := "https://github.com/hyattpd/Prodigal/archive/refs/tags/v2.6.3.bin"
	prodigal_cmd_bin := exec.Command("curl", "-L", prodigal_bin_url, "-o", "binary.bin")
	err = prodigal_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	prodigal_src_url := "https://github.com/hyattpd/Prodigal/archive/refs/tags/v2.6.3.src.tar.gz"
	prodigal_cmd_src := exec.Command("curl", "-L", prodigal_src_url, "-o", "source.tar.gz")
	err = prodigal_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	prodigal_cmd_direct := exec.Command("./binary")
	err = prodigal_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

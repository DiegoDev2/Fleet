package main

// Foreman - Manage Procfile-based applications
// Homepage: https://ddollar.github.io/foreman/

import (
	"fmt"
	
	"os/exec"
)

func installForeman() {
	// Método 1: Descargar y extraer .tar.gz
	foreman_tar_url := "https://github.com/ddollar/foreman/archive/refs/tags/v0.88.1.tar.gz"
	foreman_cmd_tar := exec.Command("curl", "-L", foreman_tar_url, "-o", "package.tar.gz")
	err := foreman_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	foreman_zip_url := "https://github.com/ddollar/foreman/archive/refs/tags/v0.88.1.zip"
	foreman_cmd_zip := exec.Command("curl", "-L", foreman_zip_url, "-o", "package.zip")
	err = foreman_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	foreman_bin_url := "https://github.com/ddollar/foreman/archive/refs/tags/v0.88.1.bin"
	foreman_cmd_bin := exec.Command("curl", "-L", foreman_bin_url, "-o", "binary.bin")
	err = foreman_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	foreman_src_url := "https://github.com/ddollar/foreman/archive/refs/tags/v0.88.1.src.tar.gz"
	foreman_cmd_src := exec.Command("curl", "-L", foreman_src_url, "-o", "source.tar.gz")
	err = foreman_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	foreman_cmd_direct := exec.Command("./binary")
	err = foreman_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

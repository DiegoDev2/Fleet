package main

// Click - Command-line interactive controller for Kubernetes
// Homepage: https://github.com/databricks/click

import (
	"fmt"
	
	"os/exec"
)

func installClick() {
	// Método 1: Descargar y extraer .tar.gz
	click_tar_url := "https://github.com/databricks/click/archive/refs/tags/v0.6.3.tar.gz"
	click_cmd_tar := exec.Command("curl", "-L", click_tar_url, "-o", "package.tar.gz")
	err := click_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	click_zip_url := "https://github.com/databricks/click/archive/refs/tags/v0.6.3.zip"
	click_cmd_zip := exec.Command("curl", "-L", click_zip_url, "-o", "package.zip")
	err = click_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	click_bin_url := "https://github.com/databricks/click/archive/refs/tags/v0.6.3.bin"
	click_cmd_bin := exec.Command("curl", "-L", click_bin_url, "-o", "binary.bin")
	err = click_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	click_src_url := "https://github.com/databricks/click/archive/refs/tags/v0.6.3.src.tar.gz"
	click_cmd_src := exec.Command("curl", "-L", click_src_url, "-o", "source.tar.gz")
	err = click_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	click_cmd_direct := exec.Command("./binary")
	err = click_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}

package main

// Keploy - Testing Toolkit creates test-cases and data mocks from API calls, DB queries
// Homepage: https://keploy.io

import (
	"fmt"
	
	"os/exec"
)

func installKeploy() {
	// Método 1: Descargar y extraer .tar.gz
	keploy_tar_url := "https://github.com/keploy/keploy/archive/refs/tags/v0.9.1.tar.gz"
	keploy_cmd_tar := exec.Command("curl", "-L", keploy_tar_url, "-o", "package.tar.gz")
	err := keploy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	keploy_zip_url := "https://github.com/keploy/keploy/archive/refs/tags/v0.9.1.zip"
	keploy_cmd_zip := exec.Command("curl", "-L", keploy_zip_url, "-o", "package.zip")
	err = keploy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	keploy_bin_url := "https://github.com/keploy/keploy/archive/refs/tags/v0.9.1.bin"
	keploy_cmd_bin := exec.Command("curl", "-L", keploy_bin_url, "-o", "binary.bin")
	err = keploy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	keploy_src_url := "https://github.com/keploy/keploy/archive/refs/tags/v0.9.1.src.tar.gz"
	keploy_cmd_src := exec.Command("curl", "-L", keploy_src_url, "-o", "source.tar.gz")
	err = keploy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	keploy_cmd_direct := exec.Command("./binary")
	err = keploy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gatsby-cli")
	exec.Command("latte", "install", "gatsby-cli").Run()
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}

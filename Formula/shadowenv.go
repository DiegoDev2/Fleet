package main

// Shadowenv - Reversible directory-local environment variable manipulations
// Homepage: https://shopify.github.io/shadowenv/

import (
	"fmt"
	
	"os/exec"
)

func installShadowenv() {
	// Método 1: Descargar y extraer .tar.gz
	shadowenv_tar_url := "https://github.com/Shopify/shadowenv/archive/refs/tags/2.1.2.tar.gz"
	shadowenv_cmd_tar := exec.Command("curl", "-L", shadowenv_tar_url, "-o", "package.tar.gz")
	err := shadowenv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	shadowenv_zip_url := "https://github.com/Shopify/shadowenv/archive/refs/tags/2.1.2.zip"
	shadowenv_cmd_zip := exec.Command("curl", "-L", shadowenv_zip_url, "-o", "package.zip")
	err = shadowenv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	shadowenv_bin_url := "https://github.com/Shopify/shadowenv/archive/refs/tags/2.1.2.bin"
	shadowenv_cmd_bin := exec.Command("curl", "-L", shadowenv_bin_url, "-o", "binary.bin")
	err = shadowenv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	shadowenv_src_url := "https://github.com/Shopify/shadowenv/archive/refs/tags/2.1.2.src.tar.gz"
	shadowenv_cmd_src := exec.Command("curl", "-L", shadowenv_src_url, "-o", "source.tar.gz")
	err = shadowenv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	shadowenv_cmd_direct := exec.Command("./binary")
	err = shadowenv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}

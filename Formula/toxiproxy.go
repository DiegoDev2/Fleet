package main

// Toxiproxy - TCP proxy to simulate network & system conditions for chaos & resiliency testing
// Homepage: https://github.com/shopify/toxiproxy

import (
	"fmt"
	
	"os/exec"
)

func installToxiproxy() {
	// Método 1: Descargar y extraer .tar.gz
	toxiproxy_tar_url := "https://github.com/Shopify/toxiproxy/archive/refs/tags/v2.9.0.tar.gz"
	toxiproxy_cmd_tar := exec.Command("curl", "-L", toxiproxy_tar_url, "-o", "package.tar.gz")
	err := toxiproxy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	toxiproxy_zip_url := "https://github.com/Shopify/toxiproxy/archive/refs/tags/v2.9.0.zip"
	toxiproxy_cmd_zip := exec.Command("curl", "-L", toxiproxy_zip_url, "-o", "package.zip")
	err = toxiproxy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	toxiproxy_bin_url := "https://github.com/Shopify/toxiproxy/archive/refs/tags/v2.9.0.bin"
	toxiproxy_cmd_bin := exec.Command("curl", "-L", toxiproxy_bin_url, "-o", "binary.bin")
	err = toxiproxy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	toxiproxy_src_url := "https://github.com/Shopify/toxiproxy/archive/refs/tags/v2.9.0.src.tar.gz"
	toxiproxy_cmd_src := exec.Command("curl", "-L", toxiproxy_src_url, "-o", "source.tar.gz")
	err = toxiproxy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	toxiproxy_cmd_direct := exec.Command("./binary")
	err = toxiproxy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}

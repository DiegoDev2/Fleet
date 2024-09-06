package main

// Proxytunnel - Create TCP tunnels through HTTPS proxies
// Homepage: https://github.com/proxytunnel/proxytunnel

import (
	"fmt"
	
	"os/exec"
)

func installProxytunnel() {
	// Método 1: Descargar y extraer .tar.gz
	proxytunnel_tar_url := "https://github.com/proxytunnel/proxytunnel/archive/refs/tags/v1.12.2.tar.gz"
	proxytunnel_cmd_tar := exec.Command("curl", "-L", proxytunnel_tar_url, "-o", "package.tar.gz")
	err := proxytunnel_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	proxytunnel_zip_url := "https://github.com/proxytunnel/proxytunnel/archive/refs/tags/v1.12.2.zip"
	proxytunnel_cmd_zip := exec.Command("curl", "-L", proxytunnel_zip_url, "-o", "package.zip")
	err = proxytunnel_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	proxytunnel_bin_url := "https://github.com/proxytunnel/proxytunnel/archive/refs/tags/v1.12.2.bin"
	proxytunnel_cmd_bin := exec.Command("curl", "-L", proxytunnel_bin_url, "-o", "binary.bin")
	err = proxytunnel_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	proxytunnel_src_url := "https://github.com/proxytunnel/proxytunnel/archive/refs/tags/v1.12.2.src.tar.gz"
	proxytunnel_cmd_src := exec.Command("curl", "-L", proxytunnel_src_url, "-o", "source.tar.gz")
	err = proxytunnel_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	proxytunnel_cmd_direct := exec.Command("./binary")
	err = proxytunnel_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asciidoc")
	exec.Command("latte", "install", "asciidoc").Run()
	fmt.Println("Instalando dependencia: xmlto")
	exec.Command("latte", "install", "xmlto").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}

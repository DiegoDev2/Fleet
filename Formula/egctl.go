package main

// Egctl - Command-line utility for operating Envoy Gateway
// Homepage: https://gateway.envoyproxy.io/

import (
	"fmt"
	
	"os/exec"
)

func installEgctl() {
	// Método 1: Descargar y extraer .tar.gz
	egctl_tar_url := "https://github.com/envoyproxy/gateway/archive/refs/tags/v1.1.0.tar.gz"
	egctl_cmd_tar := exec.Command("curl", "-L", egctl_tar_url, "-o", "package.tar.gz")
	err := egctl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	egctl_zip_url := "https://github.com/envoyproxy/gateway/archive/refs/tags/v1.1.0.zip"
	egctl_cmd_zip := exec.Command("curl", "-L", egctl_zip_url, "-o", "package.zip")
	err = egctl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	egctl_bin_url := "https://github.com/envoyproxy/gateway/archive/refs/tags/v1.1.0.bin"
	egctl_cmd_bin := exec.Command("curl", "-L", egctl_bin_url, "-o", "binary.bin")
	err = egctl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	egctl_src_url := "https://github.com/envoyproxy/gateway/archive/refs/tags/v1.1.0.src.tar.gz"
	egctl_cmd_src := exec.Command("curl", "-L", egctl_src_url, "-o", "source.tar.gz")
	err = egctl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	egctl_cmd_direct := exec.Command("./binary")
	err = egctl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: btrfs-progs")
	exec.Command("latte", "install", "btrfs-progs").Run()
}

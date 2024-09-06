package main

// Hcloud - Command-line interface for Hetzner Cloud
// Homepage: https://github.com/hetznercloud/cli

import (
	"fmt"
	
	"os/exec"
)

func installHcloud() {
	// Método 1: Descargar y extraer .tar.gz
	hcloud_tar_url := "https://github.com/hetznercloud/cli/archive/refs/tags/v1.47.0.tar.gz"
	hcloud_cmd_tar := exec.Command("curl", "-L", hcloud_tar_url, "-o", "package.tar.gz")
	err := hcloud_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hcloud_zip_url := "https://github.com/hetznercloud/cli/archive/refs/tags/v1.47.0.zip"
	hcloud_cmd_zip := exec.Command("curl", "-L", hcloud_zip_url, "-o", "package.zip")
	err = hcloud_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hcloud_bin_url := "https://github.com/hetznercloud/cli/archive/refs/tags/v1.47.0.bin"
	hcloud_cmd_bin := exec.Command("curl", "-L", hcloud_bin_url, "-o", "binary.bin")
	err = hcloud_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hcloud_src_url := "https://github.com/hetznercloud/cli/archive/refs/tags/v1.47.0.src.tar.gz"
	hcloud_cmd_src := exec.Command("curl", "-L", hcloud_src_url, "-o", "source.tar.gz")
	err = hcloud_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hcloud_cmd_direct := exec.Command("./binary")
	err = hcloud_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}

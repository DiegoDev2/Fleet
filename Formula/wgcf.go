package main

// Wgcf - Generate WireGuard profile from Cloudflare Warp account
// Homepage: https://github.com/ViRb3/wgcf

import (
	"fmt"
	
	"os/exec"
)

func installWgcf() {
	// Método 1: Descargar y extraer .tar.gz
	wgcf_tar_url := "https://github.com/ViRb3/wgcf/archive/refs/tags/v2.2.22.tar.gz"
	wgcf_cmd_tar := exec.Command("curl", "-L", wgcf_tar_url, "-o", "package.tar.gz")
	err := wgcf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wgcf_zip_url := "https://github.com/ViRb3/wgcf/archive/refs/tags/v2.2.22.zip"
	wgcf_cmd_zip := exec.Command("curl", "-L", wgcf_zip_url, "-o", "package.zip")
	err = wgcf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wgcf_bin_url := "https://github.com/ViRb3/wgcf/archive/refs/tags/v2.2.22.bin"
	wgcf_cmd_bin := exec.Command("curl", "-L", wgcf_bin_url, "-o", "binary.bin")
	err = wgcf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wgcf_src_url := "https://github.com/ViRb3/wgcf/archive/refs/tags/v2.2.22.src.tar.gz"
	wgcf_cmd_src := exec.Command("curl", "-L", wgcf_src_url, "-o", "source.tar.gz")
	err = wgcf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wgcf_cmd_direct := exec.Command("./binary")
	err = wgcf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}

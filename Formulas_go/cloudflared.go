package main

// Cloudflared - Cloudflare Tunnel client (formerly Argo Tunnel)
// Homepage: https://developers.cloudflare.com/cloudflare-one/connections/connect-apps/install-and-setup/tunnel-guide

import (
	"fmt"
	
	"os/exec"
)

func installCloudflared() {
	// Método 1: Descargar y extraer .tar.gz
	cloudflared_tar_url := "https://github.com/cloudflare/cloudflared/archive/refs/tags/2024.8.3.tar.gz"
	cloudflared_cmd_tar := exec.Command("curl", "-L", cloudflared_tar_url, "-o", "package.tar.gz")
	err := cloudflared_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cloudflared_zip_url := "https://github.com/cloudflare/cloudflared/archive/refs/tags/2024.8.3.zip"
	cloudflared_cmd_zip := exec.Command("curl", "-L", cloudflared_zip_url, "-o", "package.zip")
	err = cloudflared_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cloudflared_bin_url := "https://github.com/cloudflare/cloudflared/archive/refs/tags/2024.8.3.bin"
	cloudflared_cmd_bin := exec.Command("curl", "-L", cloudflared_bin_url, "-o", "binary.bin")
	err = cloudflared_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cloudflared_src_url := "https://github.com/cloudflare/cloudflared/archive/refs/tags/2024.8.3.src.tar.gz"
	cloudflared_cmd_src := exec.Command("curl", "-L", cloudflared_src_url, "-o", "source.tar.gz")
	err = cloudflared_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cloudflared_cmd_direct := exec.Command("./binary")
	err = cloudflared_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}

package main

// Caddy - Powerful, enterprise-ready, open source web server with automatic HTTPS
// Homepage: https://caddyserver.com/

import (
	"fmt"
	
	"os/exec"
)

func installCaddy() {
	// Método 1: Descargar y extraer .tar.gz
	caddy_tar_url := "https://github.com/caddyserver/caddy/archive/refs/tags/v2.8.4.tar.gz"
	caddy_cmd_tar := exec.Command("curl", "-L", caddy_tar_url, "-o", "package.tar.gz")
	err := caddy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	caddy_zip_url := "https://github.com/caddyserver/caddy/archive/refs/tags/v2.8.4.zip"
	caddy_cmd_zip := exec.Command("curl", "-L", caddy_zip_url, "-o", "package.zip")
	err = caddy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	caddy_bin_url := "https://github.com/caddyserver/caddy/archive/refs/tags/v2.8.4.bin"
	caddy_cmd_bin := exec.Command("curl", "-L", caddy_bin_url, "-o", "binary.bin")
	err = caddy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	caddy_src_url := "https://github.com/caddyserver/caddy/archive/refs/tags/v2.8.4.src.tar.gz"
	caddy_cmd_src := exec.Command("curl", "-L", caddy_src_url, "-o", "source.tar.gz")
	err = caddy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	caddy_cmd_direct := exec.Command("./binary")
	err = caddy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}

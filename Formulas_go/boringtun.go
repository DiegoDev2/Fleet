package main

// Boringtun - Userspace WireGuard implementation in Rust
// Homepage: https://github.com/cloudflare/boringtun

import (
	"fmt"
	
	"os/exec"
)

func installBoringtun() {
	// Método 1: Descargar y extraer .tar.gz
	boringtun_tar_url := "https://github.com/cloudflare/boringtun/archive/refs/tags/boringtun-0.5.2.tar.gz"
	boringtun_cmd_tar := exec.Command("curl", "-L", boringtun_tar_url, "-o", "package.tar.gz")
	err := boringtun_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	boringtun_zip_url := "https://github.com/cloudflare/boringtun/archive/refs/tags/boringtun-0.5.2.zip"
	boringtun_cmd_zip := exec.Command("curl", "-L", boringtun_zip_url, "-o", "package.zip")
	err = boringtun_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	boringtun_bin_url := "https://github.com/cloudflare/boringtun/archive/refs/tags/boringtun-0.5.2.bin"
	boringtun_cmd_bin := exec.Command("curl", "-L", boringtun_bin_url, "-o", "binary.bin")
	err = boringtun_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	boringtun_src_url := "https://github.com/cloudflare/boringtun/archive/refs/tags/boringtun-0.5.2.src.tar.gz"
	boringtun_cmd_src := exec.Command("curl", "-L", boringtun_src_url, "-o", "source.tar.gz")
	err = boringtun_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	boringtun_cmd_direct := exec.Command("./binary")
	err = boringtun_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}

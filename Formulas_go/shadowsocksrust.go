package main

// ShadowsocksRust - Rust port of Shadowsocks
// Homepage: https://github.com/shadowsocks/shadowsocks-rust

import (
	"fmt"
	
	"os/exec"
)

func installShadowsocksRust() {
	// Método 1: Descargar y extraer .tar.gz
	shadowsocksrust_tar_url := "https://github.com/shadowsocks/shadowsocks-rust/archive/refs/tags/v1.20.4.tar.gz"
	shadowsocksrust_cmd_tar := exec.Command("curl", "-L", shadowsocksrust_tar_url, "-o", "package.tar.gz")
	err := shadowsocksrust_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	shadowsocksrust_zip_url := "https://github.com/shadowsocks/shadowsocks-rust/archive/refs/tags/v1.20.4.zip"
	shadowsocksrust_cmd_zip := exec.Command("curl", "-L", shadowsocksrust_zip_url, "-o", "package.zip")
	err = shadowsocksrust_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	shadowsocksrust_bin_url := "https://github.com/shadowsocks/shadowsocks-rust/archive/refs/tags/v1.20.4.bin"
	shadowsocksrust_cmd_bin := exec.Command("curl", "-L", shadowsocksrust_bin_url, "-o", "binary.bin")
	err = shadowsocksrust_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	shadowsocksrust_src_url := "https://github.com/shadowsocks/shadowsocks-rust/archive/refs/tags/v1.20.4.src.tar.gz"
	shadowsocksrust_cmd_src := exec.Command("curl", "-L", shadowsocksrust_src_url, "-o", "source.tar.gz")
	err = shadowsocksrust_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	shadowsocksrust_cmd_direct := exec.Command("./binary")
	err = shadowsocksrust_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}

package main

// ShadowsocksLibev - Libev port of shadowsocks
// Homepage: https://github.com/shadowsocks/shadowsocks-libev

import (
	"fmt"
	
	"os/exec"
)

func installShadowsocksLibev() {
	// Método 1: Descargar y extraer .tar.gz
	shadowsockslibev_tar_url := "https://github.com/shadowsocks/shadowsocks-libev/releases/download/v3.3.5/shadowsocks-libev-3.3.5.tar.gz"
	shadowsockslibev_cmd_tar := exec.Command("curl", "-L", shadowsockslibev_tar_url, "-o", "package.tar.gz")
	err := shadowsockslibev_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	shadowsockslibev_zip_url := "https://github.com/shadowsocks/shadowsocks-libev/releases/download/v3.3.5/shadowsocks-libev-3.3.5.zip"
	shadowsockslibev_cmd_zip := exec.Command("curl", "-L", shadowsockslibev_zip_url, "-o", "package.zip")
	err = shadowsockslibev_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	shadowsockslibev_bin_url := "https://github.com/shadowsocks/shadowsocks-libev/releases/download/v3.3.5/shadowsocks-libev-3.3.5.bin"
	shadowsockslibev_cmd_bin := exec.Command("curl", "-L", shadowsockslibev_bin_url, "-o", "binary.bin")
	err = shadowsockslibev_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	shadowsockslibev_src_url := "https://github.com/shadowsocks/shadowsocks-libev/releases/download/v3.3.5/shadowsocks-libev-3.3.5.src.tar.gz"
	shadowsockslibev_cmd_src := exec.Command("curl", "-L", shadowsockslibev_src_url, "-o", "source.tar.gz")
	err = shadowsockslibev_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	shadowsockslibev_cmd_direct := exec.Command("./binary")
	err = shadowsockslibev_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: asciidoc")
exec.Command("latte", "install", "asciidoc")
	fmt.Println("Instalando dependencia: xmlto")
exec.Command("latte", "install", "xmlto")
	fmt.Println("Instalando dependencia: c-ares")
exec.Command("latte", "install", "c-ares")
	fmt.Println("Instalando dependencia: libev")
exec.Command("latte", "install", "libev")
	fmt.Println("Instalando dependencia: libsodium")
exec.Command("latte", "install", "libsodium")
	fmt.Println("Instalando dependencia: mbedtls@2")
exec.Command("latte", "install", "mbedtls@2")
	fmt.Println("Instalando dependencia: pcre")
exec.Command("latte", "install", "pcre")
}

package main

// Jose - C-language implementation of Javascript Object Signing and Encryption
// Homepage: https://github.com/latchset/jose

import (
	"fmt"
	
	"os/exec"
)

func installJose() {
	// Método 1: Descargar y extraer .tar.gz
	jose_tar_url := "https://github.com/latchset/jose/releases/download/v14/jose-14.tar.xz"
	jose_cmd_tar := exec.Command("curl", "-L", jose_tar_url, "-o", "package.tar.gz")
	err := jose_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jose_zip_url := "https://github.com/latchset/jose/releases/download/v14/jose-14.tar.xz"
	jose_cmd_zip := exec.Command("curl", "-L", jose_zip_url, "-o", "package.zip")
	err = jose_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jose_bin_url := "https://github.com/latchset/jose/releases/download/v14/jose-14.tar.xz"
	jose_cmd_bin := exec.Command("curl", "-L", jose_bin_url, "-o", "binary.bin")
	err = jose_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jose_src_url := "https://github.com/latchset/jose/releases/download/v14/jose-14.tar.xz"
	jose_cmd_src := exec.Command("curl", "-L", jose_src_url, "-o", "source.tar.gz")
	err = jose_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jose_cmd_direct := exec.Command("./binary")
	err = jose_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: jansson")
exec.Command("latte", "install", "jansson")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}

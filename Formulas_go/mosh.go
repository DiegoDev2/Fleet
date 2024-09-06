package main

// Mosh - Remote terminal application
// Homepage: https://mosh.org

import (
	"fmt"
	
	"os/exec"
)

func installMosh() {
	// Método 1: Descargar y extraer .tar.gz
	mosh_tar_url := "https://github.com/mobile-shell/mosh/releases/download/mosh-1.4.0/mosh-1.4.0.tar.gz"
	mosh_cmd_tar := exec.Command("curl", "-L", mosh_tar_url, "-o", "package.tar.gz")
	err := mosh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mosh_zip_url := "https://github.com/mobile-shell/mosh/releases/download/mosh-1.4.0/mosh-1.4.0.zip"
	mosh_cmd_zip := exec.Command("curl", "-L", mosh_zip_url, "-o", "package.zip")
	err = mosh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mosh_bin_url := "https://github.com/mobile-shell/mosh/releases/download/mosh-1.4.0/mosh-1.4.0.bin"
	mosh_cmd_bin := exec.Command("curl", "-L", mosh_bin_url, "-o", "binary.bin")
	err = mosh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mosh_src_url := "https://github.com/mobile-shell/mosh/releases/download/mosh-1.4.0/mosh-1.4.0.src.tar.gz"
	mosh_cmd_src := exec.Command("curl", "-L", mosh_src_url, "-o", "source.tar.gz")
	err = mosh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mosh_cmd_direct := exec.Command("./binary")
	err = mosh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: protobuf")
exec.Command("latte", "install", "protobuf")
	fmt.Println("Instalando dependencia: tmux")
exec.Command("latte", "install", "tmux")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}

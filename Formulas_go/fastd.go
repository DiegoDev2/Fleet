package main

// Fastd - Fast and Secure Tunnelling Daemon
// Homepage: https://github.com/neocturne/fastd

import (
	"fmt"
	
	"os/exec"
)

func installFastd() {
	// Método 1: Descargar y extraer .tar.gz
	fastd_tar_url := "https://github.com/neocturne/fastd.git"
	fastd_cmd_tar := exec.Command("curl", "-L", fastd_tar_url, "-o", "package.tar.gz")
	err := fastd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fastd_zip_url := "https://github.com/neocturne/fastd.git"
	fastd_cmd_zip := exec.Command("curl", "-L", fastd_zip_url, "-o", "package.zip")
	err = fastd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fastd_bin_url := "https://github.com/neocturne/fastd.git"
	fastd_cmd_bin := exec.Command("curl", "-L", fastd_bin_url, "-o", "binary.bin")
	err = fastd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fastd_src_url := "https://github.com/neocturne/fastd.git"
	fastd_cmd_src := exec.Command("curl", "-L", fastd_src_url, "-o", "source.tar.gz")
	err = fastd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fastd_cmd_direct := exec.Command("./binary")
	err = fastd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bison")
exec.Command("latte", "install", "bison")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: json-c")
exec.Command("latte", "install", "json-c")
	fmt.Println("Instalando dependencia: libsodium")
exec.Command("latte", "install", "libsodium")
	fmt.Println("Instalando dependencia: libuecc")
exec.Command("latte", "install", "libuecc")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: libcap")
exec.Command("latte", "install", "libcap")
	fmt.Println("Instalando dependencia: libmnl")
exec.Command("latte", "install", "libmnl")
}

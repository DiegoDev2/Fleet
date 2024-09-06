package main

// Inspircd - Modular C++ Internet Relay Chat daemon
// Homepage: https://www.inspircd.org/

import (
	"fmt"
	
	"os/exec"
)

func installInspircd() {
	// Método 1: Descargar y extraer .tar.gz
	inspircd_tar_url := "https://github.com/inspircd/inspircd/archive/refs/tags/v4.2.0.tar.gz"
	inspircd_cmd_tar := exec.Command("curl", "-L", inspircd_tar_url, "-o", "package.tar.gz")
	err := inspircd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	inspircd_zip_url := "https://github.com/inspircd/inspircd/archive/refs/tags/v4.2.0.zip"
	inspircd_cmd_zip := exec.Command("curl", "-L", inspircd_zip_url, "-o", "package.zip")
	err = inspircd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	inspircd_bin_url := "https://github.com/inspircd/inspircd/archive/refs/tags/v4.2.0.bin"
	inspircd_cmd_bin := exec.Command("curl", "-L", inspircd_bin_url, "-o", "binary.bin")
	err = inspircd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	inspircd_src_url := "https://github.com/inspircd/inspircd/archive/refs/tags/v4.2.0.src.tar.gz"
	inspircd_cmd_src := exec.Command("curl", "-L", inspircd_src_url, "-o", "source.tar.gz")
	err = inspircd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	inspircd_cmd_direct := exec.Command("./binary")
	err = inspircd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: argon2")
exec.Command("latte", "install", "argon2")
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: libpq")
exec.Command("latte", "install", "libpq")
	fmt.Println("Instalando dependencia: mysql-client")
exec.Command("latte", "install", "mysql-client")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: zlib")
exec.Command("latte", "install", "zlib")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
}

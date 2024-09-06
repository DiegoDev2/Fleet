package main

// Libpqxx - C++ connector for PostgreSQL
// Homepage: https://pqxx.org/development/libpqxx/

import (
	"fmt"
	
	"os/exec"
)

func installLibpqxx() {
	// Método 1: Descargar y extraer .tar.gz
	libpqxx_tar_url := "https://github.com/jtv/libpqxx/archive/refs/tags/7.9.2.tar.gz"
	libpqxx_cmd_tar := exec.Command("curl", "-L", libpqxx_tar_url, "-o", "package.tar.gz")
	err := libpqxx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libpqxx_zip_url := "https://github.com/jtv/libpqxx/archive/refs/tags/7.9.2.zip"
	libpqxx_cmd_zip := exec.Command("curl", "-L", libpqxx_zip_url, "-o", "package.zip")
	err = libpqxx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libpqxx_bin_url := "https://github.com/jtv/libpqxx/archive/refs/tags/7.9.2.bin"
	libpqxx_cmd_bin := exec.Command("curl", "-L", libpqxx_bin_url, "-o", "binary.bin")
	err = libpqxx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libpqxx_src_url := "https://github.com/jtv/libpqxx/archive/refs/tags/7.9.2.src.tar.gz"
	libpqxx_cmd_src := exec.Command("curl", "-L", libpqxx_src_url, "-o", "source.tar.gz")
	err = libpqxx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libpqxx_cmd_direct := exec.Command("./binary")
	err = libpqxx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: xmlto")
exec.Command("latte", "install", "xmlto")
	fmt.Println("Instalando dependencia: libpq")
exec.Command("latte", "install", "libpq")
}

package main

// MariadbAT108 - Drop-in replacement for MySQL
// Homepage: https://mariadb.org/

import (
	"fmt"
	
	"os/exec"
)

func installMariadbAT108() {
	// Método 1: Descargar y extraer .tar.gz
	mariadbat108_tar_url := "https://downloads.mariadb.com/MariaDB/mariadb-10.8.8/source/mariadb-10.8.8.tar.gz"
	mariadbat108_cmd_tar := exec.Command("curl", "-L", mariadbat108_tar_url, "-o", "package.tar.gz")
	err := mariadbat108_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mariadbat108_zip_url := "https://downloads.mariadb.com/MariaDB/mariadb-10.8.8/source/mariadb-10.8.8.zip"
	mariadbat108_cmd_zip := exec.Command("curl", "-L", mariadbat108_zip_url, "-o", "package.zip")
	err = mariadbat108_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mariadbat108_bin_url := "https://downloads.mariadb.com/MariaDB/mariadb-10.8.8/source/mariadb-10.8.8.bin"
	mariadbat108_cmd_bin := exec.Command("curl", "-L", mariadbat108_bin_url, "-o", "binary.bin")
	err = mariadbat108_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mariadbat108_src_url := "https://downloads.mariadb.com/MariaDB/mariadb-10.8.8/source/mariadb-10.8.8.src.tar.gz"
	mariadbat108_cmd_src := exec.Command("curl", "-L", mariadbat108_src_url, "-o", "source.tar.gz")
	err = mariadbat108_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mariadbat108_cmd_direct := exec.Command("./binary")
	err = mariadbat108_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bison")
exec.Command("latte", "install", "bison")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: fmt")
exec.Command("latte", "install", "fmt")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: groonga")
exec.Command("latte", "install", "groonga")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
	fmt.Println("Instalando dependencia: linux-pam")
exec.Command("latte", "install", "linux-pam")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}

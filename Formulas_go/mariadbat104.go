package main

// MariadbAT104 - Drop-in replacement for MySQL
// Homepage: https://mariadb.org/

import (
	"fmt"
	
	"os/exec"
)

func installMariadbAT104() {
	// Método 1: Descargar y extraer .tar.gz
	mariadbat104_tar_url := "https://archive.mariadb.org/mariadb-10.4.34/source/mariadb-10.4.34.tar.gz"
	mariadbat104_cmd_tar := exec.Command("curl", "-L", mariadbat104_tar_url, "-o", "package.tar.gz")
	err := mariadbat104_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mariadbat104_zip_url := "https://archive.mariadb.org/mariadb-10.4.34/source/mariadb-10.4.34.zip"
	mariadbat104_cmd_zip := exec.Command("curl", "-L", mariadbat104_zip_url, "-o", "package.zip")
	err = mariadbat104_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mariadbat104_bin_url := "https://archive.mariadb.org/mariadb-10.4.34/source/mariadb-10.4.34.bin"
	mariadbat104_cmd_bin := exec.Command("curl", "-L", mariadbat104_bin_url, "-o", "binary.bin")
	err = mariadbat104_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mariadbat104_src_url := "https://archive.mariadb.org/mariadb-10.4.34/source/mariadb-10.4.34.src.tar.gz"
	mariadbat104_cmd_src := exec.Command("curl", "-L", mariadbat104_src_url, "-o", "source.tar.gz")
	err = mariadbat104_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mariadbat104_cmd_direct := exec.Command("./binary")
	err = mariadbat104_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bison")
exec.Command("latte", "install", "bison")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: groonga")
exec.Command("latte", "install", "groonga")
	fmt.Println("Instalando dependencia: openssl@1.1")
exec.Command("latte", "install", "openssl@1.1")
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
	fmt.Println("Instalando dependencia: linux-pam")
exec.Command("latte", "install", "linux-pam")
}

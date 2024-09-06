package main

// MariadbAT111 - Drop-in replacement for MySQL
// Homepage: https://mariadb.org/

import (
	"fmt"
	
	"os/exec"
)

func installMariadbAT111() {
	// Método 1: Descargar y extraer .tar.gz
	mariadbat111_tar_url := "https://archive.mariadb.org/mariadb-11.1.6/source/mariadb-11.1.6.tar.gz"
	mariadbat111_cmd_tar := exec.Command("curl", "-L", mariadbat111_tar_url, "-o", "package.tar.gz")
	err := mariadbat111_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mariadbat111_zip_url := "https://archive.mariadb.org/mariadb-11.1.6/source/mariadb-11.1.6.zip"
	mariadbat111_cmd_zip := exec.Command("curl", "-L", mariadbat111_zip_url, "-o", "package.zip")
	err = mariadbat111_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mariadbat111_bin_url := "https://archive.mariadb.org/mariadb-11.1.6/source/mariadb-11.1.6.bin"
	mariadbat111_cmd_bin := exec.Command("curl", "-L", mariadbat111_bin_url, "-o", "binary.bin")
	err = mariadbat111_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mariadbat111_src_url := "https://archive.mariadb.org/mariadb-11.1.6/source/mariadb-11.1.6.src.tar.gz"
	mariadbat111_cmd_src := exec.Command("curl", "-L", mariadbat111_src_url, "-o", "source.tar.gz")
	err = mariadbat111_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mariadbat111_cmd_direct := exec.Command("./binary")
	err = mariadbat111_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bison")
	exec.Command("latte", "install", "bison").Run()
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: fmt")
	exec.Command("latte", "install", "fmt").Run()
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: groonga")
	exec.Command("latte", "install", "groonga").Run()
	fmt.Println("Instalando dependencia: lz4")
	exec.Command("latte", "install", "lz4").Run()
	fmt.Println("Instalando dependencia: lzo")
	exec.Command("latte", "install", "lzo").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
	fmt.Println("Instalando dependencia: linux-pam")
	exec.Command("latte", "install", "linux-pam").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}

package main

// Sqlsmith - Random SQL query generator
// Homepage: https://github.com/anse1/sqlsmith

import (
	"fmt"
	
	"os/exec"
)

func installSqlsmith() {
	// Método 1: Descargar y extraer .tar.gz
	sqlsmith_tar_url := "https://github.com/anse1/sqlsmith/releases/download/v1.4/sqlsmith-1.4.tar.gz"
	sqlsmith_cmd_tar := exec.Command("curl", "-L", sqlsmith_tar_url, "-o", "package.tar.gz")
	err := sqlsmith_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sqlsmith_zip_url := "https://github.com/anse1/sqlsmith/releases/download/v1.4/sqlsmith-1.4.zip"
	sqlsmith_cmd_zip := exec.Command("curl", "-L", sqlsmith_zip_url, "-o", "package.zip")
	err = sqlsmith_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sqlsmith_bin_url := "https://github.com/anse1/sqlsmith/releases/download/v1.4/sqlsmith-1.4.bin"
	sqlsmith_cmd_bin := exec.Command("curl", "-L", sqlsmith_bin_url, "-o", "binary.bin")
	err = sqlsmith_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sqlsmith_src_url := "https://github.com/anse1/sqlsmith/releases/download/v1.4/sqlsmith-1.4.src.tar.gz"
	sqlsmith_cmd_src := exec.Command("curl", "-L", sqlsmith_src_url, "-o", "source.tar.gz")
	err = sqlsmith_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sqlsmith_cmd_direct := exec.Command("./binary")
	err = sqlsmith_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: autoconf-archive")
exec.Command("latte", "install", "autoconf-archive")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libpq")
exec.Command("latte", "install", "libpq")
	fmt.Println("Instalando dependencia: libpqxx")
exec.Command("latte", "install", "libpqxx")
}

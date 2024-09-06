package main

// Sqlite - Command-line interface for SQLite
// Homepage: https://sqlite.org/index.html

import (
	"fmt"
	
	"os/exec"
)

func installSqlite() {
	// Método 1: Descargar y extraer .tar.gz
	sqlite_tar_url := "https://www.sqlite.org/2024/sqlite-autoconf-3460100.tar.gz"
	sqlite_cmd_tar := exec.Command("curl", "-L", sqlite_tar_url, "-o", "package.tar.gz")
	err := sqlite_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sqlite_zip_url := "https://www.sqlite.org/2024/sqlite-autoconf-3460100.zip"
	sqlite_cmd_zip := exec.Command("curl", "-L", sqlite_zip_url, "-o", "package.zip")
	err = sqlite_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sqlite_bin_url := "https://www.sqlite.org/2024/sqlite-autoconf-3460100.bin"
	sqlite_cmd_bin := exec.Command("curl", "-L", sqlite_bin_url, "-o", "binary.bin")
	err = sqlite_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sqlite_src_url := "https://www.sqlite.org/2024/sqlite-autoconf-3460100.src.tar.gz"
	sqlite_cmd_src := exec.Command("curl", "-L", sqlite_src_url, "-o", "source.tar.gz")
	err = sqlite_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sqlite_cmd_direct := exec.Command("./binary")
	err = sqlite_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}

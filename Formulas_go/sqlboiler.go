package main

// Sqlboiler - Generate a Go ORM tailored to your database schema
// Homepage: https://github.com/volatiletech/sqlboiler

import (
	"fmt"
	
	"os/exec"
)

func installSqlboiler() {
	// Método 1: Descargar y extraer .tar.gz
	sqlboiler_tar_url := "https://github.com/volatiletech/sqlboiler/archive/refs/tags/v4.16.2.tar.gz"
	sqlboiler_cmd_tar := exec.Command("curl", "-L", sqlboiler_tar_url, "-o", "package.tar.gz")
	err := sqlboiler_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sqlboiler_zip_url := "https://github.com/volatiletech/sqlboiler/archive/refs/tags/v4.16.2.zip"
	sqlboiler_cmd_zip := exec.Command("curl", "-L", sqlboiler_zip_url, "-o", "package.zip")
	err = sqlboiler_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sqlboiler_bin_url := "https://github.com/volatiletech/sqlboiler/archive/refs/tags/v4.16.2.bin"
	sqlboiler_cmd_bin := exec.Command("curl", "-L", sqlboiler_bin_url, "-o", "binary.bin")
	err = sqlboiler_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sqlboiler_src_url := "https://github.com/volatiletech/sqlboiler/archive/refs/tags/v4.16.2.src.tar.gz"
	sqlboiler_cmd_src := exec.Command("curl", "-L", sqlboiler_src_url, "-o", "source.tar.gz")
	err = sqlboiler_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sqlboiler_cmd_direct := exec.Command("./binary")
	err = sqlboiler_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}

package main

// Sqlcmd - Microsoft SQL Server command-line interface
// Homepage: https://github.com/microsoft/go-sqlcmd

import (
	"fmt"
	
	"os/exec"
)

func installSqlcmd() {
	// Método 1: Descargar y extraer .tar.gz
	sqlcmd_tar_url := "https://github.com/microsoft/go-sqlcmd/archive/refs/tags/v1.8.0.tar.gz"
	sqlcmd_cmd_tar := exec.Command("curl", "-L", sqlcmd_tar_url, "-o", "package.tar.gz")
	err := sqlcmd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sqlcmd_zip_url := "https://github.com/microsoft/go-sqlcmd/archive/refs/tags/v1.8.0.zip"
	sqlcmd_cmd_zip := exec.Command("curl", "-L", sqlcmd_zip_url, "-o", "package.zip")
	err = sqlcmd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sqlcmd_bin_url := "https://github.com/microsoft/go-sqlcmd/archive/refs/tags/v1.8.0.bin"
	sqlcmd_cmd_bin := exec.Command("curl", "-L", sqlcmd_bin_url, "-o", "binary.bin")
	err = sqlcmd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sqlcmd_src_url := "https://github.com/microsoft/go-sqlcmd/archive/refs/tags/v1.8.0.src.tar.gz"
	sqlcmd_cmd_src := exec.Command("curl", "-L", sqlcmd_src_url, "-o", "source.tar.gz")
	err = sqlcmd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sqlcmd_cmd_direct := exec.Command("./binary")
	err = sqlcmd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}

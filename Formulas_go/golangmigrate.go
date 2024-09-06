package main

// GolangMigrate - Database migrations CLI tool
// Homepage: https://github.com/golang-migrate/migrate

import (
	"fmt"
	
	"os/exec"
)

func installGolangMigrate() {
	// Método 1: Descargar y extraer .tar.gz
	golangmigrate_tar_url := "https://github.com/golang-migrate/migrate/archive/refs/tags/v4.17.1.tar.gz"
	golangmigrate_cmd_tar := exec.Command("curl", "-L", golangmigrate_tar_url, "-o", "package.tar.gz")
	err := golangmigrate_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	golangmigrate_zip_url := "https://github.com/golang-migrate/migrate/archive/refs/tags/v4.17.1.zip"
	golangmigrate_cmd_zip := exec.Command("curl", "-L", golangmigrate_zip_url, "-o", "package.zip")
	err = golangmigrate_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	golangmigrate_bin_url := "https://github.com/golang-migrate/migrate/archive/refs/tags/v4.17.1.bin"
	golangmigrate_cmd_bin := exec.Command("curl", "-L", golangmigrate_bin_url, "-o", "binary.bin")
	err = golangmigrate_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	golangmigrate_src_url := "https://github.com/golang-migrate/migrate/archive/refs/tags/v4.17.1.src.tar.gz"
	golangmigrate_cmd_src := exec.Command("curl", "-L", golangmigrate_src_url, "-o", "source.tar.gz")
	err = golangmigrate_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	golangmigrate_cmd_direct := exec.Command("./binary")
	err = golangmigrate_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}

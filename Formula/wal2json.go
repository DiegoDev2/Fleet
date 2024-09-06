package main

// Wal2json - Convert PostgreSQL changesets to JSON format
// Homepage: https://github.com/eulerto/wal2json

import (
	"fmt"
	
	"os/exec"
)

func installWal2json() {
	// Método 1: Descargar y extraer .tar.gz
	wal2json_tar_url := "https://github.com/eulerto/wal2json/archive/refs/tags/wal2json_2_6.tar.gz"
	wal2json_cmd_tar := exec.Command("curl", "-L", wal2json_tar_url, "-o", "package.tar.gz")
	err := wal2json_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wal2json_zip_url := "https://github.com/eulerto/wal2json/archive/refs/tags/wal2json_2_6.zip"
	wal2json_cmd_zip := exec.Command("curl", "-L", wal2json_zip_url, "-o", "package.zip")
	err = wal2json_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wal2json_bin_url := "https://github.com/eulerto/wal2json/archive/refs/tags/wal2json_2_6.bin"
	wal2json_cmd_bin := exec.Command("curl", "-L", wal2json_bin_url, "-o", "binary.bin")
	err = wal2json_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wal2json_src_url := "https://github.com/eulerto/wal2json/archive/refs/tags/wal2json_2_6.src.tar.gz"
	wal2json_cmd_src := exec.Command("curl", "-L", wal2json_src_url, "-o", "source.tar.gz")
	err = wal2json_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wal2json_cmd_direct := exec.Command("./binary")
	err = wal2json_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: postgresql@14")
	exec.Command("latte", "install", "postgresql@14").Run()
}

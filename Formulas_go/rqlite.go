package main

// Rqlite - Lightweight, distributed relational database built on SQLite
// Homepage: https://www.rqlite.io/

import (
	"fmt"
	
	"os/exec"
)

func installRqlite() {
	// Método 1: Descargar y extraer .tar.gz
	rqlite_tar_url := "https://github.com/rqlite/rqlite/archive/refs/tags/v8.30.0.tar.gz"
	rqlite_cmd_tar := exec.Command("curl", "-L", rqlite_tar_url, "-o", "package.tar.gz")
	err := rqlite_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rqlite_zip_url := "https://github.com/rqlite/rqlite/archive/refs/tags/v8.30.0.zip"
	rqlite_cmd_zip := exec.Command("curl", "-L", rqlite_zip_url, "-o", "package.zip")
	err = rqlite_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rqlite_bin_url := "https://github.com/rqlite/rqlite/archive/refs/tags/v8.30.0.bin"
	rqlite_cmd_bin := exec.Command("curl", "-L", rqlite_bin_url, "-o", "binary.bin")
	err = rqlite_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rqlite_src_url := "https://github.com/rqlite/rqlite/archive/refs/tags/v8.30.0.src.tar.gz"
	rqlite_cmd_src := exec.Command("curl", "-L", rqlite_src_url, "-o", "source.tar.gz")
	err = rqlite_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rqlite_cmd_direct := exec.Command("./binary")
	err = rqlite_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}

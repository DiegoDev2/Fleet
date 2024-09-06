package main

// Pgvector - Open-source vector similarity search for Postgres
// Homepage: https://github.com/pgvector/pgvector

import (
	"fmt"
	
	"os/exec"
)

func installPgvector() {
	// Método 1: Descargar y extraer .tar.gz
	pgvector_tar_url := "https://github.com/pgvector/pgvector/archive/refs/tags/v0.7.4.tar.gz"
	pgvector_cmd_tar := exec.Command("curl", "-L", pgvector_tar_url, "-o", "package.tar.gz")
	err := pgvector_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pgvector_zip_url := "https://github.com/pgvector/pgvector/archive/refs/tags/v0.7.4.zip"
	pgvector_cmd_zip := exec.Command("curl", "-L", pgvector_zip_url, "-o", "package.zip")
	err = pgvector_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pgvector_bin_url := "https://github.com/pgvector/pgvector/archive/refs/tags/v0.7.4.bin"
	pgvector_cmd_bin := exec.Command("curl", "-L", pgvector_bin_url, "-o", "binary.bin")
	err = pgvector_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pgvector_src_url := "https://github.com/pgvector/pgvector/archive/refs/tags/v0.7.4.src.tar.gz"
	pgvector_cmd_src := exec.Command("curl", "-L", pgvector_src_url, "-o", "source.tar.gz")
	err = pgvector_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pgvector_cmd_direct := exec.Command("./binary")
	err = pgvector_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: postgresql@14")
exec.Command("latte", "install", "postgresql@14")
}

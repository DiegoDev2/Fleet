package main

// Sqlmap - Penetration testing for SQL injection and database servers
// Homepage: https://sqlmap.org

import (
	"fmt"
	
	"os/exec"
)

func installSqlmap() {
	// Método 1: Descargar y extraer .tar.gz
	sqlmap_tar_url := "https://github.com/sqlmapproject/sqlmap/archive/refs/tags/1.8.8.tar.gz"
	sqlmap_cmd_tar := exec.Command("curl", "-L", sqlmap_tar_url, "-o", "package.tar.gz")
	err := sqlmap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sqlmap_zip_url := "https://github.com/sqlmapproject/sqlmap/archive/refs/tags/1.8.8.zip"
	sqlmap_cmd_zip := exec.Command("curl", "-L", sqlmap_zip_url, "-o", "package.zip")
	err = sqlmap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sqlmap_bin_url := "https://github.com/sqlmapproject/sqlmap/archive/refs/tags/1.8.8.bin"
	sqlmap_cmd_bin := exec.Command("curl", "-L", sqlmap_bin_url, "-o", "binary.bin")
	err = sqlmap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sqlmap_src_url := "https://github.com/sqlmapproject/sqlmap/archive/refs/tags/1.8.8.src.tar.gz"
	sqlmap_cmd_src := exec.Command("curl", "-L", sqlmap_src_url, "-o", "source.tar.gz")
	err = sqlmap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sqlmap_cmd_direct := exec.Command("./binary")
	err = sqlmap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}

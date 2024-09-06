package main

// Sqlancer - Detecting Logic Bugs in DBMS
// Homepage: https://github.com/sqlancer/sqlancer

import (
	"fmt"
	
	"os/exec"
)

func installSqlancer() {
	// Método 1: Descargar y extraer .tar.gz
	sqlancer_tar_url := "https://github.com/sqlancer/sqlancer/archive/refs/tags/v2.0.0.tar.gz"
	sqlancer_cmd_tar := exec.Command("curl", "-L", sqlancer_tar_url, "-o", "package.tar.gz")
	err := sqlancer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sqlancer_zip_url := "https://github.com/sqlancer/sqlancer/archive/refs/tags/v2.0.0.zip"
	sqlancer_cmd_zip := exec.Command("curl", "-L", sqlancer_zip_url, "-o", "package.zip")
	err = sqlancer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sqlancer_bin_url := "https://github.com/sqlancer/sqlancer/archive/refs/tags/v2.0.0.bin"
	sqlancer_cmd_bin := exec.Command("curl", "-L", sqlancer_bin_url, "-o", "binary.bin")
	err = sqlancer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sqlancer_src_url := "https://github.com/sqlancer/sqlancer/archive/refs/tags/v2.0.0.src.tar.gz"
	sqlancer_cmd_src := exec.Command("curl", "-L", sqlancer_src_url, "-o", "source.tar.gz")
	err = sqlancer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sqlancer_cmd_direct := exec.Command("./binary")
	err = sqlancer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: maven")
	exec.Command("latte", "install", "maven").Run()
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}

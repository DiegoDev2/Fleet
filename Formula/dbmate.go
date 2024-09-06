package main

// Dbmate - Lightweight, framework-agnostic database migration tool
// Homepage: https://github.com/amacneil/dbmate

import (
	"fmt"
	
	"os/exec"
)

func installDbmate() {
	// Método 1: Descargar y extraer .tar.gz
	dbmate_tar_url := "https://github.com/amacneil/dbmate/archive/refs/tags/v2.20.0.tar.gz"
	dbmate_cmd_tar := exec.Command("curl", "-L", dbmate_tar_url, "-o", "package.tar.gz")
	err := dbmate_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dbmate_zip_url := "https://github.com/amacneil/dbmate/archive/refs/tags/v2.20.0.zip"
	dbmate_cmd_zip := exec.Command("curl", "-L", dbmate_zip_url, "-o", "package.zip")
	err = dbmate_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dbmate_bin_url := "https://github.com/amacneil/dbmate/archive/refs/tags/v2.20.0.bin"
	dbmate_cmd_bin := exec.Command("curl", "-L", dbmate_bin_url, "-o", "binary.bin")
	err = dbmate_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dbmate_src_url := "https://github.com/amacneil/dbmate/archive/refs/tags/v2.20.0.src.tar.gz"
	dbmate_cmd_src := exec.Command("curl", "-L", dbmate_src_url, "-o", "source.tar.gz")
	err = dbmate_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dbmate_cmd_direct := exec.Command("./binary")
	err = dbmate_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}

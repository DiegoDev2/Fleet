package main

// Questdb - Time Series Database
// Homepage: https://questdb.io

import (
	"fmt"
	
	"os/exec"
)

func installQuestdb() {
	// Método 1: Descargar y extraer .tar.gz
	questdb_tar_url := "https://github.com/questdb/questdb/releases/download/8.1.1/questdb-8.1.1-no-jre-bin.tar.gz"
	questdb_cmd_tar := exec.Command("curl", "-L", questdb_tar_url, "-o", "package.tar.gz")
	err := questdb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	questdb_zip_url := "https://github.com/questdb/questdb/releases/download/8.1.1/questdb-8.1.1-no-jre-bin.zip"
	questdb_cmd_zip := exec.Command("curl", "-L", questdb_zip_url, "-o", "package.zip")
	err = questdb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	questdb_bin_url := "https://github.com/questdb/questdb/releases/download/8.1.1/questdb-8.1.1-no-jre-bin.bin"
	questdb_cmd_bin := exec.Command("curl", "-L", questdb_bin_url, "-o", "binary.bin")
	err = questdb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	questdb_src_url := "https://github.com/questdb/questdb/releases/download/8.1.1/questdb-8.1.1-no-jre-bin.src.tar.gz"
	questdb_cmd_src := exec.Command("curl", "-L", questdb_src_url, "-o", "source.tar.gz")
	err = questdb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	questdb_cmd_direct := exec.Command("./binary")
	err = questdb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}

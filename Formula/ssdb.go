package main

// Ssdb - NoSQL database supporting many data structures: Redis alternative
// Homepage: https://ssdb.io/

import (
	"fmt"
	
	"os/exec"
)

func installSsdb() {
	// Método 1: Descargar y extraer .tar.gz
	ssdb_tar_url := "https://github.com/ideawu/ssdb/archive/refs/tags/1.9.9.tar.gz"
	ssdb_cmd_tar := exec.Command("curl", "-L", ssdb_tar_url, "-o", "package.tar.gz")
	err := ssdb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ssdb_zip_url := "https://github.com/ideawu/ssdb/archive/refs/tags/1.9.9.zip"
	ssdb_cmd_zip := exec.Command("curl", "-L", ssdb_zip_url, "-o", "package.zip")
	err = ssdb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ssdb_bin_url := "https://github.com/ideawu/ssdb/archive/refs/tags/1.9.9.bin"
	ssdb_cmd_bin := exec.Command("curl", "-L", ssdb_bin_url, "-o", "binary.bin")
	err = ssdb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ssdb_src_url := "https://github.com/ideawu/ssdb/archive/refs/tags/1.9.9.src.tar.gz"
	ssdb_cmd_src := exec.Command("curl", "-L", ssdb_src_url, "-o", "source.tar.gz")
	err = ssdb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ssdb_cmd_direct := exec.Command("./binary")
	err = ssdb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
}

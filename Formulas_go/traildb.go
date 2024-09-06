package main

// Traildb - Blazingly-fast database for log-structured data
// Homepage: https://traildb.io/

import (
	"fmt"
	
	"os/exec"
)

func installTraildb() {
	// Método 1: Descargar y extraer .tar.gz
	traildb_tar_url := "https://github.com/traildb/traildb/archive/refs/tags/0.6.tar.gz"
	traildb_cmd_tar := exec.Command("curl", "-L", traildb_tar_url, "-o", "package.tar.gz")
	err := traildb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	traildb_zip_url := "https://github.com/traildb/traildb/archive/refs/tags/0.6.zip"
	traildb_cmd_zip := exec.Command("curl", "-L", traildb_zip_url, "-o", "package.zip")
	err = traildb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	traildb_bin_url := "https://github.com/traildb/traildb/archive/refs/tags/0.6.bin"
	traildb_cmd_bin := exec.Command("curl", "-L", traildb_bin_url, "-o", "binary.bin")
	err = traildb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	traildb_src_url := "https://github.com/traildb/traildb/archive/refs/tags/0.6.src.tar.gz"
	traildb_cmd_src := exec.Command("curl", "-L", traildb_src_url, "-o", "source.tar.gz")
	err = traildb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	traildb_cmd_direct := exec.Command("./binary")
	err = traildb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: judy")
exec.Command("latte", "install", "judy")
	fmt.Println("Instalando dependencia: libarchive")
exec.Command("latte", "install", "libarchive")
}

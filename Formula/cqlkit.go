package main

// Cqlkit - CLI tool to export Cassandra query as CSV and JSON format
// Homepage: https://github.com/tenmax/cqlkit

import (
	"fmt"
	
	"os/exec"
)

func installCqlkit() {
	// Método 1: Descargar y extraer .tar.gz
	cqlkit_tar_url := "https://github.com/tenmax/cqlkit/releases/download/v0.3.3/cqlkit-0.3.3.zip"
	cqlkit_cmd_tar := exec.Command("curl", "-L", cqlkit_tar_url, "-o", "package.tar.gz")
	err := cqlkit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cqlkit_zip_url := "https://github.com/tenmax/cqlkit/releases/download/v0.3.3/cqlkit-0.3.3.zip"
	cqlkit_cmd_zip := exec.Command("curl", "-L", cqlkit_zip_url, "-o", "package.zip")
	err = cqlkit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cqlkit_bin_url := "https://github.com/tenmax/cqlkit/releases/download/v0.3.3/cqlkit-0.3.3.zip"
	cqlkit_cmd_bin := exec.Command("curl", "-L", cqlkit_bin_url, "-o", "binary.bin")
	err = cqlkit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cqlkit_src_url := "https://github.com/tenmax/cqlkit/releases/download/v0.3.3/cqlkit-0.3.3.zip"
	cqlkit_cmd_src := exec.Command("curl", "-L", cqlkit_src_url, "-o", "source.tar.gz")
	err = cqlkit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cqlkit_cmd_direct := exec.Command("./binary")
	err = cqlkit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}

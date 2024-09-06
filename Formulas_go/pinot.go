package main

// Pinot - Realtime distributed OLAP datastore
// Homepage: https://pinot.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installPinot() {
	// Método 1: Descargar y extraer .tar.gz
	pinot_tar_url := "https://downloads.apache.org/pinot/apache-pinot-1.2.0/apache-pinot-1.2.0-bin.tar.gz"
	pinot_cmd_tar := exec.Command("curl", "-L", pinot_tar_url, "-o", "package.tar.gz")
	err := pinot_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pinot_zip_url := "https://downloads.apache.org/pinot/apache-pinot-1.2.0/apache-pinot-1.2.0-bin.zip"
	pinot_cmd_zip := exec.Command("curl", "-L", pinot_zip_url, "-o", "package.zip")
	err = pinot_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pinot_bin_url := "https://downloads.apache.org/pinot/apache-pinot-1.2.0/apache-pinot-1.2.0-bin.bin"
	pinot_cmd_bin := exec.Command("curl", "-L", pinot_bin_url, "-o", "binary.bin")
	err = pinot_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pinot_src_url := "https://downloads.apache.org/pinot/apache-pinot-1.2.0/apache-pinot-1.2.0-bin.src.tar.gz"
	pinot_cmd_src := exec.Command("curl", "-L", pinot_src_url, "-o", "source.tar.gz")
	err = pinot_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pinot_cmd_direct := exec.Command("./binary")
	err = pinot_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@11")
exec.Command("latte", "install", "openjdk@11")
}

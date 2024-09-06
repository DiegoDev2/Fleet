package main

// Maxwell - Reads MySQL binlogs and writes row updates as JSON to Kafka
// Homepage: https://maxwells-daemon.io/

import (
	"fmt"
	
	"os/exec"
)

func installMaxwell() {
	// Método 1: Descargar y extraer .tar.gz
	maxwell_tar_url := "https://github.com/zendesk/maxwell/releases/download/v1.41.2/maxwell-1.41.2.tar.gz"
	maxwell_cmd_tar := exec.Command("curl", "-L", maxwell_tar_url, "-o", "package.tar.gz")
	err := maxwell_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	maxwell_zip_url := "https://github.com/zendesk/maxwell/releases/download/v1.41.2/maxwell-1.41.2.zip"
	maxwell_cmd_zip := exec.Command("curl", "-L", maxwell_zip_url, "-o", "package.zip")
	err = maxwell_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	maxwell_bin_url := "https://github.com/zendesk/maxwell/releases/download/v1.41.2/maxwell-1.41.2.bin"
	maxwell_cmd_bin := exec.Command("curl", "-L", maxwell_bin_url, "-o", "binary.bin")
	err = maxwell_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	maxwell_src_url := "https://github.com/zendesk/maxwell/releases/download/v1.41.2/maxwell-1.41.2.src.tar.gz"
	maxwell_cmd_src := exec.Command("curl", "-L", maxwell_src_url, "-o", "source.tar.gz")
	err = maxwell_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	maxwell_cmd_direct := exec.Command("./binary")
	err = maxwell_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@11")
exec.Command("latte", "install", "openjdk@11")
}

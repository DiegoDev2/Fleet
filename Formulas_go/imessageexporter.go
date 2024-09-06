package main

// ImessageExporter - Command-line tool to export and inspect local iMessage database
// Homepage: https://github.com/ReagentX/imessage-exporter

import (
	"fmt"
	
	"os/exec"
)

func installImessageExporter() {
	// Método 1: Descargar y extraer .tar.gz
	imessageexporter_tar_url := "https://github.com/ReagentX/imessage-exporter/archive/refs/tags/2.0.1.tar.gz"
	imessageexporter_cmd_tar := exec.Command("curl", "-L", imessageexporter_tar_url, "-o", "package.tar.gz")
	err := imessageexporter_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	imessageexporter_zip_url := "https://github.com/ReagentX/imessage-exporter/archive/refs/tags/2.0.1.zip"
	imessageexporter_cmd_zip := exec.Command("curl", "-L", imessageexporter_zip_url, "-o", "package.zip")
	err = imessageexporter_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	imessageexporter_bin_url := "https://github.com/ReagentX/imessage-exporter/archive/refs/tags/2.0.1.bin"
	imessageexporter_cmd_bin := exec.Command("curl", "-L", imessageexporter_bin_url, "-o", "binary.bin")
	err = imessageexporter_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	imessageexporter_src_url := "https://github.com/ReagentX/imessage-exporter/archive/refs/tags/2.0.1.src.tar.gz"
	imessageexporter_cmd_src := exec.Command("curl", "-L", imessageexporter_src_url, "-o", "source.tar.gz")
	err = imessageexporter_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	imessageexporter_cmd_direct := exec.Command("./binary")
	err = imessageexporter_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}

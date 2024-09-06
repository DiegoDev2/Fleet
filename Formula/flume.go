package main

// Flume - Hadoop-based distributed log collection and aggregation
// Homepage: https://flume.apache.org

import (
	"fmt"
	
	"os/exec"
)

func installFlume() {
	// Método 1: Descargar y extraer .tar.gz
	flume_tar_url := "https://www.apache.org/dyn/closer.lua?path=flume/1.11.0/apache-flume-1.11.0-bin.tar.gz"
	flume_cmd_tar := exec.Command("curl", "-L", flume_tar_url, "-o", "package.tar.gz")
	err := flume_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	flume_zip_url := "https://www.apache.org/dyn/closer.lua?path=flume/1.11.0/apache-flume-1.11.0-bin.zip"
	flume_cmd_zip := exec.Command("curl", "-L", flume_zip_url, "-o", "package.zip")
	err = flume_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	flume_bin_url := "https://www.apache.org/dyn/closer.lua?path=flume/1.11.0/apache-flume-1.11.0-bin.bin"
	flume_cmd_bin := exec.Command("curl", "-L", flume_bin_url, "-o", "binary.bin")
	err = flume_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	flume_src_url := "https://www.apache.org/dyn/closer.lua?path=flume/1.11.0/apache-flume-1.11.0-bin.src.tar.gz"
	flume_cmd_src := exec.Command("curl", "-L", flume_src_url, "-o", "source.tar.gz")
	err = flume_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	flume_cmd_direct := exec.Command("./binary")
	err = flume_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: hadoop")
	exec.Command("latte", "install", "hadoop").Run()
	fmt.Println("Instalando dependencia: openjdk@11")
	exec.Command("latte", "install", "openjdk@11").Run()
}

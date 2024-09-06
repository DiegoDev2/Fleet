package main

// Storm - Distributed realtime computation system to process data streams
// Homepage: https://storm.apache.org

import (
	"fmt"
	
	"os/exec"
)

func installStorm() {
	// Método 1: Descargar y extraer .tar.gz
	storm_tar_url := "https://dlcdn.apache.org/storm/apache-storm-2.6.4/apache-storm-2.6.4.tar.gz"
	storm_cmd_tar := exec.Command("curl", "-L", storm_tar_url, "-o", "package.tar.gz")
	err := storm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	storm_zip_url := "https://dlcdn.apache.org/storm/apache-storm-2.6.4/apache-storm-2.6.4.zip"
	storm_cmd_zip := exec.Command("curl", "-L", storm_zip_url, "-o", "package.zip")
	err = storm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	storm_bin_url := "https://dlcdn.apache.org/storm/apache-storm-2.6.4/apache-storm-2.6.4.bin"
	storm_cmd_bin := exec.Command("curl", "-L", storm_bin_url, "-o", "binary.bin")
	err = storm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	storm_src_url := "https://dlcdn.apache.org/storm/apache-storm-2.6.4/apache-storm-2.6.4.src.tar.gz"
	storm_cmd_src := exec.Command("curl", "-L", storm_src_url, "-o", "source.tar.gz")
	err = storm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	storm_cmd_direct := exec.Command("./binary")
	err = storm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}

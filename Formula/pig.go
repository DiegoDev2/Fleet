package main

// Pig - Platform for analyzing large data sets
// Homepage: https://pig.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installPig() {
	// Método 1: Descargar y extraer .tar.gz
	pig_tar_url := "https://www.apache.org/dyn/closer.lua?path=pig/pig-0.17.0/pig-0.17.0.tar.gz"
	pig_cmd_tar := exec.Command("curl", "-L", pig_tar_url, "-o", "package.tar.gz")
	err := pig_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pig_zip_url := "https://www.apache.org/dyn/closer.lua?path=pig/pig-0.17.0/pig-0.17.0.zip"
	pig_cmd_zip := exec.Command("curl", "-L", pig_zip_url, "-o", "package.zip")
	err = pig_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pig_bin_url := "https://www.apache.org/dyn/closer.lua?path=pig/pig-0.17.0/pig-0.17.0.bin"
	pig_cmd_bin := exec.Command("curl", "-L", pig_bin_url, "-o", "binary.bin")
	err = pig_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pig_src_url := "https://www.apache.org/dyn/closer.lua?path=pig/pig-0.17.0/pig-0.17.0.src.tar.gz"
	pig_cmd_src := exec.Command("curl", "-L", pig_src_url, "-o", "source.tar.gz")
	err = pig_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pig_cmd_direct := exec.Command("./binary")
	err = pig_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@17")
	exec.Command("latte", "install", "openjdk@17").Run()
}

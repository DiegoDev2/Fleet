package main

// OpenZwave - Library that interfaces with selected Z-Wave PC controllers
// Homepage: https://github.com/OpenZWave/open-zwave

import (
	"fmt"
	
	"os/exec"
)

func installOpenZwave() {
	// Método 1: Descargar y extraer .tar.gz
	openzwave_tar_url := "http://old.openzwave.com/downloads/openzwave-1.6.1914.tar.gz"
	openzwave_cmd_tar := exec.Command("curl", "-L", openzwave_tar_url, "-o", "package.tar.gz")
	err := openzwave_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openzwave_zip_url := "http://old.openzwave.com/downloads/openzwave-1.6.1914.zip"
	openzwave_cmd_zip := exec.Command("curl", "-L", openzwave_zip_url, "-o", "package.zip")
	err = openzwave_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openzwave_bin_url := "http://old.openzwave.com/downloads/openzwave-1.6.1914.bin"
	openzwave_cmd_bin := exec.Command("curl", "-L", openzwave_bin_url, "-o", "binary.bin")
	err = openzwave_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openzwave_src_url := "http://old.openzwave.com/downloads/openzwave-1.6.1914.src.tar.gz"
	openzwave_cmd_src := exec.Command("curl", "-L", openzwave_src_url, "-o", "source.tar.gz")
	err = openzwave_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openzwave_cmd_direct := exec.Command("./binary")
	err = openzwave_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: doxygen")
	exec.Command("latte", "install", "doxygen").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}

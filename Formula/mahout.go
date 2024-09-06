package main

// Mahout - Library to help build scalable machine learning libraries
// Homepage: https://mahout.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installMahout() {
	// Método 1: Descargar y extraer .tar.gz
	mahout_tar_url := "https://www.apache.org/dyn/closer.lua?path=mahout/0.13.0/apache-mahout-distribution-0.13.0.tar.gz"
	mahout_cmd_tar := exec.Command("curl", "-L", mahout_tar_url, "-o", "package.tar.gz")
	err := mahout_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mahout_zip_url := "https://www.apache.org/dyn/closer.lua?path=mahout/0.13.0/apache-mahout-distribution-0.13.0.zip"
	mahout_cmd_zip := exec.Command("curl", "-L", mahout_zip_url, "-o", "package.zip")
	err = mahout_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mahout_bin_url := "https://www.apache.org/dyn/closer.lua?path=mahout/0.13.0/apache-mahout-distribution-0.13.0.bin"
	mahout_cmd_bin := exec.Command("curl", "-L", mahout_bin_url, "-o", "binary.bin")
	err = mahout_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mahout_src_url := "https://www.apache.org/dyn/closer.lua?path=mahout/0.13.0/apache-mahout-distribution-0.13.0.src.tar.gz"
	mahout_cmd_src := exec.Command("curl", "-L", mahout_src_url, "-o", "source.tar.gz")
	err = mahout_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mahout_cmd_direct := exec.Command("./binary")
	err = mahout_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: maven")
	exec.Command("latte", "install", "maven").Run()
	fmt.Println("Instalando dependencia: hadoop")
	exec.Command("latte", "install", "hadoop").Run()
	fmt.Println("Instalando dependencia: openjdk@11")
	exec.Command("latte", "install", "openjdk@11").Run()
}

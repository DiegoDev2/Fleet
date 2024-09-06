package main

// TomeePlus - Everything in TomEE Web Profile and JAX-RS, plus more
// Homepage: https://tomee.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installTomeePlus() {
	// Método 1: Descargar y extraer .tar.gz
	tomeeplus_tar_url := "https://www.apache.org/dyn/closer.lua?path=tomee/tomee-9.1.3/apache-tomee-9.1.3-plus.tar.gz"
	tomeeplus_cmd_tar := exec.Command("curl", "-L", tomeeplus_tar_url, "-o", "package.tar.gz")
	err := tomeeplus_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tomeeplus_zip_url := "https://www.apache.org/dyn/closer.lua?path=tomee/tomee-9.1.3/apache-tomee-9.1.3-plus.zip"
	tomeeplus_cmd_zip := exec.Command("curl", "-L", tomeeplus_zip_url, "-o", "package.zip")
	err = tomeeplus_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tomeeplus_bin_url := "https://www.apache.org/dyn/closer.lua?path=tomee/tomee-9.1.3/apache-tomee-9.1.3-plus.bin"
	tomeeplus_cmd_bin := exec.Command("curl", "-L", tomeeplus_bin_url, "-o", "binary.bin")
	err = tomeeplus_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tomeeplus_src_url := "https://www.apache.org/dyn/closer.lua?path=tomee/tomee-9.1.3/apache-tomee-9.1.3-plus.src.tar.gz"
	tomeeplus_cmd_src := exec.Command("curl", "-L", tomeeplus_src_url, "-o", "source.tar.gz")
	err = tomeeplus_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tomeeplus_cmd_direct := exec.Command("./binary")
	err = tomeeplus_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}

package main

// TomeeWebprofile - All-Apache Java EE 7 Web Profile stack
// Homepage: https://tomee.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installTomeeWebprofile() {
	// Método 1: Descargar y extraer .tar.gz
	tomeewebprofile_tar_url := "https://www.apache.org/dyn/closer.lua?path=tomee/tomee-9.1.3/apache-tomee-9.1.3-webprofile.tar.gz"
	tomeewebprofile_cmd_tar := exec.Command("curl", "-L", tomeewebprofile_tar_url, "-o", "package.tar.gz")
	err := tomeewebprofile_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tomeewebprofile_zip_url := "https://www.apache.org/dyn/closer.lua?path=tomee/tomee-9.1.3/apache-tomee-9.1.3-webprofile.zip"
	tomeewebprofile_cmd_zip := exec.Command("curl", "-L", tomeewebprofile_zip_url, "-o", "package.zip")
	err = tomeewebprofile_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tomeewebprofile_bin_url := "https://www.apache.org/dyn/closer.lua?path=tomee/tomee-9.1.3/apache-tomee-9.1.3-webprofile.bin"
	tomeewebprofile_cmd_bin := exec.Command("curl", "-L", tomeewebprofile_bin_url, "-o", "binary.bin")
	err = tomeewebprofile_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tomeewebprofile_src_url := "https://www.apache.org/dyn/closer.lua?path=tomee/tomee-9.1.3/apache-tomee-9.1.3-webprofile.src.tar.gz"
	tomeewebprofile_cmd_src := exec.Command("curl", "-L", tomeewebprofile_src_url, "-o", "source.tar.gz")
	err = tomeewebprofile_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tomeewebprofile_cmd_direct := exec.Command("./binary")
	err = tomeewebprofile_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}

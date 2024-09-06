package main

// LicenseEye - Tool to check and fix license headers and resolve dependency licenses
// Homepage: https://github.com/apache/skywalking-eyes

import (
	"fmt"
	
	"os/exec"
)

func installLicenseEye() {
	// Método 1: Descargar y extraer .tar.gz
	licenseeye_tar_url := "https://www.apache.org/dyn/closer.lua?path=skywalking/eyes/0.6.0/skywalking-license-eye-0.6.0-src.tgz"
	licenseeye_cmd_tar := exec.Command("curl", "-L", licenseeye_tar_url, "-o", "package.tar.gz")
	err := licenseeye_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	licenseeye_zip_url := "https://www.apache.org/dyn/closer.lua?path=skywalking/eyes/0.6.0/skywalking-license-eye-0.6.0-src.tgz"
	licenseeye_cmd_zip := exec.Command("curl", "-L", licenseeye_zip_url, "-o", "package.zip")
	err = licenseeye_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	licenseeye_bin_url := "https://www.apache.org/dyn/closer.lua?path=skywalking/eyes/0.6.0/skywalking-license-eye-0.6.0-src.tgz"
	licenseeye_cmd_bin := exec.Command("curl", "-L", licenseeye_bin_url, "-o", "binary.bin")
	err = licenseeye_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	licenseeye_src_url := "https://www.apache.org/dyn/closer.lua?path=skywalking/eyes/0.6.0/skywalking-license-eye-0.6.0-src.tgz"
	licenseeye_cmd_src := exec.Command("curl", "-L", licenseeye_src_url, "-o", "source.tar.gz")
	err = licenseeye_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	licenseeye_cmd_direct := exec.Command("./binary")
	err = licenseeye_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}

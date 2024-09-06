package main

// Checkstyle - Check Java source against a coding standard
// Homepage: https://checkstyle.sourceforge.io/

import (
	"fmt"
	
	"os/exec"
)

func installCheckstyle() {
	// Método 1: Descargar y extraer .tar.gz
	checkstyle_tar_url := "https://github.com/checkstyle/checkstyle/releases/download/checkstyle-10.18.1/checkstyle-10.18.1-all.jar"
	checkstyle_cmd_tar := exec.Command("curl", "-L", checkstyle_tar_url, "-o", "package.tar.gz")
	err := checkstyle_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	checkstyle_zip_url := "https://github.com/checkstyle/checkstyle/releases/download/checkstyle-10.18.1/checkstyle-10.18.1-all.jar"
	checkstyle_cmd_zip := exec.Command("curl", "-L", checkstyle_zip_url, "-o", "package.zip")
	err = checkstyle_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	checkstyle_bin_url := "https://github.com/checkstyle/checkstyle/releases/download/checkstyle-10.18.1/checkstyle-10.18.1-all.jar"
	checkstyle_cmd_bin := exec.Command("curl", "-L", checkstyle_bin_url, "-o", "binary.bin")
	err = checkstyle_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	checkstyle_src_url := "https://github.com/checkstyle/checkstyle/releases/download/checkstyle-10.18.1/checkstyle-10.18.1-all.jar"
	checkstyle_cmd_src := exec.Command("curl", "-L", checkstyle_src_url, "-o", "source.tar.gz")
	err = checkstyle_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	checkstyle_cmd_direct := exec.Command("./binary")
	err = checkstyle_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}

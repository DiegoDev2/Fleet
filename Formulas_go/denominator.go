package main

// Denominator - Portable Java library for manipulating DNS clouds
// Homepage: https://github.com/Netflix/denominator/tree/v4.7.1/cli

import (
	"fmt"
	
	"os/exec"
)

func installDenominator() {
	// Método 1: Descargar y extraer .tar.gz
	denominator_tar_url := "https://search.maven.org/remotecontent?filepath=com/netflix/denominator/denominator-cli/4.7.1/denominator-cli-4.7.1-fat.jar"
	denominator_cmd_tar := exec.Command("curl", "-L", denominator_tar_url, "-o", "package.tar.gz")
	err := denominator_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	denominator_zip_url := "https://search.maven.org/remotecontent?filepath=com/netflix/denominator/denominator-cli/4.7.1/denominator-cli-4.7.1-fat.jar"
	denominator_cmd_zip := exec.Command("curl", "-L", denominator_zip_url, "-o", "package.zip")
	err = denominator_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	denominator_bin_url := "https://search.maven.org/remotecontent?filepath=com/netflix/denominator/denominator-cli/4.7.1/denominator-cli-4.7.1-fat.jar"
	denominator_cmd_bin := exec.Command("curl", "-L", denominator_bin_url, "-o", "binary.bin")
	err = denominator_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	denominator_src_url := "https://search.maven.org/remotecontent?filepath=com/netflix/denominator/denominator-cli/4.7.1/denominator-cli-4.7.1-fat.jar"
	denominator_cmd_src := exec.Command("curl", "-L", denominator_src_url, "-o", "source.tar.gz")
	err = denominator_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	denominator_cmd_direct := exec.Command("./binary")
	err = denominator_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}

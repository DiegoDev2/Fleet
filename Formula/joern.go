package main

// Joern - Open-source code analysis platform based on code property graphs
// Homepage: https://joern.io/

import (
	"fmt"
	
	"os/exec"
)

func installJoern() {
	// Método 1: Descargar y extraer .tar.gz
	joern_tar_url := "https://github.com/joernio/joern/archive/refs/tags/v2.0.260.tar.gz"
	joern_cmd_tar := exec.Command("curl", "-L", joern_tar_url, "-o", "package.tar.gz")
	err := joern_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	joern_zip_url := "https://github.com/joernio/joern/archive/refs/tags/v2.0.260.zip"
	joern_cmd_zip := exec.Command("curl", "-L", joern_zip_url, "-o", "package.zip")
	err = joern_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	joern_bin_url := "https://github.com/joernio/joern/archive/refs/tags/v2.0.260.bin"
	joern_cmd_bin := exec.Command("curl", "-L", joern_bin_url, "-o", "binary.bin")
	err = joern_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	joern_src_url := "https://github.com/joernio/joern/archive/refs/tags/v2.0.260.src.tar.gz"
	joern_cmd_src := exec.Command("curl", "-L", joern_src_url, "-o", "source.tar.gz")
	err = joern_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	joern_cmd_direct := exec.Command("./binary")
	err = joern_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: sbt")
	exec.Command("latte", "install", "sbt").Run()
	fmt.Println("Instalando dependencia: astgen")
	exec.Command("latte", "install", "astgen").Run()
	fmt.Println("Instalando dependencia: coreutils")
	exec.Command("latte", "install", "coreutils").Run()
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
	fmt.Println("Instalando dependencia: php")
	exec.Command("latte", "install", "php").Run()
}

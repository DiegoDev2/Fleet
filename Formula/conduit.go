package main

// Conduit - Streams data between data stores. Kafka Connect replacement. No JVM required
// Homepage: https://conduit.io/

import (
	"fmt"
	
	"os/exec"
)

func installConduit() {
	// Método 1: Descargar y extraer .tar.gz
	conduit_tar_url := "https://github.com/ConduitIO/conduit/archive/refs/tags/v0.11.1.tar.gz"
	conduit_cmd_tar := exec.Command("curl", "-L", conduit_tar_url, "-o", "package.tar.gz")
	err := conduit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	conduit_zip_url := "https://github.com/ConduitIO/conduit/archive/refs/tags/v0.11.1.zip"
	conduit_cmd_zip := exec.Command("curl", "-L", conduit_zip_url, "-o", "package.zip")
	err = conduit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	conduit_bin_url := "https://github.com/ConduitIO/conduit/archive/refs/tags/v0.11.1.bin"
	conduit_cmd_bin := exec.Command("curl", "-L", conduit_bin_url, "-o", "binary.bin")
	err = conduit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	conduit_src_url := "https://github.com/ConduitIO/conduit/archive/refs/tags/v0.11.1.src.tar.gz"
	conduit_cmd_src := exec.Command("curl", "-L", conduit_src_url, "-o", "source.tar.gz")
	err = conduit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	conduit_cmd_direct := exec.Command("./binary")
	err = conduit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
	fmt.Println("Instalando dependencia: yarn")
	exec.Command("latte", "install", "yarn").Run()
}

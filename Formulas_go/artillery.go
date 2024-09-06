package main

// Artillery - Cloud-native performance & reliability testing for developers and SREs
// Homepage: https://artillery.io/

import (
	"fmt"
	
	"os/exec"
)

func installArtillery() {
	// Método 1: Descargar y extraer .tar.gz
	artillery_tar_url := "https://registry.npmjs.org/artillery/-/artillery-2.0.20.tgz"
	artillery_cmd_tar := exec.Command("curl", "-L", artillery_tar_url, "-o", "package.tar.gz")
	err := artillery_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	artillery_zip_url := "https://registry.npmjs.org/artillery/-/artillery-2.0.20.tgz"
	artillery_cmd_zip := exec.Command("curl", "-L", artillery_zip_url, "-o", "package.zip")
	err = artillery_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	artillery_bin_url := "https://registry.npmjs.org/artillery/-/artillery-2.0.20.tgz"
	artillery_cmd_bin := exec.Command("curl", "-L", artillery_bin_url, "-o", "binary.bin")
	err = artillery_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	artillery_src_url := "https://registry.npmjs.org/artillery/-/artillery-2.0.20.tgz"
	artillery_cmd_src := exec.Command("curl", "-L", artillery_src_url, "-o", "source.tar.gz")
	err = artillery_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	artillery_cmd_direct := exec.Command("./binary")
	err = artillery_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
	fmt.Println("Instalando dependencia: macos-term-size")
exec.Command("latte", "install", "macos-term-size")
}

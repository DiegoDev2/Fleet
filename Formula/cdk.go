package main

// Cdk - Curses development kit provides predefined curses widget for apps
// Homepage: https://invisible-island.net/cdk/

import (
	"fmt"
	
	"os/exec"
)

func installCdk() {
	// Método 1: Descargar y extraer .tar.gz
	cdk_tar_url := "https://invisible-mirror.net/archives/cdk/cdk-5.0-20240619.tgz"
	cdk_cmd_tar := exec.Command("curl", "-L", cdk_tar_url, "-o", "package.tar.gz")
	err := cdk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cdk_zip_url := "https://invisible-mirror.net/archives/cdk/cdk-5.0-20240619.tgz"
	cdk_cmd_zip := exec.Command("curl", "-L", cdk_zip_url, "-o", "package.zip")
	err = cdk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cdk_bin_url := "https://invisible-mirror.net/archives/cdk/cdk-5.0-20240619.tgz"
	cdk_cmd_bin := exec.Command("curl", "-L", cdk_bin_url, "-o", "binary.bin")
	err = cdk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cdk_src_url := "https://invisible-mirror.net/archives/cdk/cdk-5.0-20240619.tgz"
	cdk_cmd_src := exec.Command("curl", "-L", cdk_src_url, "-o", "source.tar.gz")
	err = cdk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cdk_cmd_direct := exec.Command("./binary")
	err = cdk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

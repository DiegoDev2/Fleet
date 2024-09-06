package main

// Bcpp - C(++) beautifier
// Homepage: https://invisible-island.net/bcpp/

import (
	"fmt"
	
	"os/exec"
)

func installBcpp() {
	// Método 1: Descargar y extraer .tar.gz
	bcpp_tar_url := "https://invisible-mirror.net/archives/bcpp/bcpp-20240111.tgz"
	bcpp_cmd_tar := exec.Command("curl", "-L", bcpp_tar_url, "-o", "package.tar.gz")
	err := bcpp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bcpp_zip_url := "https://invisible-mirror.net/archives/bcpp/bcpp-20240111.tgz"
	bcpp_cmd_zip := exec.Command("curl", "-L", bcpp_zip_url, "-o", "package.zip")
	err = bcpp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bcpp_bin_url := "https://invisible-mirror.net/archives/bcpp/bcpp-20240111.tgz"
	bcpp_cmd_bin := exec.Command("curl", "-L", bcpp_bin_url, "-o", "binary.bin")
	err = bcpp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bcpp_src_url := "https://invisible-mirror.net/archives/bcpp/bcpp-20240111.tgz"
	bcpp_cmd_src := exec.Command("curl", "-L", bcpp_src_url, "-o", "source.tar.gz")
	err = bcpp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bcpp_cmd_direct := exec.Command("./binary")
	err = bcpp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

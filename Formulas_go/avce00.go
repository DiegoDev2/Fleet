package main

// Avce00 - Make Arc/Info (binary) Vector Coverages appear as E00
// Homepage: http://avce00.maptools.org/avce00/index.html

import (
	"fmt"
	
	"os/exec"
)

func installAvce00() {
	// Método 1: Descargar y extraer .tar.gz
	avce00_tar_url := "http://avce00.maptools.org/dl/avce00-2.0.0.tar.gz"
	avce00_cmd_tar := exec.Command("curl", "-L", avce00_tar_url, "-o", "package.tar.gz")
	err := avce00_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	avce00_zip_url := "http://avce00.maptools.org/dl/avce00-2.0.0.zip"
	avce00_cmd_zip := exec.Command("curl", "-L", avce00_zip_url, "-o", "package.zip")
	err = avce00_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	avce00_bin_url := "http://avce00.maptools.org/dl/avce00-2.0.0.bin"
	avce00_cmd_bin := exec.Command("curl", "-L", avce00_bin_url, "-o", "binary.bin")
	err = avce00_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	avce00_src_url := "http://avce00.maptools.org/dl/avce00-2.0.0.src.tar.gz"
	avce00_cmd_src := exec.Command("curl", "-L", avce00_src_url, "-o", "source.tar.gz")
	err = avce00_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	avce00_cmd_direct := exec.Command("./binary")
	err = avce00_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

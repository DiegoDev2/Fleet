package main

// Pbzip2 - Parallel bzip2
// Homepage: http://compression.great-site.net/pbzip2/

import (
	"fmt"
	
	"os/exec"
)

func installPbzip2() {
	// Método 1: Descargar y extraer .tar.gz
	pbzip2_tar_url := "https://launchpad.net/pbzip2/1.1/1.1.13/+download/pbzip2-1.1.13.tar.gz"
	pbzip2_cmd_tar := exec.Command("curl", "-L", pbzip2_tar_url, "-o", "package.tar.gz")
	err := pbzip2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pbzip2_zip_url := "https://launchpad.net/pbzip2/1.1/1.1.13/+download/pbzip2-1.1.13.zip"
	pbzip2_cmd_zip := exec.Command("curl", "-L", pbzip2_zip_url, "-o", "package.zip")
	err = pbzip2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pbzip2_bin_url := "https://launchpad.net/pbzip2/1.1/1.1.13/+download/pbzip2-1.1.13.bin"
	pbzip2_cmd_bin := exec.Command("curl", "-L", pbzip2_bin_url, "-o", "binary.bin")
	err = pbzip2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pbzip2_src_url := "https://launchpad.net/pbzip2/1.1/1.1.13/+download/pbzip2-1.1.13.src.tar.gz"
	pbzip2_cmd_src := exec.Command("curl", "-L", pbzip2_src_url, "-o", "source.tar.gz")
	err = pbzip2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pbzip2_cmd_direct := exec.Command("./binary")
	err = pbzip2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

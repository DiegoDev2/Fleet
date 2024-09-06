package main

// Greed - Game of consumption
// Homepage: http://www.catb.org/~esr/greed/

import (
	"fmt"
	
	"os/exec"
)

func installGreed() {
	// Método 1: Descargar y extraer .tar.gz
	greed_tar_url := "http://www.catb.org/~esr/greed/greed-4.3.tar.gz"
	greed_cmd_tar := exec.Command("curl", "-L", greed_tar_url, "-o", "package.tar.gz")
	err := greed_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	greed_zip_url := "http://www.catb.org/~esr/greed/greed-4.3.zip"
	greed_cmd_zip := exec.Command("curl", "-L", greed_zip_url, "-o", "package.zip")
	err = greed_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	greed_bin_url := "http://www.catb.org/~esr/greed/greed-4.3.bin"
	greed_cmd_bin := exec.Command("curl", "-L", greed_bin_url, "-o", "binary.bin")
	err = greed_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	greed_src_url := "http://www.catb.org/~esr/greed/greed-4.3.src.tar.gz"
	greed_cmd_src := exec.Command("curl", "-L", greed_src_url, "-o", "source.tar.gz")
	err = greed_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	greed_cmd_direct := exec.Command("./binary")
	err = greed_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

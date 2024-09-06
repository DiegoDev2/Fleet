package main

// Narwhal - General purpose JavaScript platform for building applications
// Homepage: https://github.com/280north/narwhal

import (
	"fmt"
	
	"os/exec"
)

func installNarwhal() {
	// Método 1: Descargar y extraer .tar.gz
	narwhal_tar_url := "https://github.com/280north/narwhal/archive/refs/tags/v0.3.2.tar.gz"
	narwhal_cmd_tar := exec.Command("curl", "-L", narwhal_tar_url, "-o", "package.tar.gz")
	err := narwhal_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	narwhal_zip_url := "https://github.com/280north/narwhal/archive/refs/tags/v0.3.2.zip"
	narwhal_cmd_zip := exec.Command("curl", "-L", narwhal_zip_url, "-o", "package.zip")
	err = narwhal_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	narwhal_bin_url := "https://github.com/280north/narwhal/archive/refs/tags/v0.3.2.bin"
	narwhal_cmd_bin := exec.Command("curl", "-L", narwhal_bin_url, "-o", "binary.bin")
	err = narwhal_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	narwhal_src_url := "https://github.com/280north/narwhal/archive/refs/tags/v0.3.2.src.tar.gz"
	narwhal_cmd_src := exec.Command("curl", "-L", narwhal_src_url, "-o", "source.tar.gz")
	err = narwhal_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	narwhal_cmd_direct := exec.Command("./binary")
	err = narwhal_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

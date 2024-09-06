package main

// Ioping - Tool to monitor I/O latency in real time
// Homepage: https://github.com/koct9i/ioping

import (
	"fmt"
	
	"os/exec"
)

func installIoping() {
	// Método 1: Descargar y extraer .tar.gz
	ioping_tar_url := "https://github.com/koct9i/ioping/archive/refs/tags/v1.3.tar.gz"
	ioping_cmd_tar := exec.Command("curl", "-L", ioping_tar_url, "-o", "package.tar.gz")
	err := ioping_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ioping_zip_url := "https://github.com/koct9i/ioping/archive/refs/tags/v1.3.zip"
	ioping_cmd_zip := exec.Command("curl", "-L", ioping_zip_url, "-o", "package.zip")
	err = ioping_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ioping_bin_url := "https://github.com/koct9i/ioping/archive/refs/tags/v1.3.bin"
	ioping_cmd_bin := exec.Command("curl", "-L", ioping_bin_url, "-o", "binary.bin")
	err = ioping_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ioping_src_url := "https://github.com/koct9i/ioping/archive/refs/tags/v1.3.src.tar.gz"
	ioping_cmd_src := exec.Command("curl", "-L", ioping_src_url, "-o", "source.tar.gz")
	err = ioping_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ioping_cmd_direct := exec.Command("./binary")
	err = ioping_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

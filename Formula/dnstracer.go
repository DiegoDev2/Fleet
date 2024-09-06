package main

// Dnstracer - Trace a chain of DNS servers to the source
// Homepage: http://www.mavetju.org/unix/dnstracer.php

import (
	"fmt"
	
	"os/exec"
)

func installDnstracer() {
	// Método 1: Descargar y extraer .tar.gz
	dnstracer_tar_url := "http://www.mavetju.org/download/dnstracer-1.9.tar.gz"
	dnstracer_cmd_tar := exec.Command("curl", "-L", dnstracer_tar_url, "-o", "package.tar.gz")
	err := dnstracer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dnstracer_zip_url := "http://www.mavetju.org/download/dnstracer-1.9.zip"
	dnstracer_cmd_zip := exec.Command("curl", "-L", dnstracer_zip_url, "-o", "package.zip")
	err = dnstracer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dnstracer_bin_url := "http://www.mavetju.org/download/dnstracer-1.9.bin"
	dnstracer_cmd_bin := exec.Command("curl", "-L", dnstracer_bin_url, "-o", "binary.bin")
	err = dnstracer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dnstracer_src_url := "http://www.mavetju.org/download/dnstracer-1.9.src.tar.gz"
	dnstracer_cmd_src := exec.Command("curl", "-L", dnstracer_src_url, "-o", "source.tar.gz")
	err = dnstracer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dnstracer_cmd_direct := exec.Command("./binary")
	err = dnstracer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

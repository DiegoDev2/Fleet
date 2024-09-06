package main

// Libltc - POSIX-C Library for handling Linear/Logitudinal Time Code (LTC)
// Homepage: https://x42.github.io/libltc/

import (
	"fmt"
	
	"os/exec"
)

func installLibltc() {
	// Método 1: Descargar y extraer .tar.gz
	libltc_tar_url := "https://github.com/x42/libltc/releases/download/v1.3.2/libltc-1.3.2.tar.gz"
	libltc_cmd_tar := exec.Command("curl", "-L", libltc_tar_url, "-o", "package.tar.gz")
	err := libltc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libltc_zip_url := "https://github.com/x42/libltc/releases/download/v1.3.2/libltc-1.3.2.zip"
	libltc_cmd_zip := exec.Command("curl", "-L", libltc_zip_url, "-o", "package.zip")
	err = libltc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libltc_bin_url := "https://github.com/x42/libltc/releases/download/v1.3.2/libltc-1.3.2.bin"
	libltc_cmd_bin := exec.Command("curl", "-L", libltc_bin_url, "-o", "binary.bin")
	err = libltc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libltc_src_url := "https://github.com/x42/libltc/releases/download/v1.3.2/libltc-1.3.2.src.tar.gz"
	libltc_cmd_src := exec.Command("curl", "-L", libltc_src_url, "-o", "source.tar.gz")
	err = libltc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libltc_cmd_direct := exec.Command("./binary")
	err = libltc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

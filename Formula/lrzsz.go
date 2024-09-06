package main

// Lrzsz - Tools for zmodem/xmodem/ymodem file transfer
// Homepage: https://www.ohse.de/uwe/software/lrzsz.html

import (
	"fmt"
	
	"os/exec"
)

func installLrzsz() {
	// Método 1: Descargar y extraer .tar.gz
	lrzsz_tar_url := "https://www.ohse.de/uwe/releases/lrzsz-0.12.20.tar.gz"
	lrzsz_cmd_tar := exec.Command("curl", "-L", lrzsz_tar_url, "-o", "package.tar.gz")
	err := lrzsz_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lrzsz_zip_url := "https://www.ohse.de/uwe/releases/lrzsz-0.12.20.zip"
	lrzsz_cmd_zip := exec.Command("curl", "-L", lrzsz_zip_url, "-o", "package.zip")
	err = lrzsz_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lrzsz_bin_url := "https://www.ohse.de/uwe/releases/lrzsz-0.12.20.bin"
	lrzsz_cmd_bin := exec.Command("curl", "-L", lrzsz_bin_url, "-o", "binary.bin")
	err = lrzsz_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lrzsz_src_url := "https://www.ohse.de/uwe/releases/lrzsz-0.12.20.src.tar.gz"
	lrzsz_cmd_src := exec.Command("curl", "-L", lrzsz_src_url, "-o", "source.tar.gz")
	err = lrzsz_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lrzsz_cmd_direct := exec.Command("./binary")
	err = lrzsz_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

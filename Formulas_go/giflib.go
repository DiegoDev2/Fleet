package main

// Giflib - Library and utilities for processing GIFs
// Homepage: https://giflib.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installGiflib() {
	// Método 1: Descargar y extraer .tar.gz
	giflib_tar_url := "https://downloads.sourceforge.net/project/giflib/giflib-5.2.2.tar.gz"
	giflib_cmd_tar := exec.Command("curl", "-L", giflib_tar_url, "-o", "package.tar.gz")
	err := giflib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	giflib_zip_url := "https://downloads.sourceforge.net/project/giflib/giflib-5.2.2.zip"
	giflib_cmd_zip := exec.Command("curl", "-L", giflib_zip_url, "-o", "package.zip")
	err = giflib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	giflib_bin_url := "https://downloads.sourceforge.net/project/giflib/giflib-5.2.2.bin"
	giflib_cmd_bin := exec.Command("curl", "-L", giflib_bin_url, "-o", "binary.bin")
	err = giflib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	giflib_src_url := "https://downloads.sourceforge.net/project/giflib/giflib-5.2.2.src.tar.gz"
	giflib_cmd_src := exec.Command("curl", "-L", giflib_src_url, "-o", "source.tar.gz")
	err = giflib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	giflib_cmd_direct := exec.Command("./binary")
	err = giflib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

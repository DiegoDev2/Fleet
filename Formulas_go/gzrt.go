package main

// Gzrt - Gzip recovery toolkit
// Homepage: https://www.urbanophile.com/arenn/coding/gzrt/gzrt.html

import (
	"fmt"
	
	"os/exec"
)

func installGzrt() {
	// Método 1: Descargar y extraer .tar.gz
	gzrt_tar_url := "https://www.urbanophile.com/arenn/coding/gzrt/gzrt-0.8.tar.gz"
	gzrt_cmd_tar := exec.Command("curl", "-L", gzrt_tar_url, "-o", "package.tar.gz")
	err := gzrt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gzrt_zip_url := "https://www.urbanophile.com/arenn/coding/gzrt/gzrt-0.8.zip"
	gzrt_cmd_zip := exec.Command("curl", "-L", gzrt_zip_url, "-o", "package.zip")
	err = gzrt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gzrt_bin_url := "https://www.urbanophile.com/arenn/coding/gzrt/gzrt-0.8.bin"
	gzrt_cmd_bin := exec.Command("curl", "-L", gzrt_bin_url, "-o", "binary.bin")
	err = gzrt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gzrt_src_url := "https://www.urbanophile.com/arenn/coding/gzrt/gzrt-0.8.src.tar.gz"
	gzrt_cmd_src := exec.Command("curl", "-L", gzrt_src_url, "-o", "source.tar.gz")
	err = gzrt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gzrt_cmd_direct := exec.Command("./binary")
	err = gzrt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

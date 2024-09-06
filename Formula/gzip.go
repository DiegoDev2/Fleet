package main

// Gzip - Popular GNU data compression program
// Homepage: https://www.gnu.org/software/gzip/

import (
	"fmt"
	
	"os/exec"
)

func installGzip() {
	// Método 1: Descargar y extraer .tar.gz
	gzip_tar_url := "https://ftp.gnu.org/gnu/gzip/gzip-1.13.tar.gz"
	gzip_cmd_tar := exec.Command("curl", "-L", gzip_tar_url, "-o", "package.tar.gz")
	err := gzip_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gzip_zip_url := "https://ftp.gnu.org/gnu/gzip/gzip-1.13.zip"
	gzip_cmd_zip := exec.Command("curl", "-L", gzip_zip_url, "-o", "package.zip")
	err = gzip_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gzip_bin_url := "https://ftp.gnu.org/gnu/gzip/gzip-1.13.bin"
	gzip_cmd_bin := exec.Command("curl", "-L", gzip_bin_url, "-o", "binary.bin")
	err = gzip_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gzip_src_url := "https://ftp.gnu.org/gnu/gzip/gzip-1.13.src.tar.gz"
	gzip_cmd_src := exec.Command("curl", "-L", gzip_src_url, "-o", "source.tar.gz")
	err = gzip_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gzip_cmd_direct := exec.Command("./binary")
	err = gzip_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

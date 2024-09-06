package main

// Fcrackzip - Zip password cracker
// Homepage: http://oldhome.schmorp.de/marc/fcrackzip.html

import (
	"fmt"
	
	"os/exec"
)

func installFcrackzip() {
	// Método 1: Descargar y extraer .tar.gz
	fcrackzip_tar_url := "http://oldhome.schmorp.de/marc/data/fcrackzip-1.0.tar.gz"
	fcrackzip_cmd_tar := exec.Command("curl", "-L", fcrackzip_tar_url, "-o", "package.tar.gz")
	err := fcrackzip_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fcrackzip_zip_url := "http://oldhome.schmorp.de/marc/data/fcrackzip-1.0.zip"
	fcrackzip_cmd_zip := exec.Command("curl", "-L", fcrackzip_zip_url, "-o", "package.zip")
	err = fcrackzip_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fcrackzip_bin_url := "http://oldhome.schmorp.de/marc/data/fcrackzip-1.0.bin"
	fcrackzip_cmd_bin := exec.Command("curl", "-L", fcrackzip_bin_url, "-o", "binary.bin")
	err = fcrackzip_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fcrackzip_src_url := "http://oldhome.schmorp.de/marc/data/fcrackzip-1.0.src.tar.gz"
	fcrackzip_cmd_src := exec.Command("curl", "-L", fcrackzip_src_url, "-o", "source.tar.gz")
	err = fcrackzip_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fcrackzip_cmd_direct := exec.Command("./binary")
	err = fcrackzip_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

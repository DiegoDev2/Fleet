package main

// Clzip - C language version of lzip
// Homepage: https://www.nongnu.org/lzip/clzip.html

import (
	"fmt"
	
	"os/exec"
)

func installClzip() {
	// Método 1: Descargar y extraer .tar.gz
	clzip_tar_url := "https://download.savannah.gnu.org/releases/lzip/clzip/clzip-1.14.tar.gz"
	clzip_cmd_tar := exec.Command("curl", "-L", clzip_tar_url, "-o", "package.tar.gz")
	err := clzip_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clzip_zip_url := "https://download.savannah.gnu.org/releases/lzip/clzip/clzip-1.14.zip"
	clzip_cmd_zip := exec.Command("curl", "-L", clzip_zip_url, "-o", "package.zip")
	err = clzip_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clzip_bin_url := "https://download.savannah.gnu.org/releases/lzip/clzip/clzip-1.14.bin"
	clzip_cmd_bin := exec.Command("curl", "-L", clzip_bin_url, "-o", "binary.bin")
	err = clzip_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clzip_src_url := "https://download.savannah.gnu.org/releases/lzip/clzip/clzip-1.14.src.tar.gz"
	clzip_cmd_src := exec.Command("curl", "-L", clzip_src_url, "-o", "source.tar.gz")
	err = clzip_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clzip_cmd_direct := exec.Command("./binary")
	err = clzip_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

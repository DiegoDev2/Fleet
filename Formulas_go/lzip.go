package main

// Lzip - LZMA-based compression program similar to gzip or bzip2
// Homepage: https://www.nongnu.org/lzip/

import (
	"fmt"
	
	"os/exec"
)

func installLzip() {
	// Método 1: Descargar y extraer .tar.gz
	lzip_tar_url := "https://download-mirror.savannah.gnu.org/releases/lzip/lzip-1.24.1.tar.gz"
	lzip_cmd_tar := exec.Command("curl", "-L", lzip_tar_url, "-o", "package.tar.gz")
	err := lzip_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lzip_zip_url := "https://download-mirror.savannah.gnu.org/releases/lzip/lzip-1.24.1.zip"
	lzip_cmd_zip := exec.Command("curl", "-L", lzip_zip_url, "-o", "package.zip")
	err = lzip_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lzip_bin_url := "https://download-mirror.savannah.gnu.org/releases/lzip/lzip-1.24.1.bin"
	lzip_cmd_bin := exec.Command("curl", "-L", lzip_bin_url, "-o", "binary.bin")
	err = lzip_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lzip_src_url := "https://download-mirror.savannah.gnu.org/releases/lzip/lzip-1.24.1.src.tar.gz"
	lzip_cmd_src := exec.Command("curl", "-L", lzip_src_url, "-o", "source.tar.gz")
	err = lzip_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lzip_cmd_direct := exec.Command("./binary")
	err = lzip_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

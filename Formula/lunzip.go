package main

// Lunzip - Decompressor for lzip files
// Homepage: https://www.nongnu.org/lzip/lunzip.html

import (
	"fmt"
	
	"os/exec"
)

func installLunzip() {
	// Método 1: Descargar y extraer .tar.gz
	lunzip_tar_url := "https://download-mirror.savannah.gnu.org/releases/lzip/lunzip/lunzip-1.14.tar.gz"
	lunzip_cmd_tar := exec.Command("curl", "-L", lunzip_tar_url, "-o", "package.tar.gz")
	err := lunzip_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lunzip_zip_url := "https://download-mirror.savannah.gnu.org/releases/lzip/lunzip/lunzip-1.14.zip"
	lunzip_cmd_zip := exec.Command("curl", "-L", lunzip_zip_url, "-o", "package.zip")
	err = lunzip_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lunzip_bin_url := "https://download-mirror.savannah.gnu.org/releases/lzip/lunzip/lunzip-1.14.bin"
	lunzip_cmd_bin := exec.Command("curl", "-L", lunzip_bin_url, "-o", "binary.bin")
	err = lunzip_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lunzip_src_url := "https://download-mirror.savannah.gnu.org/releases/lzip/lunzip/lunzip-1.14.src.tar.gz"
	lunzip_cmd_src := exec.Command("curl", "-L", lunzip_src_url, "-o", "source.tar.gz")
	err = lunzip_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lunzip_cmd_direct := exec.Command("./binary")
	err = lunzip_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: lzip")
	exec.Command("latte", "install", "lzip").Run()
}

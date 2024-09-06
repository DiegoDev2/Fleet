package main

// Plzip - Data compressor
// Homepage: https://www.nongnu.org/lzip/plzip.html

import (
	"fmt"
	
	"os/exec"
)

func installPlzip() {
	// Método 1: Descargar y extraer .tar.gz
	plzip_tar_url := "https://download.savannah.gnu.org/releases/lzip/plzip/plzip-1.11.tar.gz"
	plzip_cmd_tar := exec.Command("curl", "-L", plzip_tar_url, "-o", "package.tar.gz")
	err := plzip_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	plzip_zip_url := "https://download.savannah.gnu.org/releases/lzip/plzip/plzip-1.11.zip"
	plzip_cmd_zip := exec.Command("curl", "-L", plzip_zip_url, "-o", "package.zip")
	err = plzip_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	plzip_bin_url := "https://download.savannah.gnu.org/releases/lzip/plzip/plzip-1.11.bin"
	plzip_cmd_bin := exec.Command("curl", "-L", plzip_bin_url, "-o", "binary.bin")
	err = plzip_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	plzip_src_url := "https://download.savannah.gnu.org/releases/lzip/plzip/plzip-1.11.src.tar.gz"
	plzip_cmd_src := exec.Command("curl", "-L", plzip_src_url, "-o", "source.tar.gz")
	err = plzip_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	plzip_cmd_direct := exec.Command("./binary")
	err = plzip_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: lzlib")
exec.Command("latte", "install", "lzlib")
}

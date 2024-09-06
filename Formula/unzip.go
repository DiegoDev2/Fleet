package main

// Unzip - Extraction utility for .zip compressed archives
// Homepage: https://infozip.sourceforge.net/UnZip.html

import (
	"fmt"
	
	"os/exec"
)

func installUnzip() {
	// Método 1: Descargar y extraer .tar.gz
	unzip_tar_url := "https://downloads.sourceforge.net/project/infozip/UnZip%206.x%20%28latest%29/UnZip%206.0/unzip60.tar.gz"
	unzip_cmd_tar := exec.Command("curl", "-L", unzip_tar_url, "-o", "package.tar.gz")
	err := unzip_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	unzip_zip_url := "https://downloads.sourceforge.net/project/infozip/UnZip%206.x%20%28latest%29/UnZip%206.0/unzip60.zip"
	unzip_cmd_zip := exec.Command("curl", "-L", unzip_zip_url, "-o", "package.zip")
	err = unzip_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	unzip_bin_url := "https://downloads.sourceforge.net/project/infozip/UnZip%206.x%20%28latest%29/UnZip%206.0/unzip60.bin"
	unzip_cmd_bin := exec.Command("curl", "-L", unzip_bin_url, "-o", "binary.bin")
	err = unzip_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	unzip_src_url := "https://downloads.sourceforge.net/project/infozip/UnZip%206.x%20%28latest%29/UnZip%206.0/unzip60.src.tar.gz"
	unzip_cmd_src := exec.Command("curl", "-L", unzip_src_url, "-o", "source.tar.gz")
	err = unzip_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	unzip_cmd_direct := exec.Command("./binary")
	err = unzip_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

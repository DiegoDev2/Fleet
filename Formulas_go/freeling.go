package main

// Freeling - Suite of language analyzers
// Homepage: https://nlp.lsi.upc.edu/freeling/

import (
	"fmt"
	
	"os/exec"
)

func installFreeling() {
	// Método 1: Descargar y extraer .tar.gz
	freeling_tar_url := "https://github.com/TALP-UPC/FreeLing/releases/download/4.2/FreeLing-src-4.2.1.tar.gz"
	freeling_cmd_tar := exec.Command("curl", "-L", freeling_tar_url, "-o", "package.tar.gz")
	err := freeling_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	freeling_zip_url := "https://github.com/TALP-UPC/FreeLing/releases/download/4.2/FreeLing-src-4.2.1.zip"
	freeling_cmd_zip := exec.Command("curl", "-L", freeling_zip_url, "-o", "package.zip")
	err = freeling_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	freeling_bin_url := "https://github.com/TALP-UPC/FreeLing/releases/download/4.2/FreeLing-src-4.2.1.bin"
	freeling_cmd_bin := exec.Command("curl", "-L", freeling_bin_url, "-o", "binary.bin")
	err = freeling_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	freeling_src_url := "https://github.com/TALP-UPC/FreeLing/releases/download/4.2/FreeLing-src-4.2.1.src.tar.gz"
	freeling_cmd_src := exec.Command("curl", "-L", freeling_src_url, "-o", "source.tar.gz")
	err = freeling_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	freeling_cmd_direct := exec.Command("./binary")
	err = freeling_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: icu4c")
exec.Command("latte", "install", "icu4c")
}

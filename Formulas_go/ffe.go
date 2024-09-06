package main

// Ffe - Parse flat file structures and print them in different formats
// Homepage: https://ff-extractor.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installFfe() {
	// Método 1: Descargar y extraer .tar.gz
	ffe_tar_url := "https://downloads.sourceforge.net/project/ff-extractor/0.3.9a/0.3.9a.tar.gz"
	ffe_cmd_tar := exec.Command("curl", "-L", ffe_tar_url, "-o", "package.tar.gz")
	err := ffe_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ffe_zip_url := "https://downloads.sourceforge.net/project/ff-extractor/0.3.9a/0.3.9a.zip"
	ffe_cmd_zip := exec.Command("curl", "-L", ffe_zip_url, "-o", "package.zip")
	err = ffe_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ffe_bin_url := "https://downloads.sourceforge.net/project/ff-extractor/0.3.9a/0.3.9a.bin"
	ffe_cmd_bin := exec.Command("curl", "-L", ffe_bin_url, "-o", "binary.bin")
	err = ffe_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ffe_src_url := "https://downloads.sourceforge.net/project/ff-extractor/0.3.9a/0.3.9a.src.tar.gz"
	ffe_cmd_src := exec.Command("curl", "-L", ffe_src_url, "-o", "source.tar.gz")
	err = ffe_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ffe_cmd_direct := exec.Command("./binary")
	err = ffe_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

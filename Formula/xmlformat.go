package main

// Xmlformat - Format XML documents
// Homepage: https://web.archive.org/web/20160929174540/http://www.kitebird.com/software/xmlformat/

import (
	"fmt"
	
	"os/exec"
)

func installXmlformat() {
	// Método 1: Descargar y extraer .tar.gz
	xmlformat_tar_url := "https://web.archive.org/web/20161110001923/http://www.kitebird.com/software/xmlformat/xmlformat-1.04.tar.gz"
	xmlformat_cmd_tar := exec.Command("curl", "-L", xmlformat_tar_url, "-o", "package.tar.gz")
	err := xmlformat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xmlformat_zip_url := "https://web.archive.org/web/20161110001923/http://www.kitebird.com/software/xmlformat/xmlformat-1.04.zip"
	xmlformat_cmd_zip := exec.Command("curl", "-L", xmlformat_zip_url, "-o", "package.zip")
	err = xmlformat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xmlformat_bin_url := "https://web.archive.org/web/20161110001923/http://www.kitebird.com/software/xmlformat/xmlformat-1.04.bin"
	xmlformat_cmd_bin := exec.Command("curl", "-L", xmlformat_bin_url, "-o", "binary.bin")
	err = xmlformat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xmlformat_src_url := "https://web.archive.org/web/20161110001923/http://www.kitebird.com/software/xmlformat/xmlformat-1.04.src.tar.gz"
	xmlformat_cmd_src := exec.Command("curl", "-L", xmlformat_src_url, "-o", "source.tar.gz")
	err = xmlformat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xmlformat_cmd_direct := exec.Command("./binary")
	err = xmlformat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

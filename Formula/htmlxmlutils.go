package main

// HtmlXmlUtils - Tools for manipulating HTML and XML files
// Homepage: https://www.w3.org/Tools/HTML-XML-utils/

import (
	"fmt"
	
	"os/exec"
)

func installHtmlXmlUtils() {
	// Método 1: Descargar y extraer .tar.gz
	htmlxmlutils_tar_url := "https://www.w3.org/Tools/HTML-XML-utils/html-xml-utils-8.6.tar.gz"
	htmlxmlutils_cmd_tar := exec.Command("curl", "-L", htmlxmlutils_tar_url, "-o", "package.tar.gz")
	err := htmlxmlutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	htmlxmlutils_zip_url := "https://www.w3.org/Tools/HTML-XML-utils/html-xml-utils-8.6.zip"
	htmlxmlutils_cmd_zip := exec.Command("curl", "-L", htmlxmlutils_zip_url, "-o", "package.zip")
	err = htmlxmlutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	htmlxmlutils_bin_url := "https://www.w3.org/Tools/HTML-XML-utils/html-xml-utils-8.6.bin"
	htmlxmlutils_cmd_bin := exec.Command("curl", "-L", htmlxmlutils_bin_url, "-o", "binary.bin")
	err = htmlxmlutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	htmlxmlutils_src_url := "https://www.w3.org/Tools/HTML-XML-utils/html-xml-utils-8.6.src.tar.gz"
	htmlxmlutils_cmd_src := exec.Command("curl", "-L", htmlxmlutils_src_url, "-o", "source.tar.gz")
	err = htmlxmlutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	htmlxmlutils_cmd_direct := exec.Command("./binary")
	err = htmlxmlutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

package main

// XmlCoreutils - Powerful interactive system for text processing
// Homepage: https://xml-coreutils.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installXmlCoreutils() {
	// Método 1: Descargar y extraer .tar.gz
	xmlcoreutils_tar_url := "https://downloads.sourceforge.net/project/xml-coreutils/xml-coreutils-0.8.1.tar.gz"
	xmlcoreutils_cmd_tar := exec.Command("curl", "-L", xmlcoreutils_tar_url, "-o", "package.tar.gz")
	err := xmlcoreutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xmlcoreutils_zip_url := "https://downloads.sourceforge.net/project/xml-coreutils/xml-coreutils-0.8.1.zip"
	xmlcoreutils_cmd_zip := exec.Command("curl", "-L", xmlcoreutils_zip_url, "-o", "package.zip")
	err = xmlcoreutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xmlcoreutils_bin_url := "https://downloads.sourceforge.net/project/xml-coreutils/xml-coreutils-0.8.1.bin"
	xmlcoreutils_cmd_bin := exec.Command("curl", "-L", xmlcoreutils_bin_url, "-o", "binary.bin")
	err = xmlcoreutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xmlcoreutils_src_url := "https://downloads.sourceforge.net/project/xml-coreutils/xml-coreutils-0.8.1.src.tar.gz"
	xmlcoreutils_cmd_src := exec.Command("curl", "-L", xmlcoreutils_src_url, "-o", "source.tar.gz")
	err = xmlcoreutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xmlcoreutils_cmd_direct := exec.Command("./binary")
	err = xmlcoreutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: s-lang")
exec.Command("latte", "install", "s-lang")
}
